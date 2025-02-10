package controller

import (
	"log"
	"main/API/model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	sqlStatement := "SELECT * FROM m_customer ORDER BY id;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println("Error executing query:", err)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var customerList []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.CustomerID, &customer.Kode, &customer.Nama, &customer.Telp); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		customerList = append(customerList, customer)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		sendErrorResponse(w, "Error retrieving customer list")
		return
	}

	if len(customerList) > 0 {
		sendSuccessResponse(w, "Successfully retrieved customerList", customerList)
	} else {
		sendErrorResponse(w, "No customer found")
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	customerID := vars["id"]

	sqlStatement := `SELECT * FROM m_customer WHERE id = $1`

	rows, errQuery := db.Query(sqlStatement, customerID)
	if errQuery != nil {
		log.Println("Error executing query:", errQuery)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var customer model.Customer
	if rows.Next() {
		// Scan data into the customer struct
		if err := rows.Scan(&customer.CustomerID, &customer.Kode, &customer.Nama, &customer.Telp); err != nil {
			log.Println("Error scanning row:", err)
			sendErrorResponse(w, "Error reading customer data")
			return
		}
		sendSuccessResponse(w, "Successfully retrieved customer", customer)
	} else {
		// No rows found
		log.Println("No customer found with the given ID")
		sendErrorResponse(w, "No customer found")
	}
}

func InsertCustomer(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	kode := r.Form.Get("kode")
	nama := r.Form.Get("nama")
	telp := r.Form.Get("telp")

	sqlStatement := "insert into m_customer (kode, name, telp) values ($1, $2, $3)"

	_, errQuery := db.Exec(sqlStatement,
		kode,
		nama,
		telp,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new customer", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert customer to database")
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	customerID := vars["id"]

	kode := r.Form.Get("kode")
	nama := r.Form.Get("nama")
	telp := r.Form.Get("telp")

	sqlStatement := `
		UPDATE m_customer
		SET kode = $1, name = $2, telp = $3
		WHERE id = $4`

	_, errQuery := db.Exec(sqlStatement,
		kode,
		nama,
		telp,
		customerID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated customer data", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to update customer data")
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	customerID := vars["id"]

	_, errQuery := db.Exec("DELETE FROM m_customer WHERE id=$1",
		customerID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted customer", nil)
	} else {
		sendErrorResponse(w, "Failed to delete customer")
	}
}
