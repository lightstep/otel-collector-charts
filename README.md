Test!
# ServiceNow Cloud Observability (formerly Lightstep) OpenTelemetry Collector Helm Charts

This is the repository for recommended [Helm](https://helm.sh/) charts for running an OpenTelemetry Collector using the [OpenTelemetry Operator for Kubernetes](https://github.com/open-telemetry/opentelemetry-operator). We recommend following the quick start documenation [here](https://docs.lightstep.com/docs/quick-start-infra-otel-first) for using these charts.

⚠️ These OpenTelemetry Helm charts are under active development and may have breaking changes between releases.

## Charts

* [otel-cloud-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/otel-cloud-stack) - **Recommended** chart for sending Kubernetes metrics to ServiceNow Cloud Observability using OpenTelemetry-native metric collection and the OpenTelemetry Operator.
* [kube-otel-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/kube-otel-stack) - Drop in replacement for [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack), which uses the same configuration for scraping Prometheus exporters and forwarding metrics to Lightstep using the OpenTelemetry Operator. Use this chart if you are looking to compare Kubernetes monitoring in Prometheus with Kubernetes monitoring using ServiceNow Cloud Observability. 

## Arrow Usage

> [!NOTE] 
> Arrow usage is in beta, please use at your own risk. Reach out if you have any issues.

In order to use an arrow trace collector, you will need to build your own custom image. We have supplied a collector builder config below. Once an image is a available, simply apply your desired helm chart with the values.yaml AND the arrow.yaml in the respective chart. Make sure to replace the image in arrow.yaml with your custom built image.

## Build configurations

Some of the features available in these charts are optional because
they rely on components that have not been released in the
OpenTelemetry Contrib Collector.  Specifically, to make use of the new
OTel-Arrow protocol requires building a customer collector at this
time.  See a [recommended custom collector build configuration](./gateway-build.yaml).
