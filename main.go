package main

type ParsedData struct {
	name   string
	age    int
	skills []string
}

func main() {
	var pars
	data := `{
		"name": "Alice",
		"age": 30,
		"skills": ["Golang", "Docker", "Kubernetes"]
	  }
	  `

	parsedData

}
