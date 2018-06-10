// go test ./api -v
package api

import (
	"testing" // Used with the "go test" command

	"github.com/stretchr/testify/assert"
)

// TestBookToJASON unit tests BookToJSON()
// Follows format: func TestXxx(*testing.T)
func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"} // Define a simple book
	json := book.ToJSON()                                                              // Issue call to ToJSON message and get response
	assert.Equal(
		t,
		`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0000000000"}`, // Expected
		string(json), // Actual
		"Book JSON marshalling wrong.",
	) // Assert the response with some expected string value
}

// TestBookFromJASON unit tests BookFromJSON()
func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0000000000"}`) // Define byte array
	book := FromJSON(json)
	assert.Equal(
		t,
		Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"}, // Expected
		book, // Actual
		"Book JSON unmarshalling wrong.",
	)
}
