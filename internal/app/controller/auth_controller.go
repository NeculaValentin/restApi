package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"restApi/internal/app/dao"
	"restApi/internal/app/service"
)

type AuthControllerImpl struct {
	svc service.AuthService
}

func NewAuthController() *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: service.NewAuthService(),
	}
}

type AuthController interface {
	GetVersion(c *gin.Context)
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

// RegisterRoutes registers the authentication routes
func (a *AuthControllerImpl) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/version", a.GetVersion)
	router.POST("/signup", a.Signup)
	router.POST("/login", a.Login)
}

// GetVersion handles the /version endpoint
func (a *AuthControllerImpl) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, a.svc.GetVersion())
}

// Signup handles the /signup endpoint
func (a *AuthControllerImpl) Signup(c *gin.Context) {
	user := checkUserInput(c)
	token := a.svc.CreateUser(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// Login handles the /login endpoint
func (a *AuthControllerImpl) Login(c *gin.Context) {
	user := checkUserInput(c)
	token := a.svc.AuthenticateUser(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// checkUserInput checks if the user input is valid
func checkUserInput(c *gin.Context) dao.User {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
	}
	return user
}
