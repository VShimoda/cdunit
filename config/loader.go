package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
)

// LoadConfig is read all current hcl files
func LoadConfig() (string, error) {
	var errors *multierror.Error
	pwd, err := os.Getwd()
	if err != nil {
		return "", multierror.Append(errors, err)
	}
	fileinfos, err := ioutil.ReadDir(pwd)
	if err != nil {
		return "", multierror.Append(errors, err)
	}
	configFiles := []string{}
	for _, fileinfo := range fileinfos {
		filename := fileinfo.Name()
		pos := strings.LastIndex(filename, ".")
		if filename[pos:] == ".hcl" {
			fmt.Println("read ", filename)
			configFiles = append(configFiles, filename)
		}
	}
	buf := [][]byte{}
	for _, configFile := range configFiles {
		buffer, err := ioutil.ReadFile(filepath.Join(pwd, configFile))
		if err != nil {
			return "", multierror.Append(errors, err)
		}
		buf = append(buf, buffer)
	}

	return string(bytes.Join(buf, []byte("\n"))), errors.ErrorOrNil()
}
