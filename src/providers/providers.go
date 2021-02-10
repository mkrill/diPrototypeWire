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

func ProvideRuntimeClient(config RuntimeConfig) (*runtimeClient.Runtime, error) {
	runtime := runtimeClient.New(
		config.Host,
		"/",
		[]string{"https"})
	runtime.DefaultAuthentication = runtimeClient.BasicAuth(config.Username, config.Password)

	// switch of verification of server side certificate (ok for prototype)
	tlsTransport, err := runtimeClient.TLSTransport(runtimeClient.TLSClientOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	runtime.Transport = tlsTransport

	return runtime, nil
}
