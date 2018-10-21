package daos

import (
	"boilerplate/app"
	"boilerplate/models"
	"fmt"
)

type AuthDAO struct{}

func NewAuthDAO() *AuthDAO {
	return &AuthDAO{}
}

func (dao *AuthDAO) Login(rs app.RequestScope, user *models.User) *models.User {
	fmt.Println(user)
	/*user := &models.User{1, "nir", "0545630117", "123"}

	  rs.Db().Create(user)*/

	return user
}
