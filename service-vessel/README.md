# ServiceVessel Service

This is the ServiceVessel service

Generated with

```
micro new --namespace=shippy --type=service service-vessel
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: shippy.service.service-vessel
- Type: service
- Alias: service-vessel

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
./service-vessel-service
```

Build a docker image
```
make docker
```