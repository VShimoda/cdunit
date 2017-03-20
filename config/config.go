package config

import (
	"log"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
)

// Config type
type Config struct {
	TestCases []TestCase `hcl:"testcase"`
}

// TestCase type
type TestCase struct {
	Name   string `hcl:",key"`
	CDN    string `hcl:"cdn"`
	URL    URL    `hcl:"url"`
	Assert Assert `hcl:"assert"`
}

type URL struct {
	Protocol string   `hcl:"protocol"`
	Host     string   `hcl:"host"`
	Path     string   `hcl:"path"`
	Headers  []Header `hcl:"header"`
	Cookies  []Cookie `hcl:"cookie"`
}

type Header struct {
	Key   string `hcl:"key"`
	Value string `hcl:"value"`
}

type Cookie struct {
	Key   string `hcl:"key"`
	Value string `hcl:"value"`
}

type Assert struct {
	StatusCode int    `hcl:"statuscode"`
	Host       string `hcl:"host"`
	Cachable   bool   `hcl:"cachable"`
	TTL        string `hcl:"ttl"`
}

// ParseConfig parse the given HCL string into a Config struct.
func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}
	var errors *multierror.Error

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(&result, hclParseTree); err != nil {
		return nil, err
	}

	log.Printf("%+v\n", result)

	return result, errors.ErrorOrNil()
}

