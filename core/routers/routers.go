package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	v1Routers "golang_rest_api_starter_project/v1/core/routers"
)

type AppRouter struct {
	v1Router *v1Routers.Router
}

func NewAppRouter(v1Router *v1Routers.Router) *AppRouter {
	return &AppRouter{
		v1Router: v1Router,
	}
}

var Module = fx.Provide(
	NewAppRouter,
)

func (router *AppRouter) RegisterRoutes(engine *gin.Engine) {
	// TODO: Register V1, V2 ... Vn routes
	router.v1Router.RegisterRoutes(engine)
}
