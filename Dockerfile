FROM golang:1.22 AS build

WORKDIR /build

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Use CGO_ENABLED so that the binary gets built statically
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o traefik-avahi-publisher


# Use distroless base image. This already contains SSL certificates and timezone data
FROM gcr.io/distroless/static

COPY --from=build /build/traefik-avahi-publisher /bin/traefik-avahi-publisher

ENTRYPOINT [ "/bin/traefik-avahi-publisher" ]