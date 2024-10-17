package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anujmritunjay/book-management-system/config"
	"github.com/anujmritunjay/book-management-system/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	r.HandleFunc("/", rootRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", r))
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{"success": true, "message": "Hello from the Go Server."}
	json.NewEncoder(w).Encode(res)

}
