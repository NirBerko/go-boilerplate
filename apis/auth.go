package apis

import (
	"boilerplate/app"
	"boilerplate/errors"
	"boilerplate/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	authService interface {
		Login(rs app.RequestScope, user *models.User)
		Register(rs app.RequestScope, user *models.User) error
	}

	authResource struct {
		service authService
	}
)

func ServeAuthResource(rg *gin.Engine, service authService) {
	r := &authResource{service}

	rg.POST("/login", r.Auth)
	rg.POST("/register", r.Register)
}

func (r *authResource) Auth(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	//r.service.Login(app.GetRequestScope(c), &user)*/

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

func (r *authResource) Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	err := r.service.Register(app.GetRequestScope(c), &user)

	if err != nil {
		errorHandler := errors.InternalServerError(err)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(200, user)
	}
}
