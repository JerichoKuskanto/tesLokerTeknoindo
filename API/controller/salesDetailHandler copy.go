package controller

import (
	"log"
	"main/API/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllSalesDetail(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	sqlStatement := "SELECT * FROM t_sales_det ORDER BY sales_id;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println("Error executing query:", err)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var salesDetailList []model.SalesDetail
	for rows.Next() {
		var salesDetail model.SalesDetail
		if err := rows.Scan(&salesDetail.Sales_ID, &salesDetail.Barang_ID, &salesDetail.Harga_bandrol, &salesDetail.Qty,
			&salesDetail.Diskon_pct, &salesDetail.Diskon_nilai, &salesDetail.Harga_diskon, &salesDetail.Total); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		salesDetailList = append(salesDetailList, salesDetail)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		sendErrorResponse(w, "Error retrieving sales detail list")
		return
	}

	if len(salesDetailList) > 0 {
		sendSuccessResponse(w, "Successfully retrieved salesDetail List", salesDetailList)
	} else {
		sendErrorResponse(w, "No salesDetail found")
	}
}

func GetAllSalesDetailSpecific(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse salesID from mux vars
	vars := mux.Vars(r)
	salesID, exists := vars["id"]
	if !exists || salesID == "" {
		log.Println("Missing or invalid sales_id parameter") // 🔍 Debugging log
		sendErrorResponse(w, "Missing or invalid sales_id parameter")
		return
	}

	// ✅ Convert salesID to integer safely
	idInt, err := strconv.Atoi(salesID)
	if err != nil {
		log.Println("Invalid sales_id format:", salesID) // 🔍 Debugging log
		sendErrorResponse(w, "Invalid sales_id format")
		return
	}

	// Prepare SQL statement with WHERE clause
	sqlStatement := "SELECT * FROM t_sales_det WHERE sales_id = $1 ORDER BY barang_id;"

	rows, err := db.Query(sqlStatement, idInt) // ✅ Use converted integer
	if err != nil {
		log.Println("Error executing query:", err)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var salesDetailList []model.SalesDetail
	for rows.Next() {
		var salesDetail model.SalesDetail
		if err := rows.Scan(&salesDetail.Sales_ID, &salesDetail.Barang_ID, &salesDetail.Harga_bandrol, &salesDetail.Qty,
			&salesDetail.Diskon_pct, &salesDetail.Diskon_nilai, &salesDetail.Harga_diskon, &salesDetail.Total); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		salesDetailList = append(salesDetailList, salesDetail)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		sendErrorResponse(w, "Error retrieving sales detail list")
		return
	}

	// ✅ Return empty array instead of error when no data is found
	sendSuccessResponse(w, "Successfully retrieved salesDetail List", salesDetailList)
}

func GetSalesDetail(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	salesDetailID := vars["id"]

	sqlStatement := `SELECT * FROM t_sales_det WHERE id = $1`

	rows, errQuery := db.Query(sqlStatement, salesDetailID)
	if errQuery != nil {
		log.Println("Error executing query:", errQuery)
		sendErrorResponse(w, "Failed to execute query")
		return
	}
	defer rows.Close()

	var salesDetail model.SalesDetail
	if rows.Next() {
		// Scan data into the sales struct
		if err := rows.Scan(&salesDetail.Sales_ID, &salesDetail.Barang_ID, &salesDetail.Harga_bandrol, &salesDetail.Qty,
			&salesDetail.Diskon_pct, &salesDetail.Diskon_nilai, &salesDetail.Harga_diskon, &salesDetail.Total); err != nil {
			log.Println("Error scanning row:", err)
			sendErrorResponse(w, "Error reading sales detail data")
			return
		}
		sendSuccessResponse(w, "Successfully retrieved sales deta", salesDetail)
	} else {
		// No rows found
		log.Println("No sales detail found with the given ID")
		sendErrorResponse(w, "No sales detail found")
	}
}

func InsertSalesDetail(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	salesID, _ := strconv.Atoi(r.Form.Get("salesID"))
	barangID, _ := strconv.Atoi(r.Form.Get("barangID"))

	hargaBandrolStr := r.Form.Get("hargaBandrol")                   // Get the value as a string
	hargaBandrolInt, err := strconv.ParseFloat(hargaBandrolStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting hargaBandrol:", err)
		return
	}
	hargaBandrol := float32(hargaBandrolInt) // Convert float64 to float32

	qty, _ := strconv.Atoi(r.Form.Get("qty"))

	diskonPctStr := r.Form.Get("diskonPct")                   // Get the value as a string
	diskonPctInt, err := strconv.ParseFloat(diskonPctStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting diskonPct:", err)
		return
	}
	diskonPct := float32(diskonPctInt) // Convert float64 to float32

	diskonNilaiStr := r.Form.Get("diskonNilai")                   // Get the value as a string
	diskonNilaiInt, err := strconv.ParseFloat(diskonNilaiStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting diskonNilai:", err)
		return
	}
	diskonNilai := float32(diskonNilaiInt) // Convert float64 to float32

	hargaDiskonStr := r.Form.Get("hargaDiskon")                   // Get the value as a string
	hargaDiskonInt, err := strconv.ParseFloat(hargaDiskonStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting hargaDiskon:", err)
		return
	}
	hargaDiskon := float32(hargaDiskonInt) // Convert float64 to float32

	totalStr := r.Form.Get("total")                   // Get the value as a string
	totalInt, err := strconv.ParseFloat(totalStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting total:", err)
		return
	}
	total := float32(totalInt) // Convert float64 to float32

	sqlStatement := `insert into t_sales_det (sales_id, barang_id, harga_bandrol, qty, diskon_pct, diskon_nilai, harga_diskon, total) 
	values ($1, $2, $3,$4, $5, $6, $7, $8)`

	_, errQuery := db.Exec(sqlStatement,
		salesID,
		barangID,
		hargaBandrol,
		qty,
		diskonPct,
		diskonNilai,
		hargaDiskon,
		total,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new sales Detail", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert sales detail to database")
	}
}

