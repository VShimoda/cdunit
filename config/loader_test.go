package config

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expected := `testcase "test1" {
	cdn = "akamai"
	url = {
		protocol = "https"
		host = "developer.akamai.com"
		path = "/"
		header = [
			{
				key = "User-Agent"
				value = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Mobile Safari/537.36"
			}
		]
	}
	assert = {
		statuscode = 200
		cachable = true
		ttl = "1h"
		host = "opencomm.download.akamai.com"
	}
}


testcase "test2" {
	cdn = "akamai"
	url = {
		protocol = "https"
		host = "developer.akamai.com"
		path = "/"
		header = [
			{
				key = "User-Agent"
				value = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
			}
		]
	}
	assert = {
		statuscode = 200
		cachable = true
		ttl = "1h"
		host = "opencomm.download.akamai.com"
	}
}

`
	if err := os.Chdir("../sample"); err != nil {
		t.Error("Filed to change directory to sample")
	}
	hclString, err := LoadConfig()
	if err != nil {
		t.Error("Failed to read config file")
	}
	if hclString != expected {
		fmt.Println(hclString)
		t.Error("Failed to read file")
	}
	fmt.Println(hclString)
}
