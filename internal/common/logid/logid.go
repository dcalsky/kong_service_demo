package logid

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	maxRandNum = 1 << 20
	LogIdKey   = "KONG_LOGID"
)

func getMsTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NewLogId generates a toy log id for DEMO
func NewLogId() string {
	r := rand.Int63n(maxRandNum)
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatInt(getMsTimestamp(), 10))
	sb.WriteString(strconv.FormatInt(r, 16))
	return sb.String()
}

func SetLogId(ctx context.Context, logid string) context.Context {
	return context.WithValue(ctx, LogIdKey, logid)
}

func LogId(ctx context.Context) string {
	id, _ := getStringFromCtx(ctx, LogIdKey)
	return id
}

func getStringFromCtx(ctx context.Context, key string) (string, bool) {
	if ctx == nil {
		return "", false
	}

	v := ctx.Value(key)
	if v == nil {
		return "", false
	}

	switch v := v.(type) {
	case string:
		return v, true
	case *string:
		if v == nil {
			return "", false
		}
		return *v, true
	}
	return "", false
}
