package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	/*
		{
			"error": "message"
		}
	*/

	respondJSON(w, code, errorResponse{Error: message})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON response: ", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
