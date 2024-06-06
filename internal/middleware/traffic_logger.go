package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/common/logid"
	"github.com/dcalsky/kong_service_demo/internal/common/logs"
)

func TrafficLogger() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx = logid.SetLogId(ctx, logid.NewLogId())
		reqStr := base.DumpHertzRequest(&c.Request)
		startAt := time.Now()
		logs.Infof(ctx, "[TrafficLogger] http request at %s, %v", startAt.String(), reqStr)
		c.Next(ctx)
		responseAt := time.Now()
		logs.Infof(ctx, "[TrafficLogger] http response cost: %d ms, body: %s", responseAt.Sub(startAt).Milliseconds(), string(c.GetResponse().Body()))
	}
}
