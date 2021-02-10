package main

import (
	"fmt"
	"os"

	"github.com/google/wire"

	"github.com/mkrill/diPrototypeWire/src/domain/models"
	"github.com/mkrill/diPrototypeWire/src/domain/services"
	"github.com/mkrill/diPrototypeWire/src/infrastructure"
	"github.com/mkrill/diPrototypeWire/src/interfaces/controller"
	"github.com/mkrill/diPrototypeWire/src/providers"
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

	var err error

	config, err := providers.ProvideConfig("config.yml")
	if err != nil {
		fmt.Printf("\nConfig: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Printf("\nConfig-File contents\n=> %s\n", config)

	// set runtimeConfig, usually from config file
	runtimeConfig := providers.RuntimeConfig{
		Host:                  os.Getenv("HOST"),
		Username:              os.Getenv("USERNAME"),
		Password:              os.Getenv("PASSWORD"),
		UseFakeAddressService: false,
	}

	raResult, err := createAndExecuteRealController(runtimeConfig)
	if err != nil {
		fmt.Printf("\nRealAdapter-Error: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Printf("\nRealAdapter-Result: %v\n", raResult)

	faResult, err := createAndExecuteFakeController()
	if err != nil {
		fmt.Printf("\nFakeAdapter-Error: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Printf("\nFakeAdapter-Result: %v\n", faResult)
}

func createAndExecuteFakeController() (*models.Address, error) {
	fakeController, err := InitializeFakeController()
	if err != nil {
		return nil, err
	}

	faResult, err := fakeController.AddressService.FullAddressInfo("Landgrabenweg", "151", "53227", "Bonn")
	if err != nil {
		return nil, err
	}
	return faResult, nil
}

func createAndExecuteRealController(runtimeConfig providers.RuntimeConfig) (*models.Address, error) {
	realController, err := InitializeRealController(runtimeConfig)
	if err != nil {
		return nil, err
	}
	raResult, err := realController.AddressService.FullAddressInfo("Landgrabenweg", "151", "53227", "Bonn")
	if err != nil {
		return nil, err
	}
	return raResult, nil
}
