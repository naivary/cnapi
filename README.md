# Cloud Native API Layout (cnapi)

This repository contains a set of best practices from the web and foundations
like CNCF providing a starting point to develop REST APIs in Go. To get started
just clone the repository and make changes to it as it would be your project :).

## Environment

Cloud native applications are always delivered as container images and deployed
on Kubernetes (or some other Kubernetes flavor) for the orchestration of the
workloads. Given these circumntances it is save to say that Kubernetes is the
runtime of the Cloud and applications should be tailored to excel in these
conditions.

## Stateless

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
In Kubernetes this is done by creating an `Service` object. Applications which
need to connect to that service will be provided with the domain specific
connection string using the configuration options provided by the application
(see: [Config](#configuration))

## Configuration

Configuration is an essential part of every application. This makes an
application adaptable to different environments like staging, production or
development. Applications store config in different forms. It is highly
recommended to seperate Code from config. A litus test to check if an
application is correctly separting code from config is where the codebase could
be made open-source at any moment and no credentials would be exposed.

The 12FA App recommends to store configuration in the environment of the
application e.g. as environemnt variables. This allows for an language agnostic
approach without accidentaly checking configuration files into repositories.
Antoher option is using CLI flags which follow the principles of environment
variables. Both of the apparoaches have the advantage of adapting to the Cloud
environment e.g. Kubernetes because the values can be provided using various
Kubernetes sources (ConfigMap, Secrets etc.).

In Go you can easily implement both approaches. If you want to implement a CLI
use the `cmd/` directory for it. You can either do it by using the standard
library or using something like `cobra`. Environment variables can be retrieved
using `os.Getenv`.

For a dynamic approach in the code base the `run` function accepts a `getenv`
function which controls the config of the environment allowing you to configure
your environment as you wish.

## Principles of Chaos Engineering

Errors are part of every software and their occurence should be accepted as a
given fact. Because of that it is important to design APIs which are able to
cope with errors.

Chaos Engineering is enforcing exactly that. It is the discipline of
experimenting on a system in order to build confidence in the system's
capabiltiy to withstand turbulent conditions in production. For a detail
description consult the
[official documentation](https://principlesofchaos.org/).

To run local chaos experiments you can use
[Chaos Monkey](https://github.com/Netflix/chaosmonkey). A tool implementing the
principles of chaos engineering by Netflix. For chaos experiments on Kubernetes
use [Chaos Mesh](https://chaos-mesh.org/), a CNCF Incubating project. It is
recommened to choose the tool based on which is representing your production
environment the most.

## Documentation

Documentation of your application should be done in a _as code_ manner. Write
the documentation in markdown or mdx to allow for easy processing and generation
of content like a documentations website. The documentation of the API resources
and endpoints are implemented using the OpenAPI Standard. To reduce the amount
of boilerplate work you can use a tool which is generating the OpenAPI
Artificats based on your code. One example is
[codemark](https://github.com/naivary/codemark) which is generating any kind of
artifact based on (comment) markers.

## Validation

Validation is a crucial part of any HTTP request. Users can provide malicous
content or mismatch information. Therefor it is important to validate the
integrity of the payload before processing the request.

One easy approach to validate data is using a common interface for it. This
project provides the `Validator` interface in `validator.go`. The validator
takes in a context and provides feedback of the validations in form of a map. If
the returned map is empty then the validation was successfull. Combining this
approach with a JSON Schema is allowing for the most flexibile and industry
standard based solution.

## Testing

Unit, E2E, Load

## API-Gateways

## Telemetry

Logging, Metrics, Traces

## Authentication and Authorization

## Encoding/Decoding and Validation

## Caching
