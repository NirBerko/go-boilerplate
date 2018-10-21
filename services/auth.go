package services

import (
	"boilerplate/app"
	"boilerplate/models"
)

type authDAO interface {
	Login(rs app.RequestScope, user *models.User) *models.User
}

type AuthService struct {
	dao authDAO
}

func NewAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (s *AuthService) Login(rs app.RequestScope, user *models.User) {
	s.dao.Login(rs, user)
}
