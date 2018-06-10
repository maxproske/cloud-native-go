package api

import (
	"encoding/json"
)

// Book type with Name, Author, and ISBN
type Book struct {
	Title  string `json:"title"` // Marshall as lowercase
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
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
