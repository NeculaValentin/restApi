package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/controller"
)

type Person struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Age      uint      `json:"age,omitempty"`
	Children []*Person `json:"children,omitempty"`
}

func getPerson(c *gin.Context) {
	username := c.Param("username")
	if username == "Vale" {
		person := Person{Age: 2, Name: "Vale", Surname: "Necu"}
		c.IndentedJSON(http.StatusOK, person)
	}

	c.String(http.StatusNotFound, "")
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(common.GlobalErrorHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	//userCtrl := controller.NewUserController()

	fileCtrl := controller.NewFileController()
	fileCtrl.RegisterRoutes(v1)

	authCtrl := controller.NewAuthController()
	authCtrl.RegisterRoutes(v1)

	//user := v1.Group("/user")
	//user.GET("", userCtrl.GetAllUserData)
	//user.POST("", userCtrl.AddUserData)
	//user.GET("/:userID", userCtrl.GetUserById)
	//user.DELETE("/:userID", userCtrl.DeleteUser)

	return r
}
