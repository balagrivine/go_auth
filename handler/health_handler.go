package handler

import (
	"net/http"
	// "fmt"
	"encoding/json"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, {"message": "StatusOK"})
}
