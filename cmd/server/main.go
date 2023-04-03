package main

import (
	"flag"
	"github.com/olegvelikanov/go-tcp-pow/internal/app/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	log.Printf("loading config from: %s", configPath)
	config, err := server.LoadConfigFromFile(configPath)
	if err != nil {
		log.Fatalf("Error reading the config: %s", err)
	}

	s, err := server.StartServer(config)
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
