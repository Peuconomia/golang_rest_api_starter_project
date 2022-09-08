package v1HomeControllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

var Module = fx.Provide(
	NewHomeController,
)

func (homeController *HomeController) RegisterControllerRoutes(rootGroup *gin.RouterGroup) {
	group := rootGroup.Group("home")

	group.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

}
