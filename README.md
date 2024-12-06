# ServiceNow Cloud Observability (formerly Lightstep) OpenTelemetry Collector Helm Charts

**⚠️ Warning:** These OpenTelemetry Helm charts are deprecated. We recommend the official OpenTelemetry community helm charts available at [https://github.com/open-telemetry/opentelemetry-helm-charts](https://github.com/open-telemetry/opentelemetry-helm-charts). 

This is the repository for [Helm](https://helm.sh/) charts for running an OpenTelemetry Collector using the [OpenTelemetry Operator for Kubernetes](https://github.com/open-telemetry/opentelemetry-operator). We recommend following the quick start documenation [here](https://docs.lightstep.com/docs/quick-start-infra-otel-first) for using these charts.


## Charts

- [otel-cloud-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/otel-cloud-stack) - **Deprecated** chart for sending Kubernetes metrics to ServiceNow Cloud Observability using OpenTelemetry-native metric collection and the OpenTelemetry Operator.
- [kube-otel-stack](https://github.com/lightstep/prometheus-k8s-opentelemetry-collector/tree/main/charts/kube-otel-stack) - **Deprecated** replacement for [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack), which uses the same configuration for scraping Prometheus exporters and forwarding metrics to Lightstep using the OpenTelemetry Operator. Use this chart if you are looking to compare Kubernetes monitoring in Prometheus with Kubernetes monitoring using ServiceNow Cloud Observability.

## Arrow Usage

> [!NOTE]
> Arrow usage is in beta, please use at your own risk. Reach out if you have any issues.

In order to use an arrow trace collector, you can use (1) the prebuilt image available via the Github Container Registry (GHCR) or you may (2) build your own custom image.

### 1. Use the prebuilt Docker image

1. We have built a Docker image using the recommended [build config](https://github.com/lightstep/otel-collector-charts/blob/main/arrow/otelcolarrow-build.yaml)
2. This Docker [image](https://github.com/lightstep/otel-collector-charts/pkgs/container/otel-collector-charts%2Fotelarrowcol-experimental) can be pulled by running: `docker pull ghcr.io/lightstep/otel-collector-charts/otelarrowcol-experimental:latest`
3. You can use the collector config (`/arrow/config/gateway-collector.yaml`) by running:
   `docker run -it -v $(PWD)/config/:/config --entrypoint /otelarrowcol ghcr.io/lightstep/otel-collector-charts/otelarrowcol-experimental:latest --config=/config/gateway-collector.yaml`

### 2. Build your own custom image

1. We have supplied a collector builder config below.
2. Once an image is a available, simply apply your desired helm chart with the values.yaml AND the arrow.yaml in the respective chart.
3. Make sure to replace the image in arrow.yaml with your custom built image.

## Build configurations

Some of the features available in these charts are optional because
they rely on components that have not been released in the
OpenTelemetry Contrib Collector. Specifically, to make use of the new
OpenTelemetry Protocol With Apache Arrow support requires using either
the prebuilt image or a customer collector build at this time.

See the [recommended custom collector build
configuration](./arrow/otelcolarrow-build.yaml) as a starting
point.

---

## Contributing and Developing

Please see [CONTRIBUTING.md.](./CONTRIBUTING.md)
