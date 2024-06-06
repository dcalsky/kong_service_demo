package logs

import (
	"context"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/dcalsky/kong_service_demo/internal/common/logid"
)

func TestZapLog(t *testing.T) {
	ctx := context.Background()
	ctx = logid.SetLogId(ctx, logid.NewLogId())
	MustInit(zapcore.DebugLevel)
	Infof(ctx, "name: %s, age: %d", "kong", 18)
}
