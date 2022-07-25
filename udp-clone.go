package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug      = kingpin.Flag("debug", "Enable debug mode").Envar("DEBUG").Bool()
	listenIP   = kingpin.Flag("listen-ip", "IP to listen in").Default("0.0.0.0").Envar("LISTEN_IP").IP()
	listenPort = kingpin.Flag("listen-port", "Port to listen on").PlaceHolder("port").Envar("LISTEN_PORT").Int()
	bodySize   = kingpin.Flag("body-size", "Size of body to read").Default("10000").Envar("BODY_SIZE").Int()
	routines   = kingpin.Flag("routines", "Set the number of listeners per port.").Default("10").Envar("ROUTINES").Int()
	forwards   = kingpin.Flag("forward", "ip:port to forward traffic to (port defaults to listen-port)").PlaceHolder("ip:port").Envar("FORWARD").Strings()

	targets []*net.UDPConn

	exit = make(chan bool, 1)
)

func main() {
	// CLI
	kingpin.Parse()

	rawJSON := []byte(`{
                "level": "info",
                "outputPaths": ["stdout"],
                "errorOutputPaths": ["stderr"],
                "encoding": "json",
                "encoderConfig": {
                        "messageKey": "message",
                        "levelKey": "level",
                        "levelEncoder": "lowercase"
                }
        }`)

	// Set up Zap logger
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := logger.Sugar()

	if *debug {
		cfg.Level.SetLevel(zap.ErrorLevel)
	}

	// Abort if no targets
	if len(*forwards) <= 0 {
		log.Fatal("Must specify at least one forward target")
	}

	// Create a list of vaild clients
	for _, forward := range *forwards {
		// Check for port
		if strings.Index(forward, ":") < 0 {
			forward = fmt.Sprintf("%s:%d", forward, *listenPort)
		}

		// Resolve
		addr, err := net.ResolveUDPAddr("udp", forward)
		if err != nil {
			log.Fatalf("Could not Resolve UDP Addr: %s (%s)", forward, err)
		}

		// Setup listenServer
		listenServer, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			log.Fatalf("Could not Dial UDP: %+v (%s)", addr, err)
		}
		defer listenServer.Close()

		targets = append(targets, listenServer)
	}

	// Listen Server
	listenServer, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: *listenPort,
		IP:   *listenIP,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer listenServer.Close()

	// Startup status
	log.Infow("Server started",
		"ip", *listenIP,
		"port", *listenPort,
	)

	for i, target := range targets {
		log.Infow("Forwarding target configured",
			"num", i+1,
			"total", len(targets),
			"addr", target.RemoteAddr(),
		)
	}

	// Launch listen routines
	for i := 0; i < *routines; i++ {
		go listenAndProxy(log, listenServer)
	}

	// Wait for exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		log.Infow("Exiting Gracefully from Interrupt Signal", "signal", sig)
		os.Exit(1)
	}
}

func listenAndProxy(log *zap.SugaredLogger, listenServer *net.UDPConn) {

	// Buffer for Packets
	buffer := make([]byte, *bodySize)

	for {
		// Read Packet
		n, source, err := listenServer.ReadFromUDP(buffer)
		if err != nil {
			log.Errorw("Failed to read from UDP",
				"error", err,
			)
			continue
		}

		// Proxy Packet
		for _, target := range targets {
			go proxyUDPpacket(log, target, source, buffer, n)
		}
	}
}

func proxyUDPpacket(log *zap.SugaredLogger, target *net.UDPConn, source *net.UDPAddr, buffer []byte, n int) {
	_, err := target.Write(buffer[:n])
	if err != nil {
		log.Warnw("Could not forward packet",
			"source", source.String(),
			"target", target.RemoteAddr(),
			"error", err,
			"time", time.Now().UTC().String(),
		)
	}
}
