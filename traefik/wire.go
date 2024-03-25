//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package traefik

import (
	"traefik-avahi-helper/util"

	"github.com/google/wire"
)

func CreateApiClient(httpClient util.HttpClient) (*TraefikApiClient, error) {
	wire.Build(loadConfig, NewApiClient)

	return &TraefikApiClient{}, nil
}
