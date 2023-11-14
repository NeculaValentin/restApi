package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)
import "github.com/gin-gonic/gin"

type Person struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Age      uint      `json:"age,omitempty"`
	Children []*Person `json:"children,omitempty"`
}

func getVersion(c *gin.Context) {
	restval := map[string]string{
		"version": "1.0",
	}
	c.JSON(http.StatusOK, restval)
}

func getPerson(c *gin.Context) {
	username := c.Param("username")
	if username == "Vale" {
		person := Person{Age: 2, Name: "Vale", Surname: "Necu"}
		c.IndentedJSON(http.StatusOK, person)
	}

	c.String(http.StatusNotFound, "")

}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Print("")
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/version", getVersion)
	}
	v1.GET("/person", getPerson)

	v1.GET("/info/:username", getPerson)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}

}
