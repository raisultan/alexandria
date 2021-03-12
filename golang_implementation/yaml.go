package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Action struct {
	Description string
	Actor       struct {
		ObjectType string `yaml:"type"`
		Properties struct {
			IsSystem    bool
			Description string
		}
	}
	Data            map[string]interface{}
	ResponseToGroup map[string]interface{} `yaml:"response_to_group"`
	ResponseToUser  map[string]interface{} `yaml:"response_to_user"`
}

func (a Action) print() {
	fmt.Printf(
		"Description: %s\nActor is system: %t\nActor description: %s\nResponse to user event name: %s\nResponse to group event name: %s\n\n",
		a.Description,
		a.Actor.Properties.IsSystem,
		a.Actor.Properties.Description,
		a.ResponseToUser["event"],
		a.ResponseToGroup["event"],
	)
}

func main() {
	ymlContent, err := ioutil.ReadFile("../examples/chat-ws.yaml")
	if err != nil {
		log.Fatal(err)
	}

	actions := make(map[string]Action)

	if err := yaml.Unmarshal([]byte(ymlContent), &actions); err != nil {
		log.Fatal(err)
	}

	for key, value := range actions {
		fmt.Println("Action: ", key)
		value.print()
	}
}
