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
	// Name is required
	Name string `hcl:",key"`
	// CDN supported "akamai" only
	CDN string `hcl:"cdn"`
	// URL is required
	URL URL `hcl:"url"`
	// Assert is required
	Assert Assert `hcl:"assert"`
}

// URL is reuqired for testing
type URL struct {
	// Protocol is http or https
	Protocol string `hcl:"protocol"`
	// Host is a domain
	Host string `hcl:"host"`
	// Method is not required, default is get Method
	Method string `hcl:"method"`
	// path is not required, "" should be ok
	Path string `hcl:"path"`
	// header is not required
	Headers []Header `hcl:"header"`
	// cookie is not required
	Cookies []Cookie `hcl:"cookie"`
}

// Header is request header, for example "User-Agent"
type Header struct {
	// if Header is not nil then Key is required
	Key string `hcl:"key"`
	// if Header is not nil then Value is required
	Value string `hcl:"value"`
}

// Cookie is request cookie, should be baked for another request
type Cookie struct {
	// if Cookie is not nil then Key is required
	Key string `hcl:"key"`
	// if Value is not nil then Value is required
	Value string `hcl:"value"`
}

// Assert is required for testing
type Assert struct {
	// StatusCode is required for testing
	StatusCode int `hcl:"statuscode"`
	//Akamai assert
	// Host is not host header, it search in "x-cache-key"
	Host string `hcl:"host"`
	// Cachable is check response header "x-check-cacheable"
	Cachable bool `hcl:"cachable"`
	// TTL is check response header in "x-cheche-key"
	TTL string `hcl:"ttl"`
}

// ParseConfig parse the given HCL string into a Config struct.
func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}
	var errors *multierror.Error

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, multierror.Append(errors, err)
	}

	if err := hcl.DecodeObject(&result, hclParseTree); err != nil {
		return nil, multierror.Append(errors, err)
	}

	log.Printf("%+v\n", result)

	return result, errors.ErrorOrNil()
}

// ValidateConfig is to validate config struct
func (config *Config) ValidateConfig() error {
	var errors *multierror.Error
	return errors.ErrorOrNil()
}
