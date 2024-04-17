package services

import (
	"time"
	"web-sales/web/models"
	"web-sales/web/repositories"
)

type SalesService interface {
	Input(*models.InputSales) (*models.Sales, error)
	Report(*models.ReportSalesReq) ([]models.ReportSalesRes, error)
}

type salesService struct {
	repo repositories.SalesRepo
}

func NewSalesService(repo repositories.SalesRepo) SalesService {
	return &salesService{repo: repo}
}
func (s *salesService) Input(data *models.InputSales) (*models.Sales, error) {
	var err error
	data.TransactionDate, err = time.Parse("2006-01-02", data.TransactionDateJson)
	if err != nil {
		return nil, err
	}

	sales, err := s.repo.Input(data)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *salesService) Report(data *models.ReportSalesReq) ([]models.ReportSalesRes, error) {
	sales, err := s.repo.Report(data)

	return sales, err
}
