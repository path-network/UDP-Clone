# UDP Clone

A tiny UDP proxy that can replicate traffic to one or more endpoints.

## Why?

We needed a way to take a stream of NetFlow/IPFIX/sFlow and send it to multiple endpoints(Prod, Staging, Testing). As this is a generic UDP replicator, it can be used for any traffic such as netflow, syslog etc.

## Usage

### Docker Compose

This project is meant to managed via single docker compose. Endpoints can easily be modified in the `docker-compose.yml` and safely reloaded by running `./up.sh`

```
justin@ops1:~/udp-clone$ ./up.sh 
[+] Building 15.5s (20/30)                                                                                                                
 => [udp-clone_netflow5 internal] load build definition from Dockerfile                                                              0.1s
 => => transferring dockerfile: 284B                                                                                                 0.0s
 => [udp-clone_netflow9 internal] load build definition from Dockerfile                                                              0.1s
 => => transferring dockerfile: 284B                                                                                                 0.0s
 => [udp-clone_sflow internal] load build definition from Dockerfile                                                                 0.1s
 => => transferring dockerfile: 284B                                                                                                 0.0s
 => [udp-clone_netflow5 internal] load .dockerignore                                                                                 0.1s
 => => transferring context: 2B                                                                                                      0.0s
 => [udp-clone_sflow internal] load .dockerignore                                                                                    0.1s
 => => transferring context: 2B                                                                                                      0.0s
 => [udp-clone_netflow9 internal] load .dockerignore                                                                                 0.1s
 => => transferring context: 2B                                                                                                      0.0s
 => [udp-clone_netflow9 internal] load metadata for docker.io/library/ubuntu:latest                                                  0.0s
 => [udp-clone_netflow9 internal] load metadata for docker.io/library/golang:1.18-bullseye                                           5.0s
 => CACHED [udp-clone_netflow5 stage-1 1/2] FROM docker.io/library/ubuntu:latest                                                     0.0s
 => [udp-clone_sflow build-img 1/6] FROM docker.io/library/golang:1.18-bullseye@sha256:db42e4bb1a7f32da1ec430906769dbbabe9f1868bd41  6.0s
 => => resolve docker.io/library/golang:1.18-bullseye@sha256:db42e4bb1a7f32da1ec430906769dbbabe9f1868bd4170751e4923f1b8948a45        0.1s
 => => sha256:e604223835ccf02d097187b5a58ca73e8598cadbb16a36202ca1943e97f56f1f 10.88MB / 10.88MB                                     0.5s
 => => sha256:db42e4bb1a7f32da1ec430906769dbbabe9f1868bd4170751e4923f1b8948a45 1.86kB / 1.86kB                                       0.0s
 => => sha256:bf168a6748997eb97b48cc86234b7ff7d8bc907645b9be99013158b3f146b272 5.16MB / 5.16MB                                       0.3s
 => => sha256:5417b4917fa7ed3ad2678a3ce6378a00c95bfd430c2ffa39936fce55130b5f2c 1.80kB / 1.80kB                                       0.0s
 => => sha256:76199a964a3fc66e31bda713381e92285f479fe8e3d4514a473f95ffc2062440 7.10kB / 7.10kB                                       0.0s
 => => sha256:e756f3fdd6a378aa16205b0f75d178b7532b110e86be7659004fc6a21183226c 55.01MB / 55.01MB                                     0.7s
 => => sha256:6d5c91c4cd86dde23108ab3af91e9eae838d0059a380ee7dfd4f370b6d985523 54.58MB / 54.58MB                                     0.8s
 => => sha256:93c221c34e03cb2bc3c5cb0e1fcf029b793cfe2c10362287dd05270d80333db9 85.87MB / 85.87MB                                     1.0s
 => => extracting sha256:e756f3fdd6a378aa16205b0f75d178b7532b110e86be7659004fc6a21183226c                                            0.6s
 => => sha256:399edca3a0ef467dadd57f6ed1ee48c7b64162ca25d1fae2940680b749c722a9 141.75MB / 141.75MB                                   1.4s
 => => sha256:00fc5c011105d0ac8b5453886bb3b836c81260e1d016938a1207d48da8f28718 155B / 155B                                           0.9s
 => => extracting sha256:bf168a6748997eb97b48cc86234b7ff7d8bc907645b9be99013158b3f146b272                                            0.1s
 => => extracting sha256:e604223835ccf02d097187b5a58ca73e8598cadbb16a36202ca1943e97f56f1f                                            0.1s
 => => extracting sha256:6d5c91c4cd86dde23108ab3af91e9eae838d0059a380ee7dfd4f370b6d985523                                            0.6s
 => => extracting sha256:93c221c34e03cb2bc3c5cb0e1fcf029b793cfe2c10362287dd05270d80333db9                                            0.8s
 => => extracting sha256:399edca3a0ef467dadd57f6ed1ee48c7b64162ca25d1fae2940680b749c722a9                                            1.7s
 => => extracting sha256:00fc5c011105d0ac8b5453886bb3b836c81260e1d016938a1207d48da8f28718                                            0.0s
 => [udp-clone_netflow9 internal] load build context                                                                                 0.2s
 => => transferring context: 4.59MB                                                                                                  0.1s
 => [udp-clone_netflow5 internal] load build context                                                                                 0.2s
 => => transferring context: 4.59MB                                                                                                  0.1s
 => [udp-clone_sflow internal] load build context                                                                                    0.2s
 => => transferring context: 4.59MB                                                                                                  0.1s
 => [udp-clone_sflow build-img 2/6] WORKDIR /go/src/app                                                                              0.6s
 => [udp-clone_netflow9 build-img 3/6] COPY go.mod go.sum ./                                                                         0.1s
 => [udp-clone_netflow9 build-img 4/6] RUN go mod download                                                                           0.7s
 => [udp-clone_netflow9 build-img 5/6] COPY . .                                                                                      0.1s
 => [udp-clone_netflow9 build-img 6/6] RUN go build                                                                                  2.4s
 => [udp-clone_netflow9 stage-1 2/2] COPY --from=build-img /go/src/app/udp-clone /bin/udp-clone                                      0.1s
 => [udp-clone_netflow5] exporting to image                                                                                          0.3s
 => => exporting layers                                                                                                              0.1s
 => => writing image sha256:34530238bc121bc2c8a8daa2bbedc48ef36ccdca90f31a81cd4e510d27ad455f                                         0.0s
 => => naming to docker.io/library/udp-clone_netflow9                                                                                0.0s
 => => naming to docker.io/library/udp-clone_sflow                                                                                   0.0s
 => => naming to docker.io/library/udp-clone_netflow5                                                                                0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
[+] Running 4/4
 _ Network udp-clone_default     Created                                                                                             0.1s
 _ Container udp-clone-netflow9  Started                                                                                             0.9s
 _ Container udp-clone-sflow     Started                                                                                             0.8s
 _ Container udp-clone-netflow5  Started                                                                                             0.7s

```


