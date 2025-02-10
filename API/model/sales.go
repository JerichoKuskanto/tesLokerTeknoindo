package model

import "time"

type Sales struct {
	SalesID     int       `json:"salesID"`
	Kode        string    `json:"kode"`
	Tgl         time.Time `json:"tgl"`
	CustID      int       `json:"custID"`
	Subtotal    float32   `json:"subtotal"`
	Diskon      float32   `json:"diskon"`
	Ongkir      float32   `json:"ongkir"`
	Total_bayar float32   `json:"total_bayar"`
}

// ConcreteSalesModelFactory struct implements SalesModelFactory interface
type SalesModelFactory interface {
	CreateSales(salesID, custID int, kode string, tgl time.Time, subtotal, diskon, ongkir, total_bayar float32) *Sales
}

// ConcreteSalesModelFactory struct implements SalesModelFactory interface
type ConcreteSalesModelFactory struct{}

func (factory *ConcreteSalesModelFactory) CreateSales(salesID, custID int, kode string, tgl time.Time, subtotal, diskon, ongkir, total_bayar float32) *Sales {
	return &Sales{
		SalesID:     salesID,
		Kode:        kode,
		Tgl:         tgl,
		CustID:      custID,
		Subtotal:    subtotal,
		Diskon:      diskon,
		Ongkir:      ongkir,
		Total_bayar: total_bayar,
	}
}

// NewSalesModelFactory creates a new Sales model factory
func NewSalesModelFactory() SalesModelFactory {
	return &ConcreteSalesModelFactory{}
}
