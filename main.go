package main

type ParsedData struct {
	name   string
	age    int
	skills []string
}

func main() {
	var parsedData ParsedData
	data := `{
		"name": "Alice",
		"age": 30,
		"skills": ["Golang", "Docker", "Kubernetes"]
	  }
	  `

	decoder := js

}
