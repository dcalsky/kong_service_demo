package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type pingResponse struct {
	Message string
}

func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, pingResponse{Message: "pong!"})
}
