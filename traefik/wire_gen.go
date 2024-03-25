// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package traefik

import (
	"traefik-avahi-helper/util"
)

// Injectors from wire.go:

func CreateApiClient(httpClient util.HttpClient) (*TraefikApiClient, error) {
	traefikConfig, err := loadConfig()
	if err != nil {
		return nil, err
	}
	traefikApiClient := NewApiClient(traefikConfig, httpClient)
	return traefikApiClient, nil
}
