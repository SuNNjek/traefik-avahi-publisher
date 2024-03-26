//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package publisher

import (
	"net/http"
	"traefik-avahi-helper/avahi"
	"traefik-avahi-helper/traefik"
	"traefik-avahi-helper/util"

	"github.com/google/wire"
)

func CreatePublisher() (*Publisher, func(), error) {
	wire.Build(
		wire.InterfaceValue(new(util.HttpClient), http.DefaultClient),
		loadConfig,
		avahi.CreateAvahiClient,
		traefik.CreateApiClient,
		NewPublisher,
	)

	return &Publisher{}, nil, nil
}
