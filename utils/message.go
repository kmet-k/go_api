package utils

import (
	"free_credit/models"

	"encoding/json"
	"net/http"
)

func SentMessage(response http.ResponseWriter, message models.Constants) {
	if message.Message == "invalid syntax" {
		response.WriteHeader(http.StatusBadRequest)
	} else if message.Message == "token not found" {
		response.WriteHeader(http.StatusNotFound)
	} else if message.Message == "invalid token" {
		response.WriteHeader(http.StatusUnauthorized)
	}
	json.NewEncoder(response).Encode(message)
}
