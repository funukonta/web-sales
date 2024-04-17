package services

import (
	"fmt"
	"web-sales/web/models"
	"web-sales/web/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(data *models.RegisterUser) (*models.RegisterUserRes, error)
	GetAll() ([]models.Users, error)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(data *models.RegisterUser) (*models.RegisterUserRes, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	data.Password = string(passwordHash)

	newCust, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	cust := &models.RegisterUserRes{
		Email:     newCust.Email,
		Password:  newCust.Password,
		Name:      newCust.Name,
		CreatedAt: newCust.CreatedAt,
	}

	return cust, err
}

func (s *userService) GetAll() ([]models.Users, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("tidak ada data")
	}

	return users, nil
}
