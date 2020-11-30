package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tmohammad78/food-api/services"
	"log"
	"net/http"
)

func AuthorizedJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := context.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		fmt.Println(tokenString)
		token, err := services.NewJWTService().ValidateToken(string(tokenString))
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println(err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
