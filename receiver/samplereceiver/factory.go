package samplereceiver

import (
	"github.com/Arthur1/opentelemetry-collector-components-tutorial/receiver/samplereceiver/internal/metadata"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(
			createMetricsReceiver,
			metadata.MetricsStability,
		),
	)
}
