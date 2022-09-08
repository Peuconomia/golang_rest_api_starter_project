package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AppRouter struct {
}

func NewAppRouter() *AppRouter {
	return &AppRouter{}
}

var Module = fx.Provide(
	NewAppRouter,
)

func (router *AppRouter) RegisterRoutes(engine *gin.Engine) {
	// TODO: Register V1, V2 ... Vn routes
}
