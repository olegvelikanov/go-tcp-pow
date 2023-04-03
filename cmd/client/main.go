package main

import (
	"github.com/olegvelikanov/go-tcp-pow/internal/app/client"
	"log"
	"os"
)

const serverAddressEnvVariable = "SERVER_ADDR"

func main() {
	addr := getAddr("127.0.0.1:3000")
	quote, err := client.FetchQuote(addr)
	if err != nil {
		log.Printf("Error fetching quote: %s", err)
		return
	}
	log.Printf("Succesfully fetched the quote: [%s]", string(quote))
}

func getAddr(defaultValue string) string {
	value, ok := os.LookupEnv(serverAddressEnvVariable)
	if !ok {
		log.Printf("env variable [%s] not found. Default value will be used", serverAddressEnvVariable)
		return defaultValue
	}
	return value
}
