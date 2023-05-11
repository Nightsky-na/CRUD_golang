// CRUD_API_USING_GO ref: https://github.com/BorntoDev/CRUD-Golang/
package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux" // go get -u github.com/gorilla/mux
)

// Create Struct
type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	} 
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	
	_ = json.NewDecoder(r.Body).Decode(&book) // Decode JSON to book struct
	
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book) // Decode JSON to book struct
			book.ID = param["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return 
		}
	}	
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}	
	json.NewEncoder(w).Encode(books)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	
	// Mock data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", r))
}