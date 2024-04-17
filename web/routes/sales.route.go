package routes

import (
	"net/http"
	"web-sales/web/handlers"
	"web-sales/web/repositories"
	"web-sales/web/services"

	"github.com/jmoiron/sqlx"
)

func SalesRoute(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewSalesRepo(db)
	service := services.NewSalesService(repo)
	salesHandler := handlers.NewSalesHandler(service)

	mux.Handle("POST /sales/input", authMiddleware(Handler(salesHandler.Input)))
	mux.Handle("POST /sales/report", authMiddleware(Handler(salesHandler.Report)))
}
