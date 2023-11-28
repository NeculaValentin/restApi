package config

import (
	"github.com/gin-gonic/gin"
	"restApi/internal/app/common"
	"restApi/internal/app/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(common.GlobalErrorHandler())

	v1 := r.Group("/api/v1")

	fileCtrl := controller.NewFileController()
	fileCtrl.RegisterRoutes(v1)

	authCtrl := controller.NewAuthController()
	authCtrl.RegisterRoutes(v1)
	return r
}
