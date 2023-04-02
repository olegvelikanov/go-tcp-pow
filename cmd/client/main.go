package main

import "github.com/olegvelikanov/word-of-wisdom/internal/app/client"

func main() {
	client.FetchQuote(3000)
}
