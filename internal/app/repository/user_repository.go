package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/dao"
)

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	err := db.AutoMigrate(&dao.User{})
	if err != nil {
		log.Error("Error init", err)
		return nil
	}
	return &UserRepositoryImpl{
		db: db,
	}
}

type UserRepository interface {
	FindAllUser() ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(user *dao.User) dao.User
	DeleteUserById(id int)
	GetUser(username string) dao.User
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (ur UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var users []dao.User
	result := ur.db.Find(&users) // returns
	if result.Error != nil {
		log.Error("Got an error when get all user. Error: ", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (ur UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	return user, nil
}

func (ur UserRepositoryImpl) GetUser(username string) dao.User {
	return dao.User{Username: username}
}

func (ur UserRepositoryImpl) Save(user *dao.User) dao.User {
	result := ur.db.Save(user)
	if result.Error != nil {
		_ = common.NewAPIError(http.StatusInternalServerError, result.Error, "error when saving user")
	}
	return *user
}

func (ur UserRepositoryImpl) DeleteUserById(id int) {
	result := ur.db.Delete(&dao.User{}, id)
	if result.Error != nil {
		_ = common.NewAPIError(http.StatusInternalServerError, result.Error, "error when deleting user")
	}
}
