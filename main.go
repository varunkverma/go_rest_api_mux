package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book DS model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Init Books var as a slice of Book
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := &Book{}
	json.NewDecoder(r.Body).Decode(book)

	book.ID = strconv.Itoa(rand.Intn(1000000)) // mock id
	books = append(books, *book)
	json.NewEncoder(w).Encode(*book)

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := &Book{}
	json.NewDecoder(r.Body).Decode(newBook)
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books[i] = *newBook
			books[i].ID = params["id"]
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init the router
	r := mux.NewRouter()

	// Mock Data
	books = append(books,
		Book{
			ID:    "1",
			Isbn:  "654681",
			Title: "Book 1",
			Author: &Author{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		Book{
			ID:    "2",
			Isbn:  "786121",
			Title: "Book 2",
			Author: &Author{
				FirstName: "Steve",
				LastName:  "Smith",
			},
		},
	)

	// Route handlers
	r.HandleFunc("/api/books", getBooks).
		Methods("GET")

	r.HandleFunc("/api/books/{id}", getBook).
		Methods("GET")

	r.HandleFunc("/api/books", createBook).
		Methods("POST")

	r.HandleFunc("/api/books/{id}", updateBook).
		Methods("PUT")

	r.HandleFunc("/api/books/{id}", deleteBook).
		Methods("DELETE")

	// http server
	log.Fatal(http.ListenAndServe(":8000", r))
}
