package handlers

import (
	"net/http"
	"web-sales/pkg"
	"web-sales/web/models"
	"web-sales/web/services"
)

type SalesHandler interface {
	Input(w http.ResponseWriter, r *http.Request) error
	Report(w http.ResponseWriter, r *http.Request) error
}

type salesHandler struct {
	serv services.SalesService
}

func NewSalesHandler(serv services.SalesService) SalesHandler {
	return &salesHandler{serv: serv}
}

func (h *salesHandler) Input(w http.ResponseWriter, r *http.Request) error {
	sale := new(models.InputSales)
	if err := pkg.GetJsonBody(r, sale); err != nil {
		return err
	}

	newSale, err := h.serv.Input(sale)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: newSale, Message: "Sales Created!"}).Send(w)
	return nil
}

func (h *salesHandler) Report(w http.ResponseWriter, r *http.Request) error {
	sale := new(models.ReportSalesReq)
	if err := pkg.GetJsonBody(r, sale); err != nil {
		return err
	}

	sales, err := h.serv.Report(sale)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: sales, Message: "Sales Created!"}).Send(w)
	return nil
}
