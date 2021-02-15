package main

import (
	"fmt"
	"os"

	"github.com/google/wire"

	"github.com/mkrill/diPrototypeWire/src/domain/models"
	"github.com/mkrill/diPrototypeWire/src/domain/services"
	"github.com/mkrill/diPrototypeWire/src/helpers"
	"github.com/mkrill/diPrototypeWire/src/infrastructure"
	"github.com/mkrill/diPrototypeWire/src/interfaces/controller"
)

var (
	realControllerSet = wire.NewSet(
		helpers.ProvideRuntimeClient,
		helpers.ProvideAdressLookupClient,
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

	const configFile = "config.yml"
	var err error

	projectConfig, err := helpers.NewConfig(configFile)
	if err != nil {
		fmt.Printf("\nError reading %s: %s\n", configFile, err.Error())
		os.Exit(2)
	}
	fmt.Printf("\nConfig-File contents\n=> %s\n", projectConfig)

	if projectConfig != nil && projectConfig.Address.UseFakeAddressService {
		faResult, err := createAndExecuteFakeController()
		if err != nil {
			fmt.Printf("\nFakeAdapter-Error: %s\n", err.Error())
			os.Exit(2)
		}
		fmt.Printf("\nFakeAdapter-Result: %v\n", faResult)
	} else if projectConfig != nil && !projectConfig.Address.UseFakeAddressService {
		raResult, err := createAndExecuteRealController()
		if err != nil {
			fmt.Printf("\nRealAdapter-Error: %s\n", err.Error())
			os.Exit(2)
		}
		fmt.Printf("\nRealAdapter-Result: %v\n", raResult)
	}
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

func createAndExecuteRealController() (*models.Address, error) {
	host := helpers.Host(os.Getenv("HOST"))
	username := helpers.Username(os.Getenv("USERNAME"))
	password := helpers.Password(os.Getenv("PASSWORD"))

	realController, err := InitializeRealController(host, username, password)
	if err != nil {
		return nil, err
	}

	raResult, err := realController.AddressService.FullAddressInfo("Ahornweg", "1", "34227", "Heimerdingen")
	if err != nil {
		return nil, err
	}
	return raResult, nil
}
