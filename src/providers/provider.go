package providers

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ghodss/yaml"
	runtimeClient "github.com/go-openapi/runtime/client"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/client"
)

type (
	RuntimeConfig struct {
		Host                  string `yaml:"host"`
		Username              string `yaml:"username"`
		Password              string `yaml:"password"`
		UseFakeAddressService bool   `yaml:"useFakeAddressService"`
	}

	Config struct {
		Address RuntimeConfig `yaml:"address"`
	}
)

func ProvideConfig(filename string) (*Config, error) {
	var (
		fileContent []byte
		err         error
	)

	if fileContent, err = ioutil.ReadFile(path.Join("config", filename)); err != nil {
		return nil, err
	}

	config := &Config{}
	if err = yaml.Unmarshal(fileContent, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) String() string {
	return fmt.Sprintf(
		"Host: %s, Username: %s, Password: %s, UseFakeAddressService: %v",
		c.Address.Host,
		c.Address.Username,
		c.Address.Password,
		c.Address.UseFakeAddressService,
	)
}

func ProvideAdressLookupClient(runtime *runtimeClient.Runtime) *client.AddressLookup {
	addressLookup := client.New(
		runtime,
		nil)
	return addressLookup
}

func ProvideRuntimeClient(config RuntimeConfig) (*runtimeClient.Runtime, error) {
	runtime := runtimeClient.New(
		config.Host,
		"/",
		[]string{"https"})
	runtime.DefaultAuthentication = runtimeClient.BasicAuth(config.Username, config.Password)

	// configure ca certificate to be trusted, set InsecureSkipVerify to switch off certificate check
	tlsTransport, err := runtimeClient.TLSTransport(runtimeClient.TLSClientOptions{
		CA:                 "certificates/ca_pub_certificate.pem",
		InsecureSkipVerify: false,
	})
	if err != nil {
		return nil, err
	}
	runtime.Transport = tlsTransport

	return runtime, nil
}
