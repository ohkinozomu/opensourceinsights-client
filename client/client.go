package client

import (
	"net/url"
)

const (
	baseURL = "https://deps.dev/_/"
)

type client struct {
	system string
	// package name
	p       string
	version string
}

type Config struct {
	System string
	// package name
	P       string
	Version string
}

func getDefaultVersion(config Config) (string, error) {
	res, err := vRequest(config.System, config.P)
	if err != nil {
		return "", err
	}
	return res.DefaultVersion, nil
}

func New(config Config) (*client, error) {
	version := config.Version
	config.P = url.QueryEscape(config.P)
	if version == "" {
		defaultVersion, err := getDefaultVersion(config)
		if err != nil {
			return nil, err
		}
		version = defaultVersion
	}
	c := client{
		system:  config.System,
		p:       config.P,
		version: version,
	}
	return &c, nil
}
