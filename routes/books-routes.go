package routes

import (
	"github.com/anujmritunjay/book-management-system/controllers"
	"github.com/gorilla/mux"
)

var BooksRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.AddBooks).Methods("POST")
}
