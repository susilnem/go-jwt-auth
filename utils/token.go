package utils

import (
	model "go-jwt/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//validatepassword
func ValidatePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}

func GenerateToken(user model.User) (string, error) {
	tokenlife, err := strconv.Atoi(os.Getenv("TOKEN_LIFE"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(tokenlife)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))

}
