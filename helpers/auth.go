package helpers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-identity-service/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expiration, err := strconv.Atoi(config.Env.TokenExpirationInHours)
	if err != nil {
		fmt.Println("Error during env var conversion")
		return
	}
	expirationTime := time.Now().Add(time.Duration(expiration) * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(config.Env.JwtSecret))
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Env.JwtSecret), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
