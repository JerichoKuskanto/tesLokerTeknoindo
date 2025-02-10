package model

type Customer struct {
	CustomerID int    `json:"customerID"`
	Kode       string `json:"kode"`
	Nama       string `json:"nama"`
	Telp       string `json:"telp"`
}

// ConcreteCustomerModelFactory struct implements CustomerModelFactory interface
type CustomerModelFactory interface {
	CreateCustomer(customerID int, telp, kode, nama string) *Customer
}

// ConcreteCustomerModelFactory struct implements CustomerModelFactory interface
type ConcreteCustomerModelFactory struct{}

func (factory *ConcreteCustomerModelFactory) CreateCustomer(customerID int, telp, kode, nama string) *Customer {
	return &Customer{
		CustomerID: customerID,
		Kode:       kode,
		Nama:       nama,
		Telp:       telp,
	}
}

// NewCustomerModelFactory creates a new Customer model factory
func NewCustomerModelFactory() CustomerModelFactory {
	return &ConcreteCustomerModelFactory{}
}
