package controllers

import (
	"net/http"
	"strconv"

	"github.com/definev/go_freecodecamp/go-bookstore/pkg/models"
	"github.com/definev/go_freecodecamp/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.StandardizedResponse(books, "", ""))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(&book, r)
	b := book.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.StandardizedResponse(b, "", ""))
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(utils.StandardizedResponse(map[string]string{}, "01", "Book id is wrong!"))
		panic(err)
	}
	_, bookDetails := models.GetBookById(bookId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.StandardizedResponse(bookDetails, "", ""))
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.StandardizedResponse(map[string]string{}, "01", "Book id is wrong!"))
		panic(err)
	}
	updateBook := &models.Book{}
	utils.ParseBody(updateBook, r)
	db, bookDetails := models.GetBookById(bookId)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.StandardizedResponse(bookDetails, "", ""))
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(utils.StandardizedResponse(map[string]string{}, "01", "Book id is wrong!"))
		panic(err)
	}
	book := models.DeleteBook(bookId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.StandardizedResponse(book, "", ""))
}
