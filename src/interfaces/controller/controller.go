package controller

import (
	"github.com/mkrill/diPrototypeWire/src/domain/services"
)

type RootController struct {
	AddressService services.AddressService
}

func ProvideController(as services.AddressService) (*RootController, error) {
	return &RootController{AddressService: as}, nil
}
