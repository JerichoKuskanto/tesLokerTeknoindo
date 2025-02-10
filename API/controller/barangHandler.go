package controller

import (
	"log"
	"main/API/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllBarang(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	sqlStatement := "SELECT * FROM m_barang ORDER BY id;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println("Error executing query:", err)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var barangList []model.Barang
	for rows.Next() {
		var barang model.Barang
		if err := rows.Scan(&barang.BarangID, &barang.Kode, &barang.Nama, &barang.Harga); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		barangList = append(barangList, barang)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		sendErrorResponse(w, "Error retrieving barang list")
		return
	}

	if len(barangList) > 0 {
		sendSuccessResponse(w, "Successfully retrieved barangList", barangList)
	} else {
		sendErrorResponse(w, "No Barang found")
	}
}

func GetBarang(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	customerID := vars["id"]

	sqlStatement := `SELECT * FROM m_barang WHERE id = $1`

	rows, errQuery := db.Query(sqlStatement, customerID)
	if errQuery != nil {
		log.Println("Error executing query:", errQuery)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var barang model.Barang
	if rows.Next() {
		// Scan data into the Barang struct
		if err := rows.Scan(&barang.BarangID, &barang.Kode, &barang.Nama, &barang.Harga); err != nil {
			log.Println("Error scanning row:", err)
			sendErrorResponse(w, "Error reading Barang data")
			return
		}
		sendSuccessResponse(w, "Successfully retrieved barang", barang)
	} else {
		// No rows found
		log.Println("No barang found with the given ID")
		sendErrorResponse(w, "No barang found")
	}
}

func InsertBarang(w http.ResponseWriter, r *http.Request) {
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
	hargaStr := r.Form.Get("harga")                   // Get the value as a string
	hargaInt, err := strconv.ParseFloat(hargaStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting harga:", err)
		return
	}
	harga := float32(hargaInt) // Convert float64 to float32

	sqlStatement := "insert into m_barang (kode, nama, harga) values ($1, $2, $3)"

	_, errQuery := db.Exec(sqlStatement,
		kode,
		nama,
		harga,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new barang", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert barang to database")
	}
}

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	barangID := vars["id"]

	kode := r.Form.Get("kode")
	nama := r.Form.Get("nama")
	hargaStr := r.Form.Get("harga")                   // Get the value as a string
	hargaInt, err := strconv.ParseFloat(hargaStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting harga:", err)
		return
	}
	harga := float32(hargaInt) // Convert float64 to float32

	sqlStatement := `
		UPDATE m_barang
		SET kode = $1, nama = $2, harga = $3
		WHERE id = $4`

	_, errQuery := db.Exec(sqlStatement,
		kode,
		nama,
		harga,
		barangID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated barang data", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to update barang data")
	}
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	barangID := vars["id"]

	_, errQuery := db.Exec("DELETE FROM m_barang WHERE id=$1",
		barangID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted barang", nil)
	} else {
		sendErrorResponse(w, "Failed to delete barang")
	}
}
