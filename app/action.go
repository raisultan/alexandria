package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type RawDocumentation struct {
	Info    Info
	Server  WSServer
	Actions map[string]RawAction
}

type Info struct {
	Title       string
	Description string
	Version     string
}

type WSServer struct {
	Url         string
	Description string
}

type RawAction struct {
	Description string
	Actor       struct {
		ObjectType string `yaml:"type"`
		Properties struct {
			IsSystem    bool `yaml:"is_system"`
			Description string
		}
	}
	Data map[string]interface{}

	ResponseToGroupRaw map[string]interface{} `yaml:"response_to_group"`
	ResponseToUserRaw  map[string]interface{} `yaml:"response_to_user"`
}

type Documentation struct {
	Info    Info
	Server  WSServer
	Actions []Action
}

type Action struct {
	Name        string
	Description string
	Actor       Actor
	DataSchema  string

	ResponseToGroupSchema string
	ResponseToUserSchema  string
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

func getDocumentationFromFile(path string) Documentation {
	rd := getRawDocumentationFromFile(path)
	info := Info{
		Title:       rd.Info.Title,
		Description: rd.Info.Description,
		Version:     rd.Info.Version,
	}
	server := WSServer{
		Url:         rd.Server.Url,
		Description: rd.Server.Description,
	}
	actions := []Action{}

	for name, action := range rd.Actions {
		actions = append(actions, newAction(name, action))
	}

	return Documentation{
		Info:    info,
		Server:  server,
		Actions: actions,
	}
}

func getRawDocumentationFromFile(path string) RawDocumentation {
	ymlContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rawDocumentation := RawDocumentation{}
	if err := yaml.Unmarshal([]byte(ymlContent), &rawDocumentation); err != nil {
		log.Fatal(err)
	}

	return rawDocumentation
}

func composeSchema(s map[string]interface{}) string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

// debug related
func (d Documentation) print() {
	fmt.Printf(
		"Info title: %s\nInfo description: %s\nInfo version: %s\nServer url: %s\nServer description: %s\n\n",
		d.Info.Title,
		d.Info.Description,
		d.Info.Version,
		d.Server.Url,
		d.Server.Description,
	)

	for _, action := range d.Actions {
		action.print()
	}
}

// debug related
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
