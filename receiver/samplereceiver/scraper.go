package samplereceiver

import (
	"context"
	"time"

	"github.com/Arthur1/opentelemetry-collector-components-tutorial/receiver/samplereceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type sampleScraper struct {
	cfg *config
	mb  *metadata.MetricsBuilder
}

func (s *sampleScraper) scrape(context.Context) (pmetric.Metrics, error) {
	ts := pcommon.NewTimestampFromTime(time.Now())
	s.mb.RecordSampleUltimateAnswerDataPoint(ts, 42, s.cfg.AnswerName)
	return s.mb.Emit(), nil
}

func newSampleScraper(cfg *config, settings receiver.Settings) *sampleScraper {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	return &sampleScraper{
		cfg: cfg,
		mb:  mb,
	}
}
