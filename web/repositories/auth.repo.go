package repositories

import (
	"fmt"
	"log"
	"strings"
	"web-sales/web/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo interface {
	Login(*models.Login) error
}

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Login(login *models.Login) error {
	query := `select email,password from users where email=$1`
	dataDb := new(models.Login)
	err := r.db.Get(dataDb, query, login.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return fmt.Errorf("email tidak ditemukan")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataDb.Password), []byte(login.Password))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
