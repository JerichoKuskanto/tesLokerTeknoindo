package model

type Barang struct {
	BarangID int     `json:"barangID"`
	Kode     string  `json:"kode"`
	Nama     string  `json:"nama"`
	Harga    float32 `json:"harga"`
}

// ConcreteBarangModelFactory struct implements BarangModelFactory interface
type BarangModelFactory interface {
	CreateBarang(barangID int, harga float32, kode, nama string) *Barang
}

// ConcreteBarangModelFactory struct implements BarangModelFactory interface
type ConcreteBarangModelFactory struct{}

func (factory *ConcreteBarangModelFactory) CreateBarang(barangID int, harga float32, kode, nama string) *Barang {
	return &Barang{
		BarangID: barangID,
		Kode:     kode,
		Nama:     nama,
		Harga:    harga,
	}
}

// NewBarangModelFactory creates a new Barang model factory
func NewBarangModelFactory() BarangModelFactory {
	return &ConcreteBarangModelFactory{}
}