func InsertSalesDetailComplete(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	// Input sales
	// Call GenerateKode function to get the new kode
	kode, err := GenerateKode(db)
	if err != nil {
		log.Println("Error generating kode:", err)
		sendErrorResponse(w, "Failed to generate kode")
		return
	}

	tgl := r.Form.Get("tgl")
	custID, _ := strconv.Atoi(r.Form.Get("custID"))

	subtotalStr := r.Form.Get("subtotal")
	subtotalInt, err := strconv.ParseFloat(subtotalStr, 32)
	if err != nil {
		log.Println("Error converting subtotal:", err)
		sendErrorResponse(w, "Invalid subtotal")
		return
	}
	subtotal := float32(subtotalInt)

	diskonStr := r.Form.Get("diskon")
	diskonInt, err := strconv.ParseFloat(diskonStr, 32)
	if err != nil {
		log.Println("Error converting diskon:", err)
		sendErrorResponse(w, "Invalid diskon")
		return
	}
	diskon := float32(diskonInt)

	ongkirStr := r.Form.Get("ongkir")
	ongkirInt, err := strconv.ParseFloat(ongkirStr, 32)
	if err != nil {
		log.Println("Error converting ongkir:", err)
		sendErrorResponse(w, "Invalid ongkir")
		return
	}
	ongkir := float32(ongkirInt)

	totalStr := r.Form.Get("total")
	totalInt, err := strconv.ParseFloat(totalStr, 32)
	if err != nil {
		log.Println("Error converting total:", err)
		sendErrorResponse(w, "Invalid total")
		return
	}
	total := float32(totalInt)

	// Insert into database
	sqlStatement := "INSERT INTO t_sales (kode, tgl, cust_id, subtotal, diskon, ongkir, total_bayar) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	var salesID int
	errQuery := db.QueryRow(sqlStatement,
		kode,
		tgl,
		custID,
		subtotal,
		diskon,
		ongkir,
		total,
	).Scan(&salesID)

	if errQuery != nil {
		log.Println("Database error:", errQuery)
		sendErrorResponse(w, "Failed to insert sales into the database")
	}

	//input detail sales
	barangID, _ := strconv.Atoi(r.Form.Get("barangID"))

	hargaBandrolStr := r.Form.Get("hargaBandrol")                   // Get the value as a string
	hargaBandrolInt, err := strconv.ParseFloat(hargaBandrolStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting hargaBandrol:", err)
		return
	}
	hargaBandrol := float32(hargaBandrolInt) // Convert float64 to float32

	qty, _ := strconv.Atoi(r.Form.Get("qty"))

	diskonPctStr := r.Form.Get("diskonPct")                   // Get the value as a string
	diskonPctInt, err := strconv.ParseFloat(diskonPctStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting diskonPct:", err)
		return
	}
	diskonPct := float32(diskonPctInt) // Convert float64 to float32

	diskonNilaiStr := r.Form.Get("diskonNilai")                   // Get the value as a string
	diskonNilaiInt, err := strconv.ParseFloat(diskonNilaiStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting diskonNilai:", err)
		return
	}
	diskonNilai := float32(diskonNilaiInt) // Convert float64 to float32

	hargaDiskonStr := r.Form.Get("hargaDiskon")                   // Get the value as a string
	hargaDiskonInt, err := strconv.ParseFloat(hargaDiskonStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting hargaDiskon:", err)
		return
	}
	hargaDiskon := float32(hargaDiskonInt) // Convert float64 to float32

	totalStr = r.Form.Get("total")                   // Get the value as a string
	totalInt, err = strconv.ParseFloat(totalStr, 32) // Convert to float64 with precision 32
	if err != nil {
		log.Println("Error converting total:", err)
		return
	}
	total = float32(totalInt) // Convert float64 to float32

	sqlStatement = `insert into t_sales_det (sales_id, barang_id, harga_bandrol, qty, diskon_pct, diskon_nilai, harga_diskon, total) 
	values ($1, $2, $3,$4, $5, $6, $7, $8)`

	_, errQuery = db.Exec(sqlStatement,
		salesID,
		barangID,
		hargaBandrol,
		qty,
		diskonPct,
		diskonNilai,
		hargaDiskon,
		total,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new sales Detail", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert sales detail to database")
	}

}

