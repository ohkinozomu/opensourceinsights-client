package client

import (
	"encoding/json"
	"fmt"
)

const dependenciesWithGraphURL = baseURL + "s/%s/p/%s/v/%s/dependenciesWithGraph"

type DependenciesWithGraphResponse struct {
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
		Errors          []struct {
			Requirement struct {
				Package struct {
					System string `json:"system"`
					Name   string `json:"name"`
				} `json:"package"`
				Version string `json:"version"`
			} `json:"requirement"`
			Error string `json:"error"`
		} `json:"errors"`
	} `json:"dependencies"`
	DependencyGraph struct {
		Nodes []struct {
			Package struct {
				System string `json:"system"`
				Name   string `json:"name"`
			} `json:"package"`
			Version string `json:"version"`
			NodeID  int    `json:"nodeID"`
		} `json:"nodes"`
		Edges []struct {
			From        int    `json:"from"`
			To          int    `json:"to"`
			Requirement string `json:"requirement"`
			Type        string `json:"type"`
		} `json:"edges"`
	} `json:"dependencyGraph"`
}

func (c *client) DependenciesWithGraph() (DependenciesWithGraphResponse, error) {
	var r DependenciesWithGraphResponse
	res, err := get(fmt.Sprintf(dependenciesWithGraphURL, c.system, c.p, c.version))
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return r, err
	}
	return r, nil
}
