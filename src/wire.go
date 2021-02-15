//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/mkrill/diPrototypeWire/src/helpers"
	"github.com/mkrill/diPrototypeWire/src/interfaces/controller"
)

func InitializeRealController(host helpers.Host, username helpers.Username, password helpers.Password) (*controller.RootController, error) {
	wire.Build(realControllerSet)
	// value returned just to satisfy the compiler, will be ignored
	return &controller.RootController{}, nil
}

func InitializeFakeController() (*controller.RootController, error) {
	wire.Build(fakeControllerSet)
	// value returned just to satisfy the compiler, will be ignored
	return &controller.RootController{}, nil
}
