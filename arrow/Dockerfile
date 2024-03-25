FROM golang:1.21 AS sandbox

WORKDIR /otelarrowcol
COPY . .
ENV CGO_ENABLED=0

# Note the version MUST MATCH otelarrowcol-build.yaml
# Future optimization - curl the release.
RUN go install go.opentelemetry.io/collector/cmd/builder@v0.96.0

# This command generates main.go, go.mod but does not update deps.
RUN builder --config=/otelarrowcol/otelcolarrow-build.yaml

# This build uses an Alpine Linux container.
FROM alpine AS release
COPY --from=sandbox /otelarrowcol/dist/otelarrowcol /otelarrowcol

# Network ports
# 4317 - OpenTelemetry gRPC services:
#      - OpenTelemetry Protocol with Apache Arrow
#      - OpenTelemetry Protocol (OTLP)
# 1777 - Profiling support
EXPOSE 4317/tcp 1777/tcp

ENTRYPOINT ["/otelarrowcol"]
