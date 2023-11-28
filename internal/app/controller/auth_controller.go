package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/dao"
	"restApi/internal/app/repository"
	"restApi/internal/app/service"
)

type AuthControllerImpl struct {
	svc service.AuthService
}

func NewAuthController() *AuthControllerImpl {
	userRepo := repository.NewUserRepository(common.ConnectToDB())
	return &AuthControllerImpl{svc: service.NewAuthService(userRepo)}
}

type AuthController interface {
	GetVersion(c *gin.Context)
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

// RegisterRoutes registers the authentication routes
func (ac *AuthControllerImpl) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/version", ac.GetVersion)
	router.POST("/signup", ac.Signup)
	router.POST("/login", ac.Login)
}

// GetVersion handles the /version endpoint
func (ac *AuthControllerImpl) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, ac.svc.GetVersion())
}

// Signup handles the /signup endpoint
func (ac *AuthControllerImpl) Signup(c *gin.Context) {
	user := checkUserInput(c)
	token := ac.svc.CreateUser(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// Login handles the /login endpoint
func (ac *AuthControllerImpl) Login(c *gin.Context) {
	user := checkUserInput(c)
	token := ac.svc.AuthenticateUser(user.Username, user.Password)
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
