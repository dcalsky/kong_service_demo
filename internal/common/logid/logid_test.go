package logid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogId(t *testing.T) {
	logid := NewLogId()
	require.NotEmpty(t, logid)

	ctx := context.Background()
	ctx = SetLogId(ctx, logid)
	ctxLogId := ctx.Value(LogIdKey)
	require.NotNil(t, ctxLogId)
	require.Equal(t, logid, ctxLogId.(string))

	logIdStr := LogId(ctx)
	require.Equal(t, logid, logIdStr)
}
