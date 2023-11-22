package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	userCtrl := controller.NewUserController()
	authCtrl := controller.NewAuthController()

	r := gin.Default()
	v1 := r.Group("/api/v1")

	//Auth
	//v1.POST("/signup", authCtrl.Signup)
	//v1.POST("/login", authCtrl.Login)
	v1.GET("/version", authCtrl.GetVersion)

	//Documents
	//v1.GET("/<string:username>/<string:doc_id>", getDocument)
	//v1.POST("/<string:username>/<string:doc_id>", postDocument)
	//v1.PUT("/<string:username>/<string:doc_id>", putDocument)
	//v1.DELETE("/<string:username>/<string:doc_id>", deleteDocument)

	//Collections
	//v1.GET("/<string:username>/_all_docs", getAllDocuments)

	user := v1.Group("/user")
	user.GET("", userCtrl.GetAllUserData)
	user.POST("", userCtrl.AddUserData)
	user.GET("/:userID", userCtrl.GetUserById)
	user.DELETE("/:userID", userCtrl.DeleteUser)

	return r
}
