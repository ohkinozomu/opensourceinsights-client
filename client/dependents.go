package client

import (
	"encoding/json"
	"fmt"
)

const (
	dependentsURL = baseURL + "s/%s/p/%s/v/%s/dependents"
)

type DependentsResponse struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Version       string `json:"version"`
	TotalCount    int    `json:"totalCount"`
	DirectCount   int    `json:"directCount"`
	IndirectCount int    `json:"indirectCount"`
	DirectSample  []struct {
		Package struct {
			System string `json:"system"`
			Name   string `json:"name"`
		} `json:"package"`
		Version string `json:"version"`
	} `json:"directSample"`
	IndirectSample []interface{} `json:"indirectSample"`
}

func (c *client) Dependents() (DependentsResponse, error) {
	var r DependentsResponse
	res, err := get(fmt.Sprintf(dependentsURL, c.system, c.p, c.version))
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return r, err
	}
	return r, nil
}
