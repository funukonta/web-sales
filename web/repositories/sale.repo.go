package repositories

import (
	"context"
	"web-sales/web/models"

	"github.com/jmoiron/sqlx"
)

type SalesRepo interface {
	Input(*models.InputSales) (*models.Sales, error)
	Report(*models.ReportSalesReq) ([]models.ReportSalesRes, error)
}

type salesRepo struct {
	db *sqlx.DB
}

func NewSalesRepo(db *sqlx.DB) SalesRepo {
	return &salesRepo{db: db}
}

func (r *salesRepo) Input(data *models.InputSales) (*models.Sales, error) {
	query := `INSERT into sales (email,sales_type,transactionDate,nominal) values ($1,$2,$3,$4)
	RETURNING email, id, sales_type, nominal, transactionDate;
	`

	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sales := &models.Sales{}
	err = tx.Get(sales, query, data.Email, data.SalesType, data.TransactionDate, data.Nominal)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *salesRepo) Report(data *models.ReportSalesReq) ([]models.ReportSalesRes, error) {
	query := `
	SELECT 
    u.name AS User,
    COUNT(DISTINCT s.transactionDate) AS Jumlah_Hari_Kerja,
    SUM(CASE WHEN s.sales_type = 'Barang' THEN 1 ELSE 0 END) AS Jumlah_Transaksi_Barang,
    SUM(CASE WHEN s.sales_type = 'Jasa' THEN 1 ELSE 0 END) AS Jumlah_Transaksi_Jasa,
    SUM(CASE WHEN s.sales_type = 'Barang' THEN s.nominal ELSE 0 END) AS Nominal_Transaksi_Barang,
    SUM(CASE WHEN s.sales_type = 'Jasa' THEN s.nominal ELSE 0 END) AS Nominal_Transaksi_Jasa
	FROM 
		Users u
	LEFT JOIN 
		Sales s ON u.email = s.email
	WHERE 
		s.transactionDate BETWEEN $1 AND $2
	GROUP BY 
		u.email, u.name`

	report := []models.ReportSalesRes{}
	err := r.db.Select(&report, query, data.StartDate, data.EndDate)
	if err != nil {
		return nil, err
	}

	return report, nil
}
