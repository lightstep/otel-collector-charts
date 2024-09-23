// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package satellitesamplerprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	typeStr   = "satellitesampler"
	stability = component.StabilityLevelStable
)

func NewFactory() processor.Factory {
	return processor.NewFactory(
		component.MustNewType(typeStr),
		createDefaultConfig,
		processor.WithTraces(createTracesProcessorHelper, stability),
	)
}

func createTracesProcessorHelper(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	tp, err := createTracesProcessor(ctx, set, cfg, nextConsumer)
	if err != nil {
		return nil, err
	}
	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		tp.processTraces,
		processorhelper.WithCapabilities(consumer.Capabilities{MutatesData: true}))
}

func createDefaultConfig() component.Config {
	return &Config{
		Percent: 100,
	}
}
