// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/liuzhaomax/ovo-user/internal/api"
	"github.com/liuzhaomax/ovo-user/internal/core"
	"github.com/liuzhaomax/ovo-user/internal/middleware"
	"github.com/liuzhaomax/ovo-user/internal/middleware/auth"
	"github.com/liuzhaomax/ovo-user/internal/middleware/reverse_proxy"
	"github.com/liuzhaomax/ovo-user/internal/middleware/tracing"
	"github.com/liuzhaomax/ovo-user/internal/middleware/validator"
	"github.com/liuzhaomax/ovo-user/internal/middleware_rpc"
	auth2 "github.com/liuzhaomax/ovo-user/internal/middleware_rpc/auth"
	tracing2 "github.com/liuzhaomax/ovo-user/internal/middleware_rpc/tracing"
	validator2 "github.com/liuzhaomax/ovo-user/internal/middleware_rpc/validator"
	"github.com/liuzhaomax/ovo-user/src/api_user/business"
	"github.com/liuzhaomax/ovo-user/src/api_user/handler"
	"github.com/liuzhaomax/ovo-user/src/api_user/model"
	business2 "github.com/liuzhaomax/ovo-user/src/api_user_rpc/business"
	model2 "github.com/liuzhaomax/ovo-user/src/api_user_rpc/model"
)

// Injectors from wire.go:

func InitInjector() (*Injector, func(), error) {
	engine := core.InitGinEngine()
	logger := core.InitLogrus()
	coreLogger := &core.Logger{
		Logger: logger,
	}
	client, cleanup, err := core.InitRedis()
	if err != nil {
		return nil, nil, err
	}
	authAuth := &auth.Auth{
		Logger: coreLogger,
		Redis:  client,
	}
	validatorValidator := &validator.Validator{
		Logger: coreLogger,
		Redis:  client,
	}
	configuration := core.InitTracer()
	tracingTracing := &tracing.Tracing{
		Logger:       coreLogger,
		TracerConfig: configuration,
	}
	reverseProxy := &reverse_proxy.ReverseProxy{
		Logger:      coreLogger,
		RedisClient: client,
	}
	middlewareMiddleware := &middleware.Middleware{
		Auth:         authAuth,
		Validator:    validatorValidator,
		Tracing:      tracingTracing,
		ReverseProxy: reverseProxy,
	}
	db, cleanup2, err := core.InitDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	modelUser := &model.ModelUser{
		DB: db,
	}
	trans := &core.Trans{
		DB: db,
	}
	businessUser := &business.BusinessUser{
		Model: modelUser,
		Tx:    trans,
		Redis: client,
	}
	response := &core.Response{
		Logger: coreLogger,
	}
	rocketMQ := &core.RocketMQ{}
	handlerUser := &handler.HandlerUser{
		Business: businessUser,
		Logger:   coreLogger,
		Res:      response,
		RocketMQ: rocketMQ,
	}
	registry := core.InitPrometheusRegistry()
	apiHandler := &api.Handler{
		Middleware:         middlewareMiddleware,
		HandlerUser:        handlerUser,
		PrometheusRegistry: registry,
	}
	injectorHTTP := InjectorHTTP{
		Engine:  engine,
		Handler: apiHandler,
		DB:      db,
		Redis:   client,
	}
	authRPC := &auth2.AuthRPC{
		Logger: coreLogger,
		Redis:  client,
	}
	validatorRPC := &validator2.ValidatorRPC{
		Logger: coreLogger,
		Redis:  client,
	}
	tracingRPC := &tracing2.TracingRPC{
		Logger:       coreLogger,
		TracerConfig: configuration,
	}
	middlewareRPC := &middleware_rpc.MiddlewareRPC{
		AuthRPC:      authRPC,
		ValidatorRPC: validatorRPC,
		TracingRPC:   tracingRPC,
	}
	modelModelUser := &model2.ModelUser{
		DB: db,
	}
	businessBusinessUser := &business2.BusinessUser{
		Model:    modelModelUser,
		Tx:       trans,
		Redis:    client,
		IRes:     response,
		RocketMQ: rocketMQ,
	}
	handlerRPC := &api.HandlerRPC{
		PrometheusRegistry: registry,
		MiddlewareRPC:      middlewareRPC,
		BusinessRPC:        businessBusinessUser,
	}
	injectorRPC := InjectorRPC{
		HandlerRPC: handlerRPC,
		DB:         db,
		Redis:      client,
	}
	injector := &Injector{
		InjectorHTTP: injectorHTTP,
		InjectorRPC:  injectorRPC,
	}
	return injector, func() {
		cleanup2()
		cleanup()
	}, nil
}
