package config

import (
	"reflect"
	"testing"
)

func TestConfigParsing(t *testing.T) {
	expected := &Config{
		TestCases: []TestCase{
			TestCase{
				Name: "test1",
				CDN:  "akamai",
				URL: URL{
					Protocol: "https",
					Host:     "www.example.com",
					Path:     "/",
					Headers: []Header{
						Header{
							Key:   "test1",
							Value: "testV",
						},
						Header{
							Key:   "User-Agent",
							Value: "iPhone",
						},
					},
					Cookies: []Cookie{
						Cookie{
							Key:   "cookie1",
							Value: "cookieV",
						},
					},
				},
				Assert: Assert{
					StatusCode: 200,
					Cachable:   true,
					TTL:        "1h",
					Host:       "www.example.com",
				},
			},
		},
	}

	config, err := ParseConfig(testConfig)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, expected) {
		t.Error("Config structure differed from expectation")
	}
}

const testConfig = `testcase "test1" {
	cdn = "akamai"
	url = {
		protocol = "https"
		host = "www.example.com"
		path = "/"
		header = [
			{
				key = "test1"
				value = "testV"
			},
			{
				key = "User-Agent"
				value = "iPhone"
			}
		]
		cookie = [
			{
				key = "cookie1"
				value = "cookieV"
			}
		]
	}
	assert = {
		statuscode = 200
		cachable = true
		ttl = "1h"
		host = "www.example.com"
	}
}`
