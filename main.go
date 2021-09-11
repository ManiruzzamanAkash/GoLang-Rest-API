package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct Schema
type Book struct {
	ID     string  `json:"id`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Book Author Schema
type Author struct {
	FirstName string `json:firstname`
	LastName  string `json:lastname`
}

// Init books var as slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/**
* Main function to handle all of the routes...
 */
func main() {
	r := mux.NewRouter()

	// Mock Data - @todo implement DB
	books = append(books, Book{ID: "1", Isbn: "348934", Title: "Book One", Author: &Author{FirstName: "Jhon", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "901209", Title: "Book Two", Author: &Author{FirstName: "Polash", LastName: "Rana"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
