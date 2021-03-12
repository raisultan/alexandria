package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

func main() {
	ymlContent, err := ioutil.ReadFile("../examples/chat-ws.yaml")
	if err != nil {
		log.Fatal(err)
	}

	actions := make(map[string]Action)

	if err := yaml.Unmarshal([]byte(ymlContent), &actions); err != nil {
		log.Fatal(err)
	}

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
