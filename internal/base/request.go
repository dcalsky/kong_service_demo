package base

import (
	"bytes"
	"net/http"

	"github.com/cloudwego/hertz/pkg/protocol"
	"moul.io/http2curl"
)

func getCompatRequest(req *protocol.Request) (*http.Request, error) {
	r, err := http.NewRequest(string(req.Method()), req.URI().String(), bytes.NewReader(req.Body()))
	if err != nil {
		return r, err
	}

	h := make(map[string][]string, req.Header.Len())
	req.Header.VisitAll(func(k, v []byte) {
		h[string(k)] = append(h[string(k)], string(v))
	})

	r.Header = h
	return r, nil
}

func DumpHttpRequest(req *http.Request) string {
	if req == nil {
		return ""
	}
	if cmd, err := http2curl.GetCurlCommand(req); err == nil {
		return cmd.String()
	}
	return ""
}

func DumpHertzRequest(req *protocol.Request) string {
	proxyRequest, err := getCompatRequest(req)
	if err != nil {
		return ""
	}
	return DumpHttpRequest(proxyRequest)
}
