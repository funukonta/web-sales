package models

import "time"

type Sales struct {
	ID              int       `json:"id" db:"id"`
	Email           string    `json:"email" db:"email"`
	TransactionDate time.Time `json:"transactionDate" db:"transactiondate"`
	SalesType       string    `json:"sales_type" db:"sales_type"`
	Nominal         float64   `json:"nominal" db:"nominal"`
	CreatedAt       time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updatedAt"`
}

type InputSales struct {
	Email               string `json:"email"`
	SalesType           string `json:"sales_type"`
	TransactionDateJson string `json:"transactionDate"`
	TransactionDate     time.Time
	Nominal             float64 `json:"nominal"`
}

type ReportSalesReq struct {
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type ReportSalesRes struct {
	Jumlah_hari_kerja        int     `json:"jumlah_hari_kerja"`
	Jumlah_transaksi_barang  int     `json:"jumlah_transaksi_barang"`
	Jumlah_transaksi_jasa    int     `json:"jumlah_transaksi_jasa"`
	Nominal_transaksi_barang float64 `json:"nominal_transaksi_barang"`
	Nominal_transaksi_jasa   float64 `json:"nominal_transaksi_jasa"`
	User                     string  `json:"user"`
}
