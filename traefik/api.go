package traefik

import (
	"encoding/json"
	"net/url"

	"traefik-avahi-helper/util"
)

type TraefikApiClient struct {
	config     *TraefikConfig
	httpClient util.HttpClient
}

func NewApiClient(config *TraefikConfig, httpClient util.HttpClient) *TraefikApiClient {
	return &TraefikApiClient{
		config,
		httpClient,
	}
}

func (c *TraefikApiClient) GetHttpRouters() ([]*Router, error) {
	url, err := combineUrls(c.config.Url, "api/http/routers")
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var res []*Router
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func combineUrls(base, rel string) (string, error) {
	baseUrl, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	relUrl, err := url.Parse(rel)
	if err != nil {
		return "", err
	}

	return baseUrl.ResolveReference(relUrl).String(), nil
}
