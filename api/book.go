package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Book type with Name, Author, and ISBN
type Book struct {
	Title       string `json:"title"` // Marshall as lowercase
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// Books declares dictionary in-memory data structure
// key (isbn) => value (individual books)
var Books = map[string]Book{
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

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	// API request for /api/book
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book) // Create the book
		if created {
			w.Header().Add("Location", "/api/books/"+isbn) // Add Location header, append isbn
			w.WriteHeader(http.StatusCreated)              // Return HTTP 201
		} else {
			w.WriteHeader(http.StatusConflict) // Return HTTP 409
		}
	default:
		w.WriteHeader(http.StatusBadRequest) // Return invalid status code
		w.Write([]byte("Unsupported request type."))

	} // Write back
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):] // Extract isbn as last part of path
	// Switch on HTTP method
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound) // Return HTTP 404
		}
	}
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	// Convert map of books to books
	all := make([]Book, 0, len(Books))
	for _, val := range Books {
		all = append(all, val)
	}
	return all
}

// Write books in JSON format as a bytearray to ResponseWriter
func writeJSON(w http.ResponseWriter, books []Book) {
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8") // Tell client about JSON data
	w.Write(b)
}

// CreateBook creates a book if it does not exist
func CreateBook(b Book) (string, bool) {
	// Check if ISBN exists in map
	if _, exists := Books[b.ISBN]; exists {
		return "", false
	}
	// Add new book to map
	Books[b.ISBN] = b
	return b.ISBN, true
}

// GetBook returns a book if it exists
func GetBook(isbn string) ([]Book, bool) {
	// Check if ISBN exists in map
	if _, exists := Books[isbn]; exists {
		return []Book{Books[isbn]}, true
	}
	return make([]Book, 0), false
}
