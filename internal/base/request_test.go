package base

import (
	"net/http"
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/require"
)

func TestDumpHttpRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "http://example.com", nil)
	require.NoError(t, err)
	reqStr := DumpHttpRequest(req)
	require.NotEmpty(t, reqStr)
	require.True(t, strings.Contains(reqStr, "curl"))

	t.Run("request is nil", func(t *testing.T) {
		reqStr = DumpHttpRequest(nil)
		require.Empty(t, reqStr)
	})
}

func TestDumpHertzRequest(t *testing.T) {
	req := &protocol.Request{}
	req.SetRequestURI("http://example.com")
	req.SetMethod("POST")
	req.SetBody([]byte("test"))
	req.SetHeader("Content-Type", "application/json")
	reqStr := DumpHertzRequest(req)
	require.NotEmpty(t, reqStr)
	require.True(t, strings.Contains(reqStr, "curl"))
}
