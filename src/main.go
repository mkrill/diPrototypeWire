package main

import (
	"fmt"
	"os"

	"github.com/google/wire"

	"github.com/mkrill/diPrototypeWire/src/controller"
	"github.com/mkrill/diPrototypeWire/src/infrastructure"
	"github.com/mkrill/diPrototypeWire/src/providers"
	"github.com/mkrill/diPrototypeWire/src/services"
)

var (
	realControllerSet = wire.NewSet(
		providers.ProvideRuntimeClient,
		providers.ProvideAdressLookupClient,
		infrastructure.ProvideRealAddressService,
		wire.Bind(new(services.AddressService), new(*infrastructure.RealAddressService)),
		controller.ProvideController,
	)
	fakeControllerSet = wire.NewSet(
		infrastructure.ProvideFakeAddressService,
		wire.Bind(new(services.AddressService), new(*infrastructure.FakeAddressService)),
		controller.ProvideController,
	)
)

func main() {
	// set runtimeConfig, usually from config file
	runtimeConfig := providers.RuntimeConfig{
		Host:                  "wiremock-acc-proxy-pacman-01.support.magic.telekom.de",
		Username:              os.Getenv("USERNAME"),
		Password:              os.Getenv("PASSWORD"),
		UseFakeAddressService: false,
	}

	err := createRealController(runtimeConfig)
	if err != nil {
		fmt.Printf("\nRealAdapter: %s\n", err.Error())
	}

	err = createFakeController()
	if err != nil {
		fmt.Printf("\nFakeAdapter: %s\n", err.Error())
	}
}

func createFakeController() error {
	fakeController, err := InitializeFakeController()
	if err != nil {
		return err
	}

	faResult, err := fakeController.AddressService.FullAddressInfo("Landgrabenweg", "151", "53227", "Bonn")
	if err != nil {
		return err
	}
	fmt.Printf("FakeAdapter: %v\n", faResult)
	return nil
}

func createRealController(runtimeConfig providers.RuntimeConfig) error {
	realController, err := InitializeRealController(runtimeConfig)
	if err != nil {
		return err
	}
	raResult, err := realController.AddressService.FullAddressInfo("Landgrabenweg", "151", "53227", "Bonn")
	if err != nil {
		return err
	}
	fmt.Printf("RealAdapter: %v\n", raResult)
	return nil
}
