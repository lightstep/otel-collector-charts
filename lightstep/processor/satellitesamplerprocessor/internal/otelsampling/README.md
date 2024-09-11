The code in this directory is is a copy of the OpenTelemetry package
in collector-contrib/pkg/sampling. The copy here has has ServiceNow
copyright because it was originally authored here.

Code organization:

# Tracestate handling

- w3ctracestate.go: the outer tracestate structure like `key=value,...`
- oteltracestate.go: the inner tracestate structure like `key:value;...`
- cloudobstracestate.go: the inner tracestate structure like `key:value;...` (internal only)
- common.go: shared parser, serializer for either tracestate

This includes an implementation of the W3C trace randomness feature,
described here: https://www.w3.org/TR/trace-context-2/#randomness-of-trace-id

# Overview of tracestate identifiers

There are two vendor codes:

- "ot" refers to formal OTel specifications
- "sn" for ServiceNow refers to internal Cloud Observability sampling (as by the Lightstep satellite)

The OTel trace state keys:

- "p" refers to the [legacy OTel power-of-two sampling](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/tracestate-probability-sampling.md)
- "r" used in the legacy convention, 1-2 decimal digits
- "th" refers to the [modern OTel 56-bit sampling](https://github.com/open-telemetry/oteps/pull/235)
- "rv" refers to the modern randomness value, 14 hex digits.

The Cloud Observability trace state keys:

- "s" refers to the satellite sampling, uses the same encoding as "th" but is modeled as an acceptance threshold. 

Note that to convert from an OTel rejection threshold to a Satellite sampler acceptance threshold, the unsigned value of the threshold should be subtracted from the maximum adjusted count,

```
satelliteSamplerThreshold, _ = UnsignedToThreshold(MaxAdjustedCount - otelModernThreshold.Unsigned())
```

# Encoding and decoding

- probability.go: defines
  `ProbabilityToThreshold()`
  `(Threshold).Probability()`
- threshold.go: defines
  `TValueToThreshold()`
  `(Threshold).TValue()`
  `(Threshold).ShouldSample()`
- randomness.go: defines 
  `TraceIDToRandomness()`
  `RValueToRandomness()`
  `(Randomness).RValue()`

# Callers of note

- In common-go/wire/oteltoegresspb/otel_to_egresspb.go:
  `TraceStateToAdjustedCount()`
   
- In internalcollector/components/satellitesamplerprocessor/traces.go:
  `createTracesProcessor()`
