package main

import "encoding/json"

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

	decoder := json.NewDecoder(data)
	err := decoder.Decode(&parsedData)

}
