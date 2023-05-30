package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SemaMolchanov/go-bookstore/packages/models"
	"github.com/SemaMolchanov/go-bookstore/packages/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(r, book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if book.Name != "" {
		bookDetails.Name = book.Name
	}

	if book.Author != "" {
		bookDetails.Author = book.Author
	}

	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
