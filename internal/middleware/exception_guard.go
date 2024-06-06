package middleware

import (
	"context"
	"fmt"
	"runtime"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/common/logs"
)

func ExceptionGuard() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if panicData := recover(); panicData != nil {
				logs.Errorf(ctx, "[ExceptionGuard] stack: %s", getStack(1<<12))
				logs.Errorf(ctx, "[ExceptionGuard] panic data: %v", panicData)
				if e, ok := panicData.(base.Exception); ok {
					base.RespondError(ctx, c, e)
					return
				}
				base.RespondJson(ctx, c, base.InternalError.WithRawError(fmt.Errorf("%v", panicData)))
				return
			}
		}()
		c.Next(ctx)
	}
}

func getStack(size int) []byte {
	buf := make([]byte, size)
	runtime.Stack(buf, false)
	return buf
}
