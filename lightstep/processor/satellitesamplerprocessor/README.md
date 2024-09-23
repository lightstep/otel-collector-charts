# Lightstep Satellite Sampler

This package contains an OpenTelemetry processor for an OpenTelemetry
traces pipeline that makes sampling decisions consistent with the
legacy Lightstep Satellite.  This component enables a slow transition
from Lightstep Satellites to OpenTelemetry Collectors without
simultaneously changing sampling algorithms.

 ## Recommended usage

This component supports operating a mixture of Lightstep Satellites
and OpenTelemetry Collectors with consistent probability sampling.
Here is a recommended sequence of steps for performing a migratation
to OpenTelemetry Collectors for Lightstep Satellite users.

### Build a custom OpenTelemetry Collector

This component is provided as a standalone component, meant for
incorporating into a custom build of the OpenTelemetry Collector using
the [OpenTelemetry Collector
builder](https://opentelemetry.io/docs/collector/custom-collector/)
tool.  In your Collector's build configuration, add the following
processor component:

```
  - gomod: github.com/lightstep/otel-collector-charts/lightstep/processor/satellitesamplerprocessor VERSIONTAG
```

where `VERSIONTAG` corresponds with the targetted OpenTelemetry
Collector release version.  At the time of this writing, the version
tag is `v0.109.0`.

Users are advised to include the OpenTelemetry Probabilistic Sampler
processor in their build, to complete this transition.  For example:

```
  - gomod: github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.109.0
```

Follow the usual steps to build your collector (e.g., `builder
--config build.yaml`).

### Configure the sampler

You will need to know the sampling probability configured with
Lightstep Satellites, in percentage terms.  Say the Lightstep
Satellite is configured with 10% sampling (i.e., 1-in-10).

Edit OpenTelemetry Collector configuration to include a
`satellitesatempler` block.  In the following example, the OTel-Arrow
receiver and exporter are configured with `satellitesampler` with 10%
sampling and [concurrent batch
processor](https://github.com/open-telemetry/otel-arrow/blob/main/collector/processor/concurrentbatchprocessor/README.md).

```
exporters:
  otelarrow:
    ...
receivers:
  otelarrow:
    ...
processors:
  satellitesampler:
    percent: 10
  concurrentbatch:
service:
  pipelines:
    traces:
      receivers: [otelarrow]
      processors: [satellitesampler, concurrentbatch]
      exporters: [otelarrow]
```

Collectors with this configuration may be deployed alongside a pool of
Lightstep Satellites sampling and the resulting traces will be
complete.

### Migrate to the OpenTelemetry Probabilistic Sampler

After decomissioning Lightstep Satellites and replacing them with
OpenTelemetry Collectors, users are advised to migrate to an
OpenTelemetry Collector processor with equivalent functionality, the
[Probabilistic Sampler Processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/processor/probabilisticsamplerprocessor/README.md).

A change of sampling configuration, either to change algorithm or to
change probability, typically results in broken traces.  Users are
advised to plan accordingly and make a quick transition between
samplers, with only a brief, planned period of broken traces.

Redeploy the pool of Collectors with the Probabilistic Sampler
processor configured instead of the Satellite sampler processor.  Make
this transition quickly, if possible, because traces will be
potentially incomplete as long as both samplers being used.

```
processors:
  probabilisticsampler:
    mode: equalizing
	sampling_percentage: 10
```

The "equalizing" mode is recommended, see that [component's
documentation](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/processor/probabilisticsamplerprocessor/README.md#equalizing).
