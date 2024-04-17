package services

import (
	"fmt"
	"log"
	"os"
	"time"
	"web-sales/web/models"
	"web-sales/web/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(*models.Login, time.Time) (string, error)
}

type authService struct {
	repo repositories.AuthRepo
}

func NewAuthService(repo repositories.AuthRepo) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(login *models.Login, expirationTime time.Time) (string, error) {

	err := s.repo.Login(login)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("email / password salah")
	}

	claims := &models.Claims{
		Username: login.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    login.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims.RegisteredClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
