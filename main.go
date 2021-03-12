package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const yamlFilePath = "./examples/chat-ws.yaml"
const templatePath = "./templates/documentation.html"

func main() {
	actions := getActionsFromFile(yamlFilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, actions)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
