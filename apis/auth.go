package apis

import (
	"boilerplate/app"
	"boilerplate/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	authService interface {
		Login(rs app.RequestScope, user *models.User)
	}

	authResource struct {
		service authService
	}
)

func ServeAuthResource(rg *gin.Engine, service authService) {
	r := &authResource{service}

	rg.POST("/login", r.Auth)
}

func (r *authResource) Auth(c *gin.Context) {
	/*var user models.User
	  c.BindJSON(&user)

	  r.service.Login(app.GetRequestScope(c), &user)*/

	user := models.User{1, "nir", "0545630117", "123"}

	token, err := app.NewJWT(jwt.MapClaims{
		"id":  user.GetID(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}, app.Config.Jwt.JWTSigningKey)

	if err != nil {
		fmt.Println("UNAUTHORIZED")
	}

	c.JSON(200, map[string]string{
		"token": token,
	})
}
