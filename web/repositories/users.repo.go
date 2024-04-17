package repositories

import (
	"context"
	"web-sales/web/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(data *models.RegisterUser) (*models.Users, error)
	GetAll() ([]models.Users, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(data *models.RegisterUser) (*models.Users, error) {
	query := `INSERT INTO users (email, name, password) 
	VALUES ($1, $2, $3)
	RETURNING name, email, password, created_at
	`
	newUser := new(models.Users)
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Get(newUser, query, data.Email, data.Name, data.Password)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	return newUser, err
}

func (r *userRepo) GetAll() ([]models.Users, error) {
	query := `select email,name,password from users`
	users := []models.Users{}
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}
