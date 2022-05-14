package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	quote, err := GetQuote()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Next quote is %s", quote.Content)
}

func GetQuote() (Quote, error) {
	// Get the quote
	response, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return Quote{}, err
	}

	// Decode the response
	var quote Quote
	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return Quote{}, err
	}

	//return the quote
	return quote, nil
}
