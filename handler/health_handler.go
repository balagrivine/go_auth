package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func HandleHealth() HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		response := map[string]string{"message": "Status OK"}
		json.NewEncoder(w).Encode(response)
	}
}
