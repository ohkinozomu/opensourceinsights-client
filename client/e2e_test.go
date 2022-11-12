package client

import (
	"testing"
)

func TestCount(t *testing.T) {
	clientConfig := Config{
		System: "go",
		P:      "github.com/guseggert/pkggodev-client",
	}
	c, err := New(clientConfig)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Describe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Versions()
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Dependencies()
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Dependents()
	if err != nil {
		t.Fatal(err)
	}
}
