package main

import "encoding/json"

type ParsedData struct {
	Name   string "js"
	Age    int
	Skills []string
}

func main() {
	var parsedData ParsedData
	data := `{
		"name": "Alice",
		"age": 30,
		"skills": ["Golang", "Docker", "Kubernetes"]
	  }
	  `

	decoder := json.NewDecoder(data)
	err := decoder.Decode(&parsedData)

}
