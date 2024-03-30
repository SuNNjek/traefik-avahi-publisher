# traefik-avahi-publisher

[![CI Workflow Status](https://img.shields.io/github/actions/workflow/status/SuNNjek/traefik-avahi-publisher/ci.yml?branch=main&label=CI&logo=github)](https://github.com/SuNNjek/traefik-avahi-publisher/actions/workflows/ci.yml)
[![Image Workflow Status](https://img.shields.io/github/actions/workflow/status/SuNNjek/traefik-avahi-publisher/docker.yml?branch=main&label=Image%20Build&logo=github)](https://github.com/SuNNjek/traefik-avahi-publisher/actions/workflows/docker.yml)
[![License](https://img.shields.io/github/license/SuNNjek/traefik-avahi-publisher)](https://github.com/SuNNjek/traefik-avahi-publisher/blob/main/LICENSE.txt)

Small helper program that continually queries the Traefik API for HTTP routes that respond to `.local` domains
and publishes them as mDNS CNAME records with Avahi.

## Usage

The recommended usage is with Docker Compose:

```yml
services:
  traefik:
    container_name: traefik
    image: traefik
    restart: unless-stopped
    ports:
      - "80:80"
    environment:
      TZ: Europe/Berlin
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro

  avahi-publisher:
    container_name: avahi-publisher
    image: ghcr.io/sunnjek/traefik-avahi-publisher:edge
    restart: unless-stopped
    depends_on:
      - traefik
    environment:
      TZ: Europe/Berlin
      TRAEFIK_URL: http://traefik:8080/
      PUBLISH_ROUTES_ENTRYPOINTS: web
    volumes:
      # Needed to communicate with the host Avahi daemon via DBus
      - /run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket
```

## Configuration

You can configure the container with the following environment variables:

| Variable                   | Description                                             |
| -------------------------- | ------------------------------------------------------- |
| `TRAEFIK_URL`              | The base URL of the Traefik server                      |
| `PUBLISH_ROUTES_ENDPOINTS` | Comma separated list of endpoints to publish            |
| `PUBLISH_ROUTES_INTERVAL`  | How often to query the Traefik API (default 10 seconds) |
| `LOG_LEVEL`                | Controls the log level (default info)                   |