package main

import (
	"github.com/olegvelikanov/word-of-wisdom/internal/app/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := server.StartServer(3000)
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
	defer s.Stop()
	waitForExit()
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
