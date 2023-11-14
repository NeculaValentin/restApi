package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/version", getVersion)
	v1.GET("/person", getPerson)
	v1.GET("/info/:username", getPerson)

	return r
}
