package publisher

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type PublishRoutesConfig struct {
	Interval    time.Duration `default:"10s"`
	Entrypoints []string      `required:"true"`
}

func loadConfig() (*PublishRoutesConfig, error) {
	var config PublishRoutesConfig
	if err := envconfig.Process("publish_routes", &config); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &config, nil
}
