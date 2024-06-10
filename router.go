package main

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/dcalsky/kong_service_demo/internal/config"
	"github.com/dcalsky/kong_service_demo/internal/handler"
	"github.com/dcalsky/kong_service_demo/internal/middleware"
)

func RegisterRoutes() {
	r := server.Default(
		server.WithHostPorts("0.0.0.0:8002"),
	)
	r.Use(middleware.TrafficLogger())
	r.Use(middleware.ExceptionGuard())
	registerHttp(r)

	r.Spin()
}

func registerHttp(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusNotFound, nil)
	})

	v1 := r.Group("/api/v1")
	jwtValidator := middleware.JwtValidator(config.Conf.KongSecret)
	constrainedV1 := v1.Group("", jwtValidator)
	{
		constrainedV1.POST("/CreateKongService", handler.CreateKongService)
		constrainedV1.POST("/DescribeKongService", handler.DescribeKongService)
		constrainedV1.POST("/DeleteKongService", handler.DeleteKongService)
		constrainedV1.POST("/ListKongServices", handler.ListKongServices)
		constrainedV1.POST("/UpdateKongService", handler.UpdateKongService)
	}
	{
		v1.POST("/Register", handler.Register)
		v1.POST("/Login", handler.Login)
	}
}
