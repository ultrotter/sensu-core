# Sensu Core

Core contains the definitions for the [sensu-go](https://github.com/sensu/sensu-go) core API.

## Codegen

Both of the core api versions have some component of protobuf-based code
generation. Each works with `go generate`, however there are system
dependencies in addition to the current supported version of go. A docker
image and Makefile is available to quickly set up the system dependencies,
but is hardcoded for x86.
