package main

import (
	"encoding/json"
	"fmt"
)

type ParsedData struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
}

func main() {
	var parsedData ParsedData
	data := `{
		"name": "Alice",
		"age": 30,
		"skills": ["Golang", "Docker", "Kubernetes"]
	  }
	  `

	json.Unmarshal([]byte(data), parsedData)

	fmt.Println("Name: ", parsedData.Name)
	fmt.Println("Age: ", parsedData.Age)

}
