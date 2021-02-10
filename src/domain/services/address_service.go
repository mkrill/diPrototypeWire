package services

import (
	"github.com/mkrill/diPrototypeWire/src/domain/models"
)

// AddressService and its functions need to be exported, because it will be used in methods of root.go controller
// of module "application"
type AddressService interface {
	// FullAddressInfo returns the full address data belonging to the address identified by fullAddressInfoParameter
	FullAddressInfo(street, streetNr, postCode, city string) (*models.Address, error)
}
