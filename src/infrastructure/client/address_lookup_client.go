// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/address_deprecated_controller"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/address_search"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/address_suggest"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/address_suggest_v2"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/basic_error_controller"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/geoposition_services"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/zipcode_municipality_name_list"
)

// Default address lookup HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "address-lookup-proxy-presales-develop-01.support.magic.telekom.de"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new address lookup HTTP client.
func NewHTTPClient(formats strfmt.Registry) *AddressLookup {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new address lookup HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *AddressLookup {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new address lookup client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *AddressLookup {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(AddressLookup)
	cli.Transport = transport
	cli.AddressDeprecatedController = address_deprecated_controller.New(transport, formats)
	cli.AddressSearch = address_search.New(transport, formats)
	cli.AddressSuggest = address_suggest.New(transport, formats)
	cli.AddressSuggestV2 = address_suggest_v2.New(transport, formats)
	cli.BasicErrorController = basic_error_controller.New(transport, formats)
	cli.GeopositionServices = geoposition_services.New(transport, formats)
	cli.ZipcodeMunicipalityNameList = zipcode_municipality_name_list.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// AddressLookup is a client for address lookup
type AddressLookup struct {
	AddressDeprecatedController address_deprecated_controller.ClientService

	AddressSearch address_search.ClientService

	AddressSuggest address_suggest.ClientService

	AddressSuggestV2 address_suggest_v2.ClientService

	BasicErrorController basic_error_controller.ClientService

	GeopositionServices geoposition_services.ClientService

	ZipcodeMunicipalityNameList zipcode_municipality_name_list.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *AddressLookup) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AddressDeprecatedController.SetTransport(transport)
	c.AddressSearch.SetTransport(transport)
	c.AddressSuggest.SetTransport(transport)
	c.AddressSuggestV2.SetTransport(transport)
	c.BasicErrorController.SetTransport(transport)
	c.GeopositionServices.SetTransport(transport)
	c.ZipcodeMunicipalityNameList.SetTransport(transport)
}