package providers

import (
	runtimeClient "github.com/go-openapi/runtime/client"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/client"
)

type RuntimeConfig struct {
	Host                  string
	Username              string
	Password              string
	UseFakeAddressService bool
}

func ProvideAdressLookupClient(runtime *runtimeClient.Runtime) *client.AddressLookup {
	addressLookup := client.New(
		runtime,
		nil)
	return addressLookup
}

func ProvideRuntimeClient(config RuntimeConfig) *runtimeClient.Runtime {
	runtime := runtimeClient.New(
		config.Host,
		"/",
		[]string{"https"})
	runtime.DefaultAuthentication = runtimeClient.BasicAuth(config.Username, config.Password)
	return runtime
}
