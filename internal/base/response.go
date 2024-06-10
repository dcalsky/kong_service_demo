package base

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/common/env"
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

func RespondJson(kongArgs KongArgs, c *app.RequestContext, data any) {
	c.JSON(200, CommonResponse{
		Meta: &ResponseMetaData{RequestId: kongArgs.RequestId},
		Data: data,
	})
	c.Abort()
}

func RespondError(kongArgs KongArgs, c *app.RequestContext, except Exception) {
	detail := ""
	if env.InLocal() {
		detail = except.RawError
	}
	c.JSON(except.StatusCode, ErrorResponse{
		Meta: &ResponseMetaData{
			RequestId: kongArgs.RequestId,
			Error: &ErrorObj{
				Message: except.Message,
				Detail:  detail,
			},
		},
	})
	c.Abort()
}
