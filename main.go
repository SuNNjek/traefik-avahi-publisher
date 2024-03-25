package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"traefik-avahi-helper/avahi"
)

// Listen to SIGINT (Ctrl+C) and SIGTERM (docker stop) signals
var cancelSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), cancelSignals...)
	defer cancel()

	ticker := time.NewTicker(5 * time.Second)

	client, cleanup, err := avahi.CreateAvahiClient()
	if err != nil {
		log.Fatalln(err)
	}

	defer cleanup()

	fqdn, err := client.GetHostNameFqdn()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		err = client.PublishCnames([]string{
			"asdf." + fqdn,
			"test." + fqdn,
		})

		if err != nil {
			log.Fatalln(err)
		}

		select {
		case <-ticker.C:
			continue

		case <-ctx.Done():
			os.Exit(0)
		}
	}
}
