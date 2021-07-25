# GetCaptcha Service

This is the GetCaptcha service

Generated with

```
micro new bj38web/service/getCaptcha --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.getCaptcha
- Type: srv
- Alias: getCaptcha

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./getCaptcha-srv
```

Build a docker image
```
make docker
```