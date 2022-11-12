# opensourceinsights-client

Unofficial [Open Source Insights](https://deps.dev/) client library

This project is not officially supported or endorsed by Google in any way.

## Example

```go
package main

import (
	"fmt"

	"github.com/ohkinozomu/opensourceinsights-client/client"
)

func main() {
	clientConfig := client.Config{
		System: "go",
		P:      "github.com/guseggert/pkggodev-client",
	}
	c, err := client.New(clientConfig)
	if err != nil {
		panic(err)
	}
	describeResponse, err := c.Describe()
	if err != nil {
		panic(err)
	}
	fmt.Println(describeResponse)
	versionsResponse, err := c.Versions()
	if err != nil {
		panic(err)
	}
	fmt.Println(versionsResponse)
	dependenciesReponse, err := c.Dependencies()
	if err != nil {
		panic(err)
	}
	fmt.Println(dependenciesReponse)
	dependentsResponse, err := c.Dependents()
	if err != nil {
		panic(err)
	}
	fmt.Println(dependentsResponse)
}

```