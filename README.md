# Cloud Native API Layout (cnapi)

This repository contains a set of best practices from the web and foundations
like CNCF providing a starting point to develop REST APIs in Go. To get started
just clone the repository and make changes to it as it would be your project :).

## Environement

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

## Configuration

Configuration is an essential part of every application. This makes an
application adaptable to different environments like staging, production or
development. Applications store config in different forms. It is highly
recommended to seperate Code from config. A litus test to check if an
application is correctly separting code from config is where the codebase could
be made open-source at any moment and no credentials would be exposed.

The 12FA App recommends to store configuration in the environment of the
application e.g. as environemnt variables. This allows for an language agnostic
approach without accidentaly checking configuration files into
repositories.Antoher option recommended is using CLI flags which follow the
principles of environment variables.

In Go you can easily implement both approaches. If you want to implement a CLI
use the `cmd/` directory for it. You can either do it by using the standard
library or using something like `cobra`. Environment variables can be retrieved
using `os.Getenv`.

For a dynamic approach in the code base the `run` function accepts a `getenv`
function which controls the config of the environment allowing you to configure
your environment as you wish.

## Testing

## API-Gateways

## Principles of Chaos Engineering

### k8s chaos mesh

### monkey choas (netflix)

## Telemetry

Logging, Metrics, Traces

## Authentication and Authorization

## Encoding/Decoding and Validation

## Caching

## Kubernetes as Runtime

## Testing

### Unit

### Load

### Fault Tolerancy

## 12FA
