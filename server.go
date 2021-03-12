package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"
)

const templatePath = "./templates/documentation.html"
const defaultPort = "8080"

func setup() (*http.Server, *log.Logger) {
	addr := ":" + os.Getenv("ALEXANDRIA_PORT")
	if addr == ":" {
		addr = ":" + defaultPort
	}
	hs := &http.Server{Addr: addr, Handler: &server{}}

	return hs, log.New(os.Stdout, "", 0)
}

func graceful(hs *http.Server, logger *log.Logger, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Printf("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	yamlFilePath := os.Getenv("ALEXANDRIA_YAML")
	if yamlFilePath == "" {
		yamlFilePath = "./examples/chat-ws.yaml"
		// log.Fatal("ALEXANDRIA_YAML environment var is not provided")
	}

	actions := getActionsFromFile(yamlFilePath)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	tmpl.Execute(w, actions)
}
