package samplereceiver

import (
	"github.com/Arthur1/opentelemetry-collector-components-tutorial/receiver/samplereceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	AnswerName                     string `mapstructure:"answerer_name"`
}

func createDefaultConfig() component.Config {
	cc := scraperhelper.NewDefaultControllerConfig()
	mbc := metadata.DefaultMetricsBuilderConfig()
	return &config{
		ControllerConfig:     cc,
		MetricsBuilderConfig: mbc,
		AnswerName:           "",
	}
}

var _ component.CreateDefaultConfigFunc = createDefaultConfig
