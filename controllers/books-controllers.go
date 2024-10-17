package controllers

import (
	"encoding/json"
	"net/http"
)

var AddBooks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Hello from the Add Book Controllers"})
}
