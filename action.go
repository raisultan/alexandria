package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type RawAction struct {
	Description string
	Actor       struct {
		ObjectType string `yaml:"type"`
		Properties struct {
			IsSystem    bool
			Description string
		}
	}
	Data map[string]interface{}

	ResponseToGroupRaw map[string]interface{} `yaml:"response_to_group"`
	ResponseToUserRaw  map[string]interface{} `yaml:"response_to_user"`
}

type Action struct {
	Name        string
	Description string
	Actor       Actor
	DataSchema  string

	ResponseToGroupSchema string
	ResponseToUserSchema  string
}

func (a Action) print() {
	fmt.Printf(
		"Action: %s\nDescription: %s\nActor is system: %t\nActor description: %s\nData schema: %s\nResponse to user event schema: %s\nResponse to group event schema: %s\n\n",
		a.Name,
		a.Description,
		a.Actor.IsSystem,
		a.Actor.Description,
		a.DataSchema,
		a.ResponseToUserSchema,
		a.ResponseToGroupSchema,
	)
}

type Actor struct {
	IsSystem    bool
	Description string
}

func newAction(n string, r RawAction) Action {
	actor := Actor{
		IsSystem:    r.Actor.Properties.IsSystem,
		Description: r.Actor.Properties.Description,
	}
	dataSchema := composeSchema(r.Data)
	responseToUserSchema := composeSchema(r.ResponseToUserRaw)
	responseToGroupSchema := composeSchema(r.ResponseToGroupRaw)

	action := Action{
		Name:                  n,
		Description:           r.Description,
		Actor:                 actor,
		DataSchema:            dataSchema,
		ResponseToUserSchema:  responseToUserSchema,
		ResponseToGroupSchema: responseToGroupSchema,
	}

	return action
}

func getActionsFromFile(path string) []Action {
	rawActions := getRawActionsFromFile(path)
	actions := []Action{}

	for name, action := range rawActions {
		actions = append(actions, newAction(name, action))
	}

	return actions
}

func getRawActionsFromFile(path string) map[string]RawAction {
	ymlContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rawActions := make(map[string]RawAction)
	if err := yaml.Unmarshal([]byte(ymlContent), &rawActions); err != nil {
		log.Fatal(err)
	}

	return rawActions
}

func composeSchema(s map[string]interface{}) string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}
