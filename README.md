# lwr

A lightweight replacement for the Lansweeper Agent on Linux that doesn't require privileged access and is small enough to audit.

## Docs

[API documentation](doc/api.md) describes the observed requests and responses from the local and cloud services.

## Build

Use `make build` to build `lwr`.

## Run

Use `lwr help` to discover available options.

### Basic usage

1. Create configuration file: `lwr configure -server lansweeper.example.org > /etc/lwr.conf`
2. Print a sample system report to review: `lwr print`
3. Send a report to the configured Lansweeper server: `lwr report`
