package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/internal/app/service"
)

func NewAuthController() *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: service.NewAuthService(),
	}
}

type AuthController interface {
	GetVersion(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (a AuthControllerImpl) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, a.svc.GetVersion())
}
