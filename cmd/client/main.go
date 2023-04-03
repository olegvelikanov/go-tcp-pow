package main

import (
	"github.com/olegvelikanov/go-tcp-pow/internal/app/client"
	"log"
)

func main() {
	quote, err := client.FetchQuote(3000)
	if err != nil {
		log.Printf("Error fetching quote: %s", err)
		return
	}
	log.Printf("Succesfully fetched the quote: %s", string(quote))
}
