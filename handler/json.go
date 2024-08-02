package handler

import (
	"net/http"
	"encoding/json"
	"log"
)

// RespondWithJSON: sends a JSON response with given HTTP status code and payload
// w: http.ResponseWriter to write response to
// statusCode: HTTP status code to set ofr the response
// payload: data to be marshalled into JSON and sent as the response body

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Write JSON data to the response body
	w.Write(data)
}
