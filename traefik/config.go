package traefik

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type TraefikConfig struct {
	Url string `required:"true"`
}

func loadConfig() (*TraefikConfig, error) {
	var config TraefikConfig
	if err := envconfig.Process("traefik", &config); err != nil {
		return nil, fmt.Errorf("failed to load traefik config: %w", err)
	}

	return &config, nil
}
