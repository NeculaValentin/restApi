package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	Save(user *dao.User) (dao.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var users []dao.User
	result := u.db.Find(&users) // returns
	if result.Error != nil {
		log.Error("Got an error when get all user. Error: ", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}
