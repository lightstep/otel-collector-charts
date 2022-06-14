# Lightstep's OpenTelemetry Collector Prometheus Replacement
This is the repository for Lightstep's recommendations for running an opentelemetry collector that scrapes Prometheus targets. You can find the documentation for how to use this [here](https://docs.lightstep.com/docs/replace-prometheus-with-an-otel-collector-on-kubernetes)

# Deployment modes
OpenTelemetry Collector is a flexible system that can be deployed in several ways, here we list our recommended and tested modes.

## Deployment with single replica
For deploying the OpenTelemetry Collector as a deployment with a single replica, use the `collector_k8s` helm instructions.

## Daemonset + Deployment
For deploying the OpenTelemetry Collector with a mixed deployed composed of one Daemonset OpenTelemetry Collector - for collecting general application metrics - and a Deployment with single replica - for collecting insfrastructure targets , use the `daemonset_collector` helm instructions.