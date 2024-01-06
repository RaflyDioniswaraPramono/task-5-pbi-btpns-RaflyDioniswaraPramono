package helpers

import (	
	"os"

	"example.com/go-rest-api-gin/models"	
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user models.Users) string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)

	claims["user"] = user
	token, _ := sign.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return token
}