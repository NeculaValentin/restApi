package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/dao"
	"restApi/internal/app/repository"
	"strconv"
)

type UserServiceImpl struct {
	userRepository *repository.UserRepositoryImpl
}

func NewUserService(repo *repository.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: repo,
	}
}

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)

	}

	c.JSON(http.StatusOK, common.BuildResponse(common.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {

	log.Info("start to execute program add data user")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)

	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)

	}

	c.JSON(http.StatusOK, common.BuildResponse(common.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {

	log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)

	}

	c.JSON(http.StatusOK, common.BuildResponse(common.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)

	}

	c.JSON(http.StatusOK, common.BuildResponse(common.Success, common.Null()))
}
