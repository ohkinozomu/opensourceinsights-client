package client

import (
	"encoding/json"
	"fmt"
)

const vURL = baseURL + "s/%s/p/%s/v/"

type DescribeResponse struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Owners  []interface{} `json:"owners"`
	Version struct {
		Version                string        `json:"version"`
		SymbolicVersions       []interface{} `json:"symbolicVersions"`
		RefreshedAt            int           `json:"refreshedAt"`
		IsDefault              bool          `json:"isDefault"`
		Licenses               []string      `json:"licenses"`
		DependentCount         int           `json:"dependentCount"`
		DependentCountDirect   int           `json:"dependentCountDirect"`
		DependentCountIndirect int           `json:"dependentCountIndirect"`
		Links                  struct {
			Origins []string `json:"origins"`
			Repo    string   `json:"repo"`
		} `json:"links"`
		Projects []struct {
			Type       string `json:"type"`
			Name       string `json:"name"`
			ObservedAt int    `json:"observedAt"`
			Issues     int    `json:"issues"`
			Forks      int    `json:"forks"`
			Stars      int    `json:"stars"`
			License    string `json:"license"`
			Link       string `json:"link"`
		} `json:"projects"`
		Advisories      []interface{} `json:"advisories"`
		RelatedPackages struct {
		} `json:"relatedPackages"`
	} `json:"version"`
	DefaultVersion string `json:"defaultVersion"`
}

func vRequest(system, p string) (DescribeResponse, error) {
	var r DescribeResponse
	describeURL := fmt.Sprintf(vURL, system, p)
	res, err := get(describeURL)
	if err != nil {
		return r, err
	}
	if err := json.Unmarshal(res, &r); err != nil {
		return r, err
	}
	return r, nil
}

func (c *client) Describe() (DescribeResponse, error) {
	r, err := vRequest(c.system, c.p)
	if err != nil {
		return r, err
	}
	return r, nil
}
