package main

import (
	"net/http"
	"time"
)

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
