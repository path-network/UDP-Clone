FROM golang:1.18-bullseye AS build-img
ENV CGO_ENABLED=0
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

FROM ubuntu:latest
COPY --from=build-img /go/src/app/udp-clone /bin/udp-clone
ENTRYPOINT ["udp-clone"]