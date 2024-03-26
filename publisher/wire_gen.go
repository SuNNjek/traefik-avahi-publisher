// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package publisher

import (
	"net/http"
	"traefik-avahi-helper/avahi"
	"traefik-avahi-helper/traefik"
)

// Injectors from wire.go:

func CreatePublisher() (*Publisher, func(), error) {
	publishRoutesConfig, err := loadConfig()
	if err != nil {
		return nil, nil, err
	}
	avahiClient, cleanup, err := avahi.CreateAvahiClient()
	if err != nil {
		return nil, nil, err
	}
	httpClient := _wireClientValue
	traefikApiClient, err := traefik.CreateApiClient(httpClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	publisher := NewPublisher(publishRoutesConfig, avahiClient, traefikApiClient)
	return publisher, func() {
		cleanup()
	}, nil
}

var (
	_wireClientValue = http.DefaultClient
)