### Direct

To run this project directly, build it manually using `go build`. 

```
./udp-clone --listen-port=9500 --forward=192.0.2.1 --forward=198.51.100.5:9100
time="2019-09-23T14:29:50+01:00" level=info msg="Server started" ip=0.0.0.0 port=9500
time="2019-09-23T14:29:50+01:00" level=info msg="Forwarding target configured" addr="192.0.2.1:9500" num=1 total=2
time="2019-09-23T14:29:50+01:00" level=info msg="Forwarding target configured" addr="198.51.100.5:9100" num=2 total=2
```

The above command will:

- Start a UDP server listening on `0.0.0.0` port `9500`
- Add a forward target of `192.0.2.1:9500` (uses `listen-port` for destination port as not specified in configuration)
- Add another forward target of `198.51.100.5:9100`

The server will start listening on `0.0.0.0:9500`, any packet it receives will be replicated and sent to both `192.0.2.1:9500` and `198.51.100.5:9100`


## Configuration

```
usage: udp-clone [<flags>]

Flags:
  --help                 Show context-sensitive help (also try --help-long and --help-man).
  --debug                Enable debug mode
  --listen-ip=0.0.0.0    IP to listen in
  --listen-port=port     Port to listen on
  --body-size=10000      Size of body to read
  --forward=ip:port ...  ip:port to forward traffic to (port defaults to listen-port)
  --routines=10          Set the number of listeners per port
```

## Copyright

2022 Path Network 

Written By: Justin Timperio
