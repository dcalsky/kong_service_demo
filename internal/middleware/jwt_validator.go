package middleware

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

const (
	authorizationHeaderKey = "Authorization"
)

var (
	jwtClaimsPool = sync.Pool{
		New: func() any {
			return &dto.JwtClaims{}
		},
	}
)

func JwtValidator(secret string) app.HandlerFunc {
	secretBytes := []byte(secret)
	keyFunc := func(token *jwt.Token) (any, error) {
		return secretBytes, nil
	}
	jwtParser := jwt.NewParser(jwt.WithoutClaimsValidation())
	return func(c context.Context, ctx *app.RequestContext) {
		authorizationToken := string(ctx.GetHeader(authorizationHeaderKey))
		tokenStr, foundPrefix := strings.CutPrefix(authorizationToken, "Bearer ")
		if len(tokenStr) == 0 || !foundPrefix {
			panic(base.AuthenticationDenied.WithRawError(fmt.Errorf("invalid token, token: %s", authorizationToken)))
		}
		claims := jwtClaimsPool.Get().(*dto.JwtClaims)
		token, err := jwtParser.ParseWithClaims(tokenStr, claims, keyFunc)
		defer jwtClaimsPool.Put(claims)
		if err != nil {
			panic(base.AuthenticationDenied.WithRawError(err))
		}
		if !token.Valid {
			panic(base.AuthenticationDenied.WithRawError(fmt.Errorf("invalid token, token: %s", authorizationToken)))
		}
		if claims.ExpiredAt.Before(time.Now()) {
			panic(base.AuthenticationExpired)
		}
		base.SetKongArgsAccountId(ctx, claims.AccountId)
		ctx.Next(c)
	}
}
