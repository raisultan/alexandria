package main

import (
	"net/http"
	"time"
)

const yamlFilePath = "./examples/chat-ws.yaml"
const templatePath = "./templates/documentation.html"

func main() {
	hs, logger := setup()

	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)

		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	graceful(hs, logger, 5*time.Second)
}
