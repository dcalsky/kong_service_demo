package middleware

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/common/logs"
)

func ExceptionGuard() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if panicData := recover(); panicData != nil {
				logs.Errorf(ctx, "[ExceptionGuard] stack: \n%s", getStack())
				logs.Errorf(ctx, "[ExceptionGuard] panic data: %v", panicData)
				kongArgs := base.GetKongArgs(ctx, c)
				if e, ok := panicData.(base.Exception); ok {
					base.RespondError(kongArgs, c, e)
					return
				}
				base.RespondJson(kongArgs, c, base.InternalError.WithRawError(fmt.Errorf("%v", panicData)))
				return
			}
		}()
		c.Next(ctx)
	}
}

func getStack() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 2048))
	pc := make([]uintptr, 12)
	n := runtime.Callers(3, pc)
	if n == 0 {
		return make([]byte, 0)
	}
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()
		buf.WriteString(frame.Function)
		buf.WriteByte('\n')
		buf.WriteByte('\t')
		buf.WriteString(frame.File)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(frame.Line))
		buf.WriteByte('\n')
		if !more {
			break
		}
	}
	return buf.Bytes()
}
