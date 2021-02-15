package helpers

import (
	runtimeClient "github.com/go-openapi/runtime/client"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/client"
)

type (
	Host     string
	Username string
	Password string
)

func ProvideRuntimeClient(host Host, username Username, password Password) (*runtimeClient.Runtime, error) {
	runtime := runtimeClient.New(
		string(host),
		"/",
		[]string{"https"})
	runtime.DefaultAuthentication = runtimeClient.BasicAuth(string(username), string(password))

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

func ProvideAdressLookupClient(runtime *runtimeClient.Runtime) *client.AddressLookup {
	addressLookup := client.New(
		runtime,
		nil)
	return addressLookup
}
