package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web-sales/pkg"
	"web-sales/web/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mux := http.NewServeMux()

	db, err := pkg.ConnectPostgre()
	if err != nil {
		log.Panic(err.Error())
		return
	}

	routes.Routes(mux, db)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Println("API Server listening to port", port)
	http.ListenAndServe(port, mux)
}

/*
- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- migrate create -ext sql -dir migrate create_table_users
- migrate create -ext sql -dir migrate create_table_sales
*/
