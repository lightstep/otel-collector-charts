## recommended-collector-config

**Recommended** chart for sending Prometheus metrics without Kubernetes to ServiceNow Cloud Observability using OpenTelemetry-native metric collection and the OpenTelemetry Operator.

# Locally Testing OpenTelemetry Collector Builder

This program generates a custom OpenTelemetry Collector binary based on the example configuration.

## Requirements

- **Go Binaries** allows developers to install Go programs from the command-line, without requiring Go to be installed on your machine.

## Installation

To install the OpenTelemetry Collector builder, run:

```bash
cd ~/otel-collector-charts/arrow
GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@latest
```

### Updating to latest version

1. Open the otelcolarrow-build.yaml file and update any references to your desired version.

2. Uncomment lines in example/config.yaml that refer to required components, such as:

```yaml
# - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor v0.103.0
# - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.103.0
# - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver v0.103.0
# - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver v0.103.0
# - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver v0.103.0
```

### Usage

- Generate the custom collector binary:

```bash
builder --config ./otelcolarrow-build.yaml
```

- Run the custom collector:

```bash
./dist/otelarrowcol --config ../example/vm/config.yaml
```

#### Troubleshooting

- Note this example config does not use the concurrent batch processor. Uncomment this line instead:

` - gomod: go.opentelemetry.io/collector/processor/batchprocessor v0.103.0`
