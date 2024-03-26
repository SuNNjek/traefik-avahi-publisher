package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"traefik-avahi-helper/publisher"
)

// Listen to SIGINT (Ctrl+C) and SIGTERM (docker stop) signals
var cancelSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), cancelSignals...)
	defer cancel()

	cmd, cleanup, err := publisher.CreatePublisher()
	if err != nil {
		log.Fatalln(err)
	}

	defer cleanup()

	if err := cmd.Run(ctx); errors.Is(err, context.Canceled) {
		log.Println("Shutdown requested, shutting down...")
	} else if err != nil {
		log.Fatalln(err)
	}
}
