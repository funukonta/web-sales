package routes

import (
	"net/http"
	"web-sales/web/handlers"
	"web-sales/web/repositories"
	"web-sales/web/services"

	"github.com/jmoiron/sqlx"
)

func UserRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewUserRepo(db)
	service := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(service)

	mux.Handle("POST /register", Handler(userHandler.Create))
	mux.Handle("GET /customers", authMiddleware(Handler(userHandler.GetAll)))
}
