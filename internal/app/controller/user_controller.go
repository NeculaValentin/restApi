package controller

import (
	"github.com/gin-gonic/gin"
	"restApi/internal/app/repository"
	"restApi/internal/app/service"
)

type UserControllerImpl struct {
	svc service.UserService
}

func NewUserController() *UserControllerImpl {
	userRepo := repository.NewUserRepository(repository.ConnectToDB())
	return &UserControllerImpl{svc: service.NewUserService(userRepo)}
}

type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}
