package middleware

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/stretchr/testify/require"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/common/logs"
	"github.com/dcalsky/kong_service_demo/internal/tests/hertz_bundle"
)

func TestExceptionGuard(t *testing.T) {
	opt := config.NewOptions([]config.Option{})
	engine := route.NewEngine(opt)
	engine.Use(ExceptionGuard())
	engine.POST("/", func(c context.Context, ctx *app.RequestContext) {
		base.PanicIf(true, base.InvalidParamErr.WithRawError(fmt.Errorf("mock error")))
	})
	recorder := hertz_bundle.PerformRequest(engine, "POST", "/", nil)
	require.NotNil(t, recorder)
	var response base.CommonResponse
	getRecordResult(t, recorder, &response)
}

func getRecordResult(t *testing.T, recorder *hertz_bundle.ResponseRecorder, result any) {
	data := recorder.Result().Body()
	response := base.CommonResponse{
		Data: result,
	}
	logs.Infof(context.Background(), "[Test] response body: %v", string(data))
	err := json.Unmarshal(data, &response)
	require.NoError(t, err)
}
