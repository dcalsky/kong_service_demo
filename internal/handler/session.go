package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/service"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var req dto.LoginRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.SessionSvc.Login(ctx, req)
	base.RespondJson(args, c, data)
}

func Register(ctx context.Context, c *app.RequestContext) {
	var req dto.RegisterRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.SessionSvc.Register(ctx, req)
	base.RespondJson(args, c, data)
}
