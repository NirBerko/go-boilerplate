package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func Transactional(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := GetRequestScope(c)
		rs.SetDB(db)
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := GetRequestScope(c)

		authentication := c.GetHeader("Authorization")

		claims := jwt.MapClaims{}
		token, _ := jwt.ParseWithClaims(strings.Split(authentication, "Bearer ")[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(Config.Jwt.JWTSigningKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rs.SetUserID(int(claims["id"].(float64)))
		} else {
			// TODO: make errors manager (APIError)
		}
	}
}
