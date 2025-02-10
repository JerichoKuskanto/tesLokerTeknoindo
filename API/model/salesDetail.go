package model

type SalesDetail struct {
	Sales_ID      int     `json:"sales_ID"`
	Barang_ID     int     `json:"barang_ID"`
	Harga_bandrol float32 `json:"harga_bandrol"`
	Qty           int     `json:"qty"`
	Diskon_pct    float32 `json:"diskon_pct"`
	Diskon_nilai  float32 `json:"diskon_nilai"`
	Harga_diskon  float32 `json:"harga_diskon"`
	Total         float32 `json:"total"`
}

// ConcreteSalesDetailModelFactory struct implements SalesDetailModelFactory interface
type SalesDetailModelFactory interface {
	CreateSalesDetail(sales_ID, barang_ID, qty int, harga_bandrol, diskon_pct, diskon_nilai, Harga_diskon, total float32) *SalesDetail
}

// ConcreteSalesDetailModelFactory struct implements SalesDetailModelFactory interface
type ConcreteSalesDetailModelFactory struct{}

func (factory *ConcreteSalesDetailModelFactory) CreateSalesDetail(sales_ID, barang_ID, qty int, harga_bandrol, diskon_pct, diskon_nilai, Harga_diskon, total float32) *SalesDetail {
	return &SalesDetail{
		Sales_ID:      sales_ID,
		Barang_ID:     barang_ID,
		Harga_bandrol: harga_bandrol,
		Qty:           qty,
		Diskon_pct:    diskon_pct,
		Diskon_nilai:  diskon_nilai,
		Harga_diskon:  Harga_diskon,
		Total:         total,
	}
}

// NewSalesDetailModelFactory creates a new SalesDetail model factory
func NewSalesDetailModelFactory() SalesDetailModelFactory {
	return &ConcreteSalesDetailModelFactory{}
}
