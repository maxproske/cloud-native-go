package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/maxproske/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc) // Responds to http://localhost:8080/api/books
	http.HandleFunc("/api/books/", api.BookHandleFunc) // Retreive individual books, update, and delete them

	http.ListenAndServe(port(), nil)
}

func port() string {
	// $ export PORT=3020
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world!")
}

// ResponseWriter responds content to client. Request respond to HTTP request receieved
func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0] // Extract first message parameter
	fmt.Fprintf(w, message)
}
