package controller

import (
	"encoding/json"
	"log"
	"main/API/model"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content=Type", "application/json")
	responseFactory := model.NewErrorResponseModelFactory()
	response := responseFactory.CreateErrorResponse(message)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendSuccessResponse(w http.ResponseWriter, message string, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseFactory := model.NewSuccessResponseModelFactory()
	response := responseFactory.CreateSuccessResponse(
		message,
		value,
	)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
