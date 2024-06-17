//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/ovo-user/internal/api"
	"github.com/liuzhaomax/ovo-user/internal/core"
	"github.com/liuzhaomax/ovo-user/internal/middleware"
	"github.com/liuzhaomax/ovo-user/internal/middleware_rpc"
	"github.com/liuzhaomax/ovo-user/src/set"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		core.InitLogrus,
		core.InitGinEngine,
		core.InitDB,
		core.InitRedis,
		core.InitTracer,
		core.InitPrometheusRegistry,
		api.APISet,
		api.APIRPCSet,
		set.HandlerSet,
		set.BusinessSet,
		set.ModelSet,
		core.LoggerSet,
		core.ResponseSet,
		core.TransactionSet,
		core.RocketMQSet,
		middleware.MwsSet,
		middleware.MiddlewareSet,
		middleware_rpc.MwsRPCSet,
		middleware_rpc.MiddlewareRPCSet,
		InjectorHTTPSet,
		InjectorRPCSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
