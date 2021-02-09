// package application contains the implementation of all backend or databases called by the domain "address"
// => out calls from module address
package infrastructure

import (
	"errors"

	"github.com/mkrill/diPrototypeWire/src/models"
)

// FakeAddressService has no attributes, because it does not need a backend address or a client to be stored
// in the structure (in contrast to the real implementation of the AddressService)
type FakeAddressService struct{}

func ProvideFakeAddressService() *FakeAddressService {
	return &FakeAddressService{}
}

// (f *FakeAddressService) FullAddressInfo returns an error, when being called with streetNr "0"
// otherwise it returns the input address parameters
func (f *FakeAddressService) FullAddressInfo(street, streetNr, postCode, city string) (
	*models.Address,
	error,
) {
	// return error, if streetNr is 0
	if streetNr == "0" {
		return nil, errors.New("housenumber invalid")
	}

	// return data passed in FullAddressInfo call, set KlsId to streetNr
	return &models.Address{
		Street:   street,
		StreetNr: streetNr,
		PostCode: postCode,
		City:     city,
		KlsId:    streetNr,
	}, nil
}
