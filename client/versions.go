package client

import (
	"encoding/json"
	"fmt"
)

const versionsURL = baseURL + "s/%s/p/%s/versions"

type VersionsResponse struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Versions []struct {
		Version          string        `json:"version"`
		SymbolicVersions []interface{} `json:"symbolicVersions"`
		IsDefault        bool          `json:"isDefault,omitempty"`
		DependentCount   int           `json:"dependentCount"`
	} `json:"versions"`
}

func (c *client) Versions() (VersionsResponse, error) {
	var r VersionsResponse
	res, err := get(fmt.Sprintf(versionsURL, c.system, c.p))
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return r, err
	}
	return r, nil
}
