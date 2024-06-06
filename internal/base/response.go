package base

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/common/env"
	"github.com/dcalsky/kong_service_demo/internal/common/logid"
)

type ErrorObj struct {
	Message string
	Detail  string `json:",omitempty"`
}

type ResponseMetaData struct {
	RequestId string    `json:"RequestId"`
	Error     *ErrorObj `json:",omitempty"`
}

type ErrorResponse struct {
	Meta *ResponseMetaData `json:",omitempty"`
}

type CommonResponse struct {
	Meta *ResponseMetaData `json:",omitempty"` // for future
	Data any
}

func RespondJson(ctx context.Context, c *app.RequestContext, data any) {
	c.JSON(200, CommonResponse{
		Meta: &ResponseMetaData{RequestId: logid.LogId(ctx)},
		Data: data,
	})
	c.Abort()
}

func RespondError(ctx context.Context, c *app.RequestContext, except Exception) {
	detail := ""
	if env.InLocal() {
		detail = except.RawError
	}
	c.JSON(except.StatusCode, ErrorResponse{
		Meta: &ResponseMetaData{
			RequestId: logid.LogId(ctx),
			Error: &ErrorObj{
				Message: except.Message,
				Detail:  detail,
			},
		},
	})
	c.Abort()
}