func UpdateSalesDetail(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse form data")
		return
	}

	vars := mux.Vars(r)
	salesID := vars["id"]

	barangIDStr := r.Form.Get("barangID")
	if barangIDStr == "" {
		sendErrorResponse(w, "barangID is required")
		return
	}
	barangID, err := strconv.Atoi(barangIDStr)
	if err != nil {
		log.Println("Error converting barangID:", err)
		sendErrorResponse(w, "Invalid barangID format")
		return
	}

	// Fetch existing sales detail data before updating
	var existingQty int
	var existingHargaBandrol, existingDiskonPct, existingDiskonNilai, existingHargaDiskon, existingTotal float32

	query := "SELECT qty, harga_bandrol, diskon_pct, diskon_nilai, harga_diskon, total FROM t_sales_det WHERE sales_id = $1 AND barang_id = $2"
	err = db.QueryRow(query, salesID, barangID).Scan(&existingQty, &existingHargaBandrol, &existingDiskonPct, &existingDiskonNilai, &existingHargaDiskon, &existingTotal)
	if err != nil {
		log.Println("Error fetching existing sales detail:", err)
		sendErrorResponse(w, "Failed to fetch existing sales detail data")
		return
	}

	// Get new values from request
	qtyStr := r.Form.Get("qty")
	hargaBandrolStr := r.Form.Get("hargaBandrol")
	diskonPctStr := r.Form.Get("diskonPct")
	diskonNilaiStr := r.Form.Get("diskonNilai")
	hargaDiskonStr := r.Form.Get("hargaDiskon")
	totalStr := r.Form.Get("total")

	// Preserve existing values if new ones are empty
	qty := existingQty
	if qtyStr != "" {
		parsedQty, err := strconv.Atoi(qtyStr)
		if err != nil {
			log.Println("Error converting qty:", err)
			sendErrorResponse(w, "Invalid qty format")
			return
		}
		qty = parsedQty
	}

	hargaBandrol := existingHargaBandrol
	if hargaBandrolStr != "" {
		parsedHargaBandrol, err := strconv.ParseFloat(hargaBandrolStr, 32)
		if err != nil {
			log.Println("Error converting hargaBandrol:", err)
			sendErrorResponse(w, "Invalid hargaBandrol format")
			return
		}
		hargaBandrol = float32(parsedHargaBandrol)
	}

	diskonPct := existingDiskonPct
	if diskonPctStr != "" {
		parsedDiskonPct, err := strconv.ParseFloat(diskonPctStr, 32)
		if err != nil {
			log.Println("Error converting diskonPct:", err)
			sendErrorResponse(w, "Invalid diskonPct format")
			return
		}
		diskonPct = float32(parsedDiskonPct)
	}

	diskonNilai := existingDiskonNilai
	if diskonNilaiStr != "" {
		parsedDiskonNilai, err := strconv.ParseFloat(diskonNilaiStr, 32)
		if err != nil {
			log.Println("Error converting diskonNilai:", err)
			sendErrorResponse(w, "Invalid diskonNilai format")
			return
		}
		diskonNilai = float32(parsedDiskonNilai)
	}

	hargaDiskon := existingHargaDiskon
	if hargaDiskonStr != "" {
		parsedHargaDiskon, err := strconv.ParseFloat(hargaDiskonStr, 32)
		if err != nil {
			log.Println("Error converting hargaDiskon:", err)
			sendErrorResponse(w, "Invalid hargaDiskon format")
			return
		}
		hargaDiskon = float32(parsedHargaDiskon)
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
		UPDATE t_sales_det
		SET harga_bandrol = $1, qty = $2, diskon_pct = $3, diskon_nilai = $4, harga_diskon = $5, total = $6
		WHERE sales_id = $7 AND barang_id = $8`

	_, errQuery := db.Exec(sqlStatement, hargaBandrol, qty, diskonPct, diskonNilai, hargaDiskon, total, salesID, barangID)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated sales detail data", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to update sales detail data")
	}
}

func UpdateQtySalesDetail(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	salesID := vars["salesid"]
	barangID, _ := strconv.Atoi(r.Form.Get("barangID"))
	qty, _ := strconv.Atoi(r.Form.Get("qty"))

	sqlStatement := `
		UPDATE t_sales_det
		SET qty = $1
		WHERE sales_id = $2 AND barang_id = $3`

	_, errQuery := db.Exec(sqlStatement,
		qty,
		salesID,
		barangID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated sales detail qty data", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to update sales detail qty data")
	}
}

func DeleteSalesDetail(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	salesID := vars["id"]

	_, errQuery := db.Exec("DELETE FROM t_sales_det WHERE sales_id=$1",
		salesID,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted sales detail", nil)
	} else {
		sendErrorResponse(w, "Failed to delete sales detail")
	}
}
