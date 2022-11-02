# Lightstep OpenTelemetry Collector Helm Charts

This is the repository for Lightstep's recommended [Helm](https://helm.sh/) charts for running an OpenTelemetry Collector using the [OpenTelemetry Operator for Kubernetes](https://github.com/open-telemetry/opentelemetry-operator). You can find documentation and tutorials for how to use these charts [here](https://docs.lightstep.com/docs/ingest-prometheus).

⚠️ Lightstep's OpenTelemetry Helm charts are under active development and may have breaking changes between releases.

## Charts

* [collector-k8s](https://github.com/lightstep/lightstep/otel-collector-charts/tree/main/charts/collector-k8s) - Chart for using the OpenTelemetry Collector to scape static or dynamic metric targets.
* [kube-otel-stack](https://github.com/lightstep/otel-collector-charts/tree/main/charts/kube-otel-stack) - Chart for sending Kubernetes metrics to Lightstep using the OpenTelemetry Operator.
* [http-check](https://github.com/lightstep/otel-collector-charts/tree/main/charts/kube-otel-stack) - Chart for running synthetic checks against HTTP endpoints.
