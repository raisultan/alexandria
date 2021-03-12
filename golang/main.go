package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	actions := getActionsFromFile("../examples/chat-ws.yaml")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./ws-documentation.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, actions)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
