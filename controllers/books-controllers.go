package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anujmritunjay/book-management-system/models"
	"github.com/anujmritunjay/book-management-system/utils"
	"github.com/gorilla/mux"
)

var AddBooks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	w.WriteHeader(http.StatusOK)
	res := map[string]interface{}{"success": true, "data": b}
	json.NewEncoder(w).Encode(res)
}

var GetBookById = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", ID)
		return
	}

	bookDetail, _ := models.GetBookById(ID)
	res := map[string]interface{}{"success": true, "data": bookDetail}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

var GetAllBooks = func(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res := map[string]interface{}{"success": true, "data": newBooks}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	vars := mux.Vars(r)
	bookId, _ := strconv.ParseInt(vars["bookId"], 0, 0)
	var bookDetail, db = models.GetBookById(bookId)

	utils.ParseBody(r, updateBook)

	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}
	db.Save(&bookDetail)
	res := map[string]interface{}{
		"success": true,
		"data":    bookDetail,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

var DeleteBookById = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid book ID",
		})
		return
	}

	models.DeleteBookById(bookId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Book deleted successfully",
	})
}
