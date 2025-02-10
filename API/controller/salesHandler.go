package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"main/API/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func GetAllSales(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	sqlStatement := "SELECT * FROM t_sales ORDER BY id;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println("Error executing query:", err)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var salesList []model.Sales
	for rows.Next() {
		var sales model.Sales
		if err := rows.Scan(&sales.SalesID, &sales.Kode, &sales.Tgl, &sales.CustID, &sales.Subtotal,
			&sales.Diskon, &sales.Ongkir, &sales.Total_bayar); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		salesList = append(salesList, sales)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		sendErrorResponse(w, "Error retrieving sales list")
		return
	}

	if len(salesList) > 0 {
		sendSuccessResponse(w, "Successfully retrieved salesList", salesList)
	} else {
		sendErrorResponse(w, "No sales found")
	}
}

func GetSales(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	salesID := vars["id"]

	sqlStatement := `SELECT * FROM t_sales WHERE id = $1`

	rows, errQuery := db.Query(sqlStatement, salesID)
	if errQuery != nil {
		log.Println("Error executing query:", errQuery)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var sales model.Sales
	if rows.Next() {
		// Scan data into the sales struct
		if err := rows.Scan(&sales.SalesID, &sales.Kode, &sales.Tgl, &sales.CustID, &sales.Subtotal,
			&sales.Diskon, &sales.Ongkir, &sales.Total_bayar); err != nil {
			log.Println("Error scanning row:", err)
			sendErrorResponse(w, "Error reading sales data")
			return
		}
		sendSuccessResponse(w, "Successfully retrieved sales", sales)
	} else {
		// No rows found
		log.Println("No sales found with the given ID")
		sendErrorResponse(w, "No sales found")
	}
}

func InsertSales(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse request")
		return
	}

	// Generate new kode
	kode, err := GenerateKode(db)
	if err != nil {
		sendErrorResponse(w, "Failed to generate kode")
		return
	}

	tgl := r.Form.Get("tgl")
	custID, _ := strconv.Atoi(r.Form.Get("custID"))

	subtotal, err := strconv.ParseFloat(r.Form.Get("subtotal"), 32)
	if err != nil {
		sendErrorResponse(w, "Invalid subtotal")
		return
	}

	diskon, err := strconv.ParseFloat(r.Form.Get("diskon"), 32)
	if err != nil {
		sendErrorResponse(w, "Invalid diskon")
		return
	}

	ongkir, err := strconv.ParseFloat(r.Form.Get("ongkir"), 32)
	if err != nil {
		sendErrorResponse(w, "Invalid ongkir")
		return
	}

	total, err := strconv.ParseFloat(r.Form.Get("total"), 32)
	if err != nil {
		sendErrorResponse(w, "Invalid total")
		return
	}

	// Insert into database and return inserted ID
	sqlStatement := `
		INSERT INTO t_sales (kode, tgl, cust_id, subtotal, diskon, ongkir, total_bayar) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	var salesID int
	err = db.QueryRow(sqlStatement, kode, tgl, custID, float32(subtotal), float32(diskon), float32(ongkir), float32(total)).Scan(&salesID)
	if err != nil {
		sendErrorResponse(w, "Failed to insert sales into the database")
		return
	}

	// ✅ Ensure the response contains the salesID
	response := map[string]interface{}{
		"success": true,
		"message": "Successfully inserted new sales",
		"salesID": salesID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // ✅ Send JSON response
}

func GenerateKode(db *sqlx.DB) (string, error) {
	currentYearMonth := time.Now().Format("200601") // YYYYMM format
	var lastNumber int

	// Find the highest sequence number for the current month
	query := `SELECT COALESCE(MAX(SUBSTRING(kode FROM 8)::INTEGER), 0) 
	          FROM t_sales WHERE SUBSTRING(kode FROM 1 FOR 6) = $1`
	err := db.QueryRow(query, currentYearMonth).Scan(&lastNumber)
	if err != nil {
		return "", err
	}

	// Increment or reset the counter
	nextNumber := lastNumber + 1

	// Format the kode as "YYYYMM-XXXX"
	newKode := fmt.Sprintf("%s-%04d", currentYearMonth, nextNumber)
	return newKode, nil
}

func UpdateSales(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse form data")
		return
	}

	vars := mux.Vars(r)
	salesID := vars["id"]

	// Fetch existing sales data before updating
	var existingKode, existingTgl string
	var existingCustID int
	var existingSubtotal, existingDiskon, existingOngkir, existingTotal float32

	query := "SELECT kode, tgl, cust_id, subtotal, diskon, ongkir, total_bayar FROM t_sales WHERE id = $1"
	err = db.QueryRow(query, salesID).Scan(&existingKode, &existingTgl, &existingCustID, &existingSubtotal, &existingDiskon, &existingOngkir, &existingTotal)
	if err != nil {
		log.Println("Error fetching existing sales data:", err)
		sendErrorResponse(w, "Failed to fetch existing sales data")
		return
	}

	// Get new values from request
	kode := r.Form.Get("kode")
	tgl := r.Form.Get("tgl")
	custIDStr := r.Form.Get("custID")
	subtotalStr := r.Form.Get("subtotal")
	diskonStr := r.Form.Get("diskon")
	ongkirStr := r.Form.Get("ongkir")
	totalStr := r.Form.Get("total")

	// Keep existing values if new ones are empty
	if kode == "" {
		kode = existingKode
	}
	if tgl == "" {
		tgl = existingTgl
	}

	// Convert integers (custID)
	custID := existingCustID
	if custIDStr != "" {
		parsedCustID, err := strconv.Atoi(custIDStr)
		if err != nil {
			log.Println("Error converting custID:", err)
			sendErrorResponse(w, "Invalid custID format")
			return
		}
		custID = parsedCustID
	}

	// Convert float values only if they are not empty
	subtotal := existingSubtotal
	if subtotalStr != "" {
		parsedSubtotal, err := strconv.ParseFloat(subtotalStr, 32)
		if err != nil {
			log.Println("Error converting subtotal:", err)
			sendErrorResponse(w, "Invalid subtotal format")
			return
		}
		subtotal = float32(parsedSubtotal)
	}

	diskon := existingDiskon
	if diskonStr != "" {
		parsedDiskon, err := strconv.ParseFloat(diskonStr, 32)
		if err != nil {
			log.Println("Error converting diskon:", err)
			sendErrorResponse(w, "Invalid diskon format")
			return
		}
		diskon = float32(parsedDiskon)
	}

	ongkir := existingOngkir
	if ongkirStr != "" {
		parsedOngkir, err := strconv.ParseFloat(ongkirStr, 32)
		if err != nil {
			log.Println("Error converting ongkir:", err)
			sendErrorResponse(w, "Invalid ongkir format")
			return
		}
		ongkir = float32(parsedOngkir)
	}

	total := existingTotal
	if totalStr != "" {
		parsedTotal, err := strconv.ParseFloat(totalStr, 32)
		if err != nil {
			log.Println("Error converting total:", err)
			sendErrorResponse(w, "Invalid total format")
			return
		}
		total = float32(parsedTotal)
	}

	// Update query
	sqlStatement := `
		UPDATE t_sales
		SET kode = $1, tgl = $2, cust_id = $3, subtotal = $4, diskon = $5, ongkir = $6, total_bayar = $7
		WHERE id = $8`

	_, errQuery := db.Exec(sqlStatement, kode, tgl, custID, subtotal, diskon, ongkir, total, salesID)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated sales data", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to update sales data")
	}
}

func DeleteSales(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	salesID := vars["id"]

	_, errQuery := db.Exec("DELETE FROM t_sales WHERE id=$1",
		salesID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted sales", nil)
	} else {
		sendErrorResponse(w, "Failed to delete sales")
	}
}
