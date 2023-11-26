package service

type AuthServiceImpl struct {
}

func NewAuthService() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

type AuthService interface {
	GetVersion() map[string]string
	AuthenticateUser(username string, password string) string
	CreateUser(username string, password string) string
}

func (a AuthServiceImpl) GetVersion() map[string]string {
	return map[string]string{
		"version": "1.0",
	}
}

func (a AuthServiceImpl) AuthenticateUser(username string, password string) string {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) CreateUser(username string, password string) string {
	//TODO implement me
	panic("implement me")
}
