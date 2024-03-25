//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package avahi

import (
	"github.com/godbus/dbus/v5"
	"github.com/google/wire"
)

func provideSystemBus() (*dbus.Conn, func(), error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, nil, err
	}

	return conn, func() { conn.Close() }, nil
}

func provideAvahiClient(conn *dbus.Conn) (*AvahiClient, func(), error) {
	client, err := NewAvahiClient(conn)
	if err != nil {
		return nil, nil, err
	}

	return client, func() { client.Close() }, nil
}

func CreateAvahiClient() (*AvahiClient, func(), error) {
	wire.Build(provideSystemBus, provideAvahiClient)

	return &AvahiClient{}, func() {}, nil
}
