package v1CoreControllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang_rest_api_starter_project/v1/feature_name/controllers"
)

type Controller struct {
	homeController *v1HomeControllers.HomeController
}

func NewController(homeController *v1HomeControllers.HomeController) *Controller {
	return &Controller{
		homeController: homeController,
	}
}

var Module = fx.Provide(
	NewController,
)

func (controller *Controller) RegisterControllerRoutes(apiGroup *gin.RouterGroup) {

	controller.homeController.RegisterControllerRoutes(apiGroup)
}
