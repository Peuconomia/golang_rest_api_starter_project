package v1CoreRouters

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	v1Controllers "golang_rest_api_starter_project/v1/core/controllers"
)

type Router struct {
	controller *v1Controllers.Controller
}

func NewV1Router(controller *v1Controllers.Controller) *Router {
	return &Router{
		controller: controller,
	}
}

var Module = fx.Provide(
	NewV1Router,
)

func (router *Router) RegisterRoutes(engine *gin.Engine) {
	// TODO: Register all v1 Routes here
	group := engine.Group("api/v1")

	router.controller.RegisterControllerRoutes(group)
}
