# Cloud Native API Layout (cnapi)

This repository contains a set of best practices from the web and foundations
like CNCF providing a starting point to develop REST APIs in Go. To get started
just clone the repository and make changes to it as it would be your project :).

## Stateless

Cloud native applications are always delivered as containers. So, it is save to
say that containers (and in a broader context Kubernetes) has become the
standard platform on which applications are running. For applications to be
efficient in a container environment they must follow a set of best practices.

Applications in containers execl when they are stateless. Meaning no data is
associate with the container or the application itself or in other words the
container could be deleted and started again without any effects on the
application. One example for stateful applications are databases.

Stateless applications are especially important because scaling horizontally is
straightforward by creating more replicas of the application.

## Service Discovery

If you application need to communicate with another service to run successfully
you need to somehow connect to the application. Using the IP-Address of the
container hosting the service in a Cloud-native environment is not recommended
because IP-Addresses are considered volatile and might change at any time. But
also it is not the responsiblity of the application to implement service
discovery.

The recommended approach to this problem is using your environment efficiently.
In most cases containers are not running without an orchestrater like
Kubernetes. These have built-in support for discovering service. In Kubernetes
this is done by creating an `Service` object. Applications which need to connect
to that service will be provided with the domain specific connection string
using environment variables, or CLI flags (see: [Config](#configuration))

## API-Gateways

## Principles of Chaos Engineering

### k8s chaos mesh

### monkey choas (netflix)

## Telemetry

## Configuration

## Authentication and Authorization

## Encoding/Decoding and Validation

## Caching

## Kubernetes as Runtime

## Testing

### Unit

### Load

### Fault Tolerancy

## 12FA
