//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/xiaoyan648/project-layout/internal/data"
	"github.com/xiaoyan648/project-layout/internal/model"
	"github.com/xiaoyan648/project-layout/internal/pkg/conf"
	"github.com/xiaoyan648/project-layout/internal/service"

	"github.com/go-leo/leo"
	"github.com/go-leo/leo/log"
	"github.com/google/wire"
)

// wireApp init leo application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*leo.App, func(), error) {
	panic(wire.Build(data.ProviderSet, model.ProviderSet, service.ProviderSet, newApp))
}
