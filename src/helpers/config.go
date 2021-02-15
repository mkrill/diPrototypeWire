package helpers

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ghodss/yaml"
)

type ProjectConfig struct {
	Address struct {
		UseFakeAddressService bool
	} `yaml:"address"`
}

func NewConfig(filename string) (*ProjectConfig, error) {
	var (
		fileContent []byte
		err         error
	)

	if fileContent, err = ioutil.ReadFile(path.Join("config", filename)); err != nil {
		return nil, err
	}

	config := &ProjectConfig{}
	if err = yaml.Unmarshal(fileContent, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *ProjectConfig) String() string {
	return fmt.Sprintf(
		"UseFakeAddressService: %v",
		c.Address.UseFakeAddressService,
	)
}
