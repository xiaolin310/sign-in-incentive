// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"sign-in/app/reward/job/internal/conf"
	"sign-in/app/reward/job/internal/data"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*data.RewardRepo, func(), error) {
	panic(wire.Build(data.ProviderSet))
}
