# Lightstep OpenTelemetry Collector Helm Charts

This is the repository for Lightstep's recommended [Helm](https://helm.sh/) charts for running an OpenTelemetry Collector using the [OpenTelemetry Operator for Kubernetes](https://github.com/open-telemetry/opentelemetry-operator). You can find documentation and tutorials for how to use these charts [here](https://docs.lightstep.com/docs/ingest-prometheus).

⚠️ Lightstep's OpenTelemetry Helm charts are under active development and may have breaking changes between releases.

## Charts

* [collector-k8s](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/collector-k8s) - Chart for using the OpenTelemetry Collector to scrape static or dynamic metric targets.
* [kube-otel-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/kube-otel-stack) - Chart for sending Kubernetes metrics to Lightstep using the OpenTelemetry Operator.
* [otel-cloud-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/otel-cloud-stack) - Chart for sending Kubernetes metrics to Lightstep using Otel native metric collection and the OpenTelemetry Operator.
