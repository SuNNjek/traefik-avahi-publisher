package publisher

import (
	"context"
	"regexp"
	"slices"
	"strings"
	"time"
	"traefik-avahi-helper/avahi"
	"traefik-avahi-helper/log"
	"traefik-avahi-helper/traefik"
)

var (
	localHostsRegex = regexp.MustCompile("Host\\(`([a-zA-Z0-9_.-]+\\.local)`\\)")
)

type Publisher struct {
	config *PublishRoutesConfig
	logger log.Logger

	avahiClient      *avahi.AvahiClient
	traefikApiClient *traefik.TraefikApiClient
}

func NewPublisher(
	config *PublishRoutesConfig,
	logger log.Logger,
	avahiClient *avahi.AvahiClient,
	traefikApiClient *traefik.TraefikApiClient,
) *Publisher {
	return &Publisher{
		config,
		logger,
		avahiClient,
		traefikApiClient,
	}
}

func (cmd *Publisher) Run(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)

	if fqdn, err := cmd.avahiClient.GetHostNameFqdn(); err == nil {
		cmd.logger.Debug("Host FQDN: %s", fqdn)
	} else {
		return err
	}

	for {
		if err := cmd.publishLocalRouters(); err != nil {
			return err
		}

		select {
		case <-ticker.C:
			continue

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func overlap[T comparable](a []T, b []T) bool {
	return slices.ContainsFunc(a, func(itemA T) bool {
		return slices.Contains(b, itemA)
	})
}

func (cmd *Publisher) publishLocalRouters() error {
	routers, err := cmd.traefikApiClient.GetHttpRouters()
	if err != nil {
		return err
	}

	names := make([]string, 0)
	for _, router := range routers {
		// Check if the router has any overlap with the configured entrypoints.
		// If not, then we don't want to publish this router
		if !overlap(cmd.config.Entrypoints, router.EntryPoints) {
			continue
		}

		matches := localHostsRegex.FindAllStringSubmatch(router.Rule, -1)
		if len(matches) == 0 {
			continue
		}

		for _, match := range matches {
			names = append(names, match[1])
		}
	}

	if len(names) == 0 {
		return nil
	}

	cmd.logger.Debug("Publishing %s", strings.Join(names, ", "))
	return cmd.avahiClient.PublishCnames(names)
}
