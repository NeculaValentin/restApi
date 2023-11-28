package service

import (
	"restApi/internal/app/common"
	"restApi/internal/app/dao"
	"restApi/internal/app/repository"
)

type AuthServiceImpl struct {
	ur repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		ur: repo,
	}
}

type AuthService interface {
	CreateUser(username, password string) string
	AuthenticateUser(username, password string) (string, error)
	GetVersion() map[string]string
}

func (svc *AuthServiceImpl) GetVersion() map[string]string {
	return map[string]string{
		"version": "1.0",
	}
}

func (svc *AuthServiceImpl) CreateUser(username, password string) string {
	hashedPassword := common.HashPassword(password)
	user := &dao.User{Username: username, Password: hashedPassword} // create a pointer
	svc.ur.Save(user)
	return common.GenerateToken(username)
}

func (svc *AuthServiceImpl) AuthenticateUser(username, password string) (string, error) {
	user, userError := svc.ur.GetUser(username)
	if userError != nil {
		return "", userError
	}
	err := common.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", err
	}
	return common.GenerateToken(username), nil
}
