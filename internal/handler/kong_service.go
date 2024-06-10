package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/service"
)

func CreateKongService(ctx context.Context, c *app.RequestContext) {
	var req dto.CreateKongServiceRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.KongServiceSvc.CreateKongService(ctx, args, req)
	base.RespondJson(args, c, data)
}

func DescribeKongService(ctx context.Context, c *app.RequestContext) {
	var req dto.DescribeKongServiceRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.KongServiceSvc.DescribeKongService(ctx, args, req)
	base.RespondJson(args, c, data)
}

func DeleteKongService(ctx context.Context, c *app.RequestContext) {
	var req dto.DeleteKongServiceRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	service.KongServiceSvc.DeleteKongService(ctx, args, req)
	base.RespondJson(args, c, dto.EmptyResponse{})
}

func UpdateKongService(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateKongServiceRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.KongServiceSvc.UpdateKongService(ctx, args, req)
	base.RespondJson(args, c, data)
}

func ListKongServices(ctx context.Context, c *app.RequestContext) {
	var req dto.ListKongServicesRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.KongServiceSvc.ListKongServices(ctx, args, req)
	base.RespondJson(args, c, data)
}
