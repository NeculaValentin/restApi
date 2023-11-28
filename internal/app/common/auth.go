package common

import (
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY") // Get the secret key from environment variable
var expiryStr, _ = strconv.Atoi(os.Getenv("JWT_EXPIRATION_MINUTES"))

func GenerateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * time.Duration(expiryStr)).Unix(), // Token expires after 5 minutes
	})

	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		log.Error("Got an error when generating the token", err)
		return ""
	}

	return tokenString
}

// HashPassword with bcrypt hashing method, the salt is automatically generated as part of the hashing process
// and is included within the resulting hashed password string.
// The format is generally something like $2a$[cost]$[22 character salt][31 character hash].
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Got an error when generating the hash", err)
		return ""
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
