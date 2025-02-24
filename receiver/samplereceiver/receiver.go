package samplereceiver

import (
	"context"

	"github.com/Arthur1/opentelemetry-collector-components-tutorial/receiver/samplereceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.opentelemetry.io/collector/scraper"
)

func createMetricsReceiver(
	_ context.Context,
	settings receiver.Settings,
	rawCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	cfg := rawCfg.(*config)

	s := newSampleScraper(cfg, settings)
	ms, err := scraper.NewMetrics(s.scrape)
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(
		&cfg.ControllerConfig,
		settings,
		consumer,
		scraperhelper.AddScraper(metadata.Type, ms),
	)
}

var _ receiver.CreateMetricsFunc = createMetricsReceiver
