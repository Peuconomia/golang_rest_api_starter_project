package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang_rest_api_starter_project/core/configs"
	"golang_rest_api_starter_project/core/routers"
)

func main() {

	// region - DI
	options := fx.Options(
		configs.Module,
		routers.Module,
	)
	// endregion

	fx.New(
		options,
		fx.Invoke(func(appConfig *configs.AppConfig, router *routers.AppRouter) {
			// Setup Gin Engine
			ginEngine := gin.Default()

			// set ginMode using appConfig
			gin.SetMode(appConfig.GinMode)

			// Register app routes
			router.RegisterRoutes(ginEngine)

			// RUN API Server on predefined port on
			ginEngine.Run()

		}),
	)
}
