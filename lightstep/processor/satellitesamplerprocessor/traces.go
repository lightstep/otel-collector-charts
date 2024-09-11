// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package satellitesamplerprocessor

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.uber.org/zap"

	"github.com/lightstep/otel-collector-charts/lightstep/processor/satellitesamplerprocessor/internal/otelsampling"
	"github.com/lightstep/otel-collector-charts/lightstep/processor/satellitesamplerprocessor/internal/sampler"
)

type tracesProcessor struct {
	logger           *zap.Logger
	traceSampler     sampler.TraceSampler
	singleTraceState string
	staticThreshold  otelsampling.Threshold
	next             consumer.Traces
}

func createTracesProcessor(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Traces) (*tracesProcessor, error) {
	oCfg := cfg.(*Config)

	// Synthesize a trace-state for use when no incoming tracestate is present.
	th, err := otelsampling.ProbabilityToThreshold(oCfg.Percent / 100.0)
	if err != nil {
		return nil, err
	}
	w3c, _ := otelsampling.NewW3CTraceState("")
	if err := w3c.CloudObsValue().UpdateSValueWithSampling(th); err != nil {
		return nil, err
	}
	var w strings.Builder
	if err := w3c.Serialize(&w); err != nil {
		return nil, err
	}

	return &tracesProcessor{
		logger:           set.Logger,
		traceSampler:     sampler.NewTraceSampler(oCfg.Percent),
		singleTraceState: w.String(),
		staticThreshold:  th,
		next:             nextConsumer,
	}, nil
}

func (p *tracesProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	td.ResourceSpans().RemoveIf(func(rs ptrace.ResourceSpans) bool {
		rs.ScopeSpans().RemoveIf(func(ils ptrace.ScopeSpans) bool {
			ils.Spans().RemoveIf(func(s ptrace.Span) bool {
				// The IsSampledOut() method takes a string.  What kind of string?
				// The expression is (annotate_grpc.go):
				//
				//   wire.TranslateIDToGUID(span.SpanContext.TraceId)
				//
				// which is a 16-byte zero-padded hex encoding capturing the least
				// significant 64 bits, i.e., the TraceId above calculated using
				//
				//   if len(id) > 8 {
				//     return binary.BigEndian.Uint64(id[len(id)-8:])
				//   }
				//   return binary.BigEndian.Uint64(id)
				//
				// where, in this case (OTLP ingest) we use the expression:
				//
				//   wire.TranslateBytesToID(traceID)
				//
				// However, note that this pipeline will reject incorrect length,
				// so length is always 16.  We take the second 8 bytes as a hex
				// string, zero pad it, then call the sampler.
				tid := s.TraceID()
				if p.traceSampler.IsSampledOut(fmt.Sprintf("%016x", tid[8:])) {
					return true
				}

				tsIn := s.TraceState().AsRaw()
				if tsIn == "" {
					// In case there is no incoming tracestate, the output is the
					// static encoding resulting from just one sampling stage.
					s.TraceState().FromRaw(p.singleTraceState)
					return false
				}
				// Combination logic is needed.  In case the SDK has sampled
				// using an OTel spec (legacy or modern), the existing fields
				// are preserved.  This sampler uses a dedicated tracestate
				// field to convey the satellite sampling threshold.
				incoming, err := otelsampling.NewW3CTraceState(tsIn)
				if err != nil {
					p.logger.Error("W3C tracecontext: invalid incoming tracestate", zap.Error(err))
					return false
				}
				if err := incoming.CloudObsValue().UpdateSValueWithSampling(p.staticThreshold); err != nil {
					p.logger.Error("W3C tracecontext: arriving trace state", zap.Error(err))
				}

				// Serialize and update.
				var out strings.Builder
				if err := incoming.Serialize(&out); err != nil {
					p.logger.Error("W3C tracecontext: serialize tracestate", zap.Error(err))
				}
				s.TraceState().FromRaw(out.String())
				return false
			})
			// Filter out empty ScopeMetrics
			return ils.Spans().Len() == 0
		})
		// Filter out empty ResourceMetrics
		return rs.ScopeSpans().Len() == 0
	})
	// Maybe skip the data item entirely.
	if td.ResourceSpans().Len() == 0 {
		return td, processorhelper.ErrSkipProcessingData
	}
	return td, nil
}
