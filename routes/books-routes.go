package routes

import (
	"github.com/anujmritunjay/book-management-system/controllers"
	"github.com/gorilla/mux"
)

var BooksRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.AddBooks).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/update/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/delete/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
