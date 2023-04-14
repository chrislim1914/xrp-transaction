package handler

import (
	"encoding/json"
	"net/http"
)

func jsonResponseHandler(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(data)
}

func jsonErrorResponseHandler(w http.ResponseWriter, errMsg string, statuscode int) {
	newErr := make(map[string]string)
	newErr["error"] = errMsg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(newErr)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse := make(map[string]string)
	jsonResponse["message"] = "API is working..."
	jsonResponseHandler(w, jsonResponse, 200)
}
