package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json: "-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncode.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "Learncode.in", "bcd123", []string{"Fullstack", "js"}},
		{"MERN Bootcamp", 299, "Learncode.in", "sid123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "sid", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
