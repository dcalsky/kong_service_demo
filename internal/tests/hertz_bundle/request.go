package hertz_bundle

import (
	"context"
	"io"

	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/route"
)

type Header struct {
	Key   string
	Value string
}

type Body struct {
	Body io.Reader
	Len  int
}

// PerformRequest sends a constructed request to given engine without network transporting
func PerformRequest(engine *route.Engine, method, url string, body *Body, headers ...Header) *ResponseRecorder {
	ctx := engine.NewContext()

	var r *protocol.Request
	if body != nil && body.Body != nil {
		r = protocol.NewRequest(method, url, body.Body)
		r.CopyTo(&ctx.Request)
		if engine.GetOptions().StreamRequestBody {
			ctx.Request.SetBodyStream(body.Body, body.Len)
		} else {
			buf, err := io.ReadAll(&io.LimitedReader{R: body.Body, N: int64(body.Len)})
			ctx.Request.SetBody(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
		}
	} else {
		r = protocol.NewRequest(method, url, nil)
		r.CopyTo(&ctx.Request)
	}

	for _, v := range headers {
		if ctx.Request.Header.Get(v.Key) != "" {
			ctx.Request.Header.Add(v.Key, v.Value)
		} else {
			ctx.Request.Header.Set(v.Key, v.Value)
		}
	}

	engine.ServeHTTP(context.Background(), ctx)

	w := NewRecorder()
	h := w.Header()
	ctx.Response.Header.CopyTo(h)

	w.WriteHeader(ctx.Response.StatusCode())
	w.Write(ctx.Response.Body())
	w.Flush()
	return w
}
