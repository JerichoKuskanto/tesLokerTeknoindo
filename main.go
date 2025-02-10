package main

import (
	"fmt"
	"log"
	"main/API/controller"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Handlers
	router.HandleFunc("/barang", controller.GetAllBarang).Methods("GET")
	router.HandleFunc("/barang/{id}", controller.GetBarang).Methods("GET")
	router.HandleFunc("/barang", controller.InsertBarang).Methods("POST")
	router.HandleFunc("/barang/{id}", controller.UpdateBarang).Methods("PUT")
	router.HandleFunc("/barang/{id}", controller.DeleteBarang).Methods("DELETE")

	router.HandleFunc("/customer", controller.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.GetCustomer).Methods("GET")
	router.HandleFunc("/customer", controller.InsertCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", controller.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", controller.DeleteCustomer).Methods("DELETE")

	router.HandleFunc("/sales", controller.GetAllSales).Methods("GET")
	router.HandleFunc("/sales/{id}", controller.GetSales).Methods("GET")
	router.HandleFunc("/sales", controller.InsertSales).Methods("POST")
	router.HandleFunc("/sales/{id}", controller.UpdateSales).Methods("PUT")
	router.HandleFunc("/sales/{id}", controller.DeleteSales).Methods("DELETE")

	router.HandleFunc("/salesDet", controller.GetAllSalesDetail).Methods("GET")
	router.HandleFunc("/salesDetSpec/{id}", controller.GetAllSalesDetailSpecific).Methods("GET")
	router.HandleFunc("/salesDet/{id}", controller.GetSalesDetail).Methods("GET")
	router.HandleFunc("/salesDet", controller.InsertSalesDetail).Methods("POST")
	router.HandleFunc("/salesDet/{id}", controller.UpdateSalesDetail).Methods("PUT")
	router.HandleFunc("/salesDet/{id}", controller.DeleteSalesDetail).Methods("DELETE")

	// CORS Configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8181"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	// 🔹 Set port dynamically using an environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8181 if PORT is not set
	}

	fmt.Println("Connected to port", port)
	log.Println("Connected to port", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
