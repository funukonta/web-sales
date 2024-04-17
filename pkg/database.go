package pkg

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgre() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("error ping", err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Panicln("getwd", err)
	}

	var folderMigrate string
	if strings.Contains(dir, "cmd") {
		folderMigrate = "../migrate"
	} else {
		folderMigrate = "./migrate"
	}

	m, err := migrate.New(fmt.Sprintf("file://%s", folderMigrate), connStr)
	if err != nil {
		log.Panicln("error migrate", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panicln("error migrate up", err.Error())
	}

	return db, nil
}
