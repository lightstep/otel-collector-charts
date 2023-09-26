# ServiceNow Cloud Observability (formerly Lightstep) OpenTelemetry Collector Helm Charts

This is the repository for recommended [Helm](https://helm.sh/) charts for running an OpenTelemetry Collector using the [OpenTelemetry Operator for Kubernetes](https://github.com/open-telemetry/opentelemetry-operator). We recommend following the quick start documenation [here](https://docs.lightstep.com/docs/quick-start-infra-otel-first) for using these charts.

⚠️ These OpenTelemetry Helm charts are under active development and may have breaking changes between releases.

## Charts

* [otel-cloud-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/otel-cloud-stack) - **Recommended** chart for sending Kubernetes metrics to ServiceNow Cloud Observability using OpenTelemetry-native metric collection and the OpenTelemetry Operator.
* [kube-otel-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/kube-otel-stack) - Drop in replacement for [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack), which uses the same configuration for scraping Prometheus exporters and forwarding metrics to Lightstep using the OpenTelemetry Operator.

