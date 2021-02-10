package controller

import (
	"errors"
	"time"

	"github.com/mkrill/diPrototypeWire/src/domain/services"
)

type RootController struct {
	AddressService services.AddressService
}

func ProvideController(as services.AddressService) (*RootController, error) {
	if time.Now().Unix()%3 == 0 {
		return nil, errors.New("time can be divided by 3, cannot create controller")
	}
	return &RootController{AddressService: as}, nil
}
