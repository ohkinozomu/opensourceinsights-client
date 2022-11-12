package client

import (
	"encoding/json"
	"fmt"
)

const dependenciesURL = baseURL + "s/%s/p/%s/v/%s/dependencies"

type DependenciesResponse struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Version         string `json:"version"`
	DependencyCount int    `json:"dependencyCount"`
	DirectCount     int    `json:"directCount"`
	IndirectCount   int    `json:"indirectCount"`
	Dependencies    []struct {
		Package struct {
			System string `json:"system"`
			Name   string `json:"name"`
		} `json:"package"`
		Version         string        `json:"version"`
		Type            string        `json:"type"`
		Owners          []interface{} `json:"owners"`
		Licenses        []string      `json:"licenses"`
		Advisories      []interface{} `json:"advisories"`
		Distance        int           `json:"distance"`
		DependencyCount int           `json:"dependencyCount"`
		Errors          []interface{} `json:"errors"`
	} `json:"dependencies"`
}

func (c *client) Dependencies() (DependenciesResponse, error) {
	var r DependenciesResponse
	res, err := get(fmt.Sprintf(dependenciesURL, c.system, c.p, c.version))
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return r, err
	}
	return r, nil
}
