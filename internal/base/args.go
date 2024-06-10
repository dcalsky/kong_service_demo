package base

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/dcalsky/kong_service_demo/internal/common/logid"
)

const (
	KongAccountIdKey = "KONG_ACCOUNT_ID"
)

type KongAccountId string

func (s KongAccountId) String() string {
	return string(s)
}

type KongArgs struct {
	TraceInfo
	AccountId KongAccountId
}

type TraceInfo struct {
	RequestId string
}

func GetKongArgs(ctx context.Context, c *app.RequestContext) KongArgs {
	out := KongArgs{
		TraceInfo: TraceInfo{
			RequestId: logid.LogId(ctx),
		},
		AccountId: KongAccountId(c.GetString(KongAccountIdKey)),
	}
	return out
}
