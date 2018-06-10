package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author, and ISBN
type Book struct {
	Title       string `json:"title"` // Marshall as lowercase
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// Declare dictionary in-memory data structure
// key (isbn) => value (individual books)
var books = map[string]Book{
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"},
}

// ToJSON returns byte array of marshalled data
func (b Book) ToJSON() []byte {
	// Returns JSON encoding of book
	toJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return toJSON
}

// FromJSON returns unmarshalled book struct
func FromJSON(data []byte) Book {
	// Unmarshal data parameter received
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// Books is a slice of all known books
var Books = []Book{
	Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"},
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	// API request for /api/book
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest) // Return invalid status code
		w.Write([]byte("Unsupported request type."))

	} // Write back
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	// API request for /api/book/<isbn>
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	return Books
}

// Write books in JSON format as a bytearray to ResponseWriter
func writeJSON(w http.ResponseWriter, books []Book) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8") // Tell client about JSON data
	w.Write(b)
}
