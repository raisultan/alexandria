package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Message struct {
	Environments map[string]models `yaml:"Environments"`
}
type models map[string][]Model
type Model struct {
	AppType     string `yaml:"app-type"`
	ServiceType string `yaml:"service-type"`
}

func main() {
	data := []byte(`
Environments:
 sys1:
    models:
    - app-type: app1
      service-type: fds
    - app-type: app2
      service-type: era
 sys2:
    models:
    - app-type: app1
      service-type: fds
    - app-type: app2
      service-type: era
`)
	y := Message{}

	err := yaml.Unmarshal([]byte(data), &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%+v\n", y)

}
