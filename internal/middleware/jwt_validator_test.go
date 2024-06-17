package middleware

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/tests/hertz_bundle"
)

func TestJwtValidator(t *testing.T) {
	opt := config.NewOptions([]config.Option{})
	engine := route.NewEngine(opt)
	const secret = "kong_test"
	const accountId = 1024
	const email = "foo@example.com"
	engine.Use(JwtValidator(secret))

	claims := &dto.JwtClaims{
		ExpiredAt: time.Now().Add(time.Minute * 5),
		AccountId: uint(accountId),
		Email:     email,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	require.NoError(t, err)
	engine.POST("/", func(ctx context.Context, c *app.RequestContext) {
		kongArgs := base.GetKongArgs(ctx, c)
		require.Equal(t, base.KongAccountId(accountId), kongArgs.AccountId)
		base.RespondJson(kongArgs, c, map[string]string{
			"email": kongArgs.AccountEmail,
		})
	})
	recorder := hertz_bundle.PerformRequest(engine, "POST", "/", nil, hertz_bundle.Header{
		Key:   authorizationHeaderKey,
		Value: "Bearer " + token,
	})
	require.NotNil(t, recorder)
	var response base.CommonResponse
	getRecordResult(t, recorder, &response)
}
