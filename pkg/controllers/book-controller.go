package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CAndresFernandez/go-books/pkg/models"
	"github.com/CAndresFernandez/go-books/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// returns json response of the slice of books in the db
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// call the GetAllBooks function from models package
	newBooks := models.GetAllBooks()
	// set the response as a json of the slice of books
	res, _ := json.Marshal(newBooks)
	// set the header type
	w.Header().Set("Content-Type", "pkglication/json")
	// http response
	w.WriteHeader(http.StatusOK)
	// return slice of books from db
	w.Write(res)
}

// returns json response of a book by its id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// returns route variables for the sent request
	vars := mux.Vars(r)
	// define the variable
	id := vars["id"]
	// converts string to int64
	ID, err := strconv.ParseInt(id, 0, 0)
	// handles errors while parsing, if there are any
	if err != nil {
		fmt.Println("error while parsing")
	}
	// recuperate the details from the db
	bookDetails, _:= models.GetBookById(ID)
	// set the response as json of the book details
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	// return json response of the book
	w.Write(res)
}

// create a book
func CreateBook(w http.ResponseWriter, r*http.Request) {
	// set a variable that points to the type of object to be created : Book
	CreateBook := &models.Book{}
	// parse the user request (received in json) into something the db comprehends
	utils.ParseBody(r, CreateBook)
	// b is the result of the creation in the db, returned by the CreateBook function in the models package
	b := CreateBook.CreateBook()
	// convert back to json for the response
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// delete a book from the db
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// update a book in the db by its id
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
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
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}