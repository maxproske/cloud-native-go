// go test ./api -v
package api

import (
	"testing" // Used with the "go test" command

	"github.com/stretchr/testify/assert"
)

// TestBookToJASON unit tests BookToJSON()
// Follows format: func TestXxx(*testing.T)
func TestBookToJSON(t *testing.T) {
	// Define a simple book
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	// Issue call to ToJSON message and get response
	json := book.ToJSON()
	// Assert the response with some expected string value
	assert.Equal(t, // Testing
		`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`, // Expected
		string(json),                   // Actual
		"Book JSON marshalling wrong.") // Message in case of error
}

// TestBookFromJASON unit tests BookFromJSON()
func TestBookFromJSON(t *testing.T) {
	// Define byte array
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`) // Get byte array
	book := FromJSON(json)
	assert.Equal(t, // Testing
		Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}, // Expected
		book, // Actual
		"Book JSON unmarshalling wrong.") // Message in case of error
}
