package service

type AuthServiceImpl struct {
}

func NewAuthService() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

type AuthService interface {
	GetVersion() map[string]string
}

func (a AuthServiceImpl) GetVersion() map[string]string {
	return map[string]string{
		"version": "1.0",
	}
}
