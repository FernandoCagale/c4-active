//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-active/api/routers"
	"github.com/FernandoCagale/c4-active/pkg"
	"github.com/google/wire"
)

func SetupApplication() (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}
