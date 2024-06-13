package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/service"
)

func CreateOrganization(ctx context.Context, c *app.RequestContext) {
	var req dto.CreateOrganizationRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.OrganizationSvc.CreateOrganization(ctx, args, req)
	base.RespondJson(args, c, data)
}

func AddAccountToOrganization(ctx context.Context, c *app.RequestContext) {
	var req dto.JoinOrganizationRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	service.OrganizationSvc.AddAccountToOrganization(ctx, args, req)
	base.RespondJson(args, c, dto.EmptyResponse{})
}

func RemoveAccountFromOrganization(ctx context.Context, c *app.RequestContext) {
	var req dto.QuitOrganizationRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	service.OrganizationSvc.RemoveAccountFromOrganization(ctx, args, req)
	base.RespondJson(args, c, dto.EmptyResponse{})
}

func DescribeOrganization(ctx context.Context, c *app.RequestContext) {
	var req dto.DescribeOrganizationRequest
	err := c.BindAndValidate(&req)
	base.PanicIfErr(err, base.InvalidParamErr.WithRawError(fmt.Errorf("invalid req: %+v, err: %v", req, err)))

	args := base.GetKongArgs(ctx, c)
	data := service.OrganizationSvc.DescribeOrganization(ctx, args, req)
	base.RespondJson(args, c, data)
}
