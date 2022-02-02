package main

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name      string
	Capital   string
	Continent string
}

func main() {
	var country []Country

	Data := []byte(`
	[
		{"Name": "Japan", "Capital": "Tokyo", "Continent": "Asia"},
		{"Name": "Japan1", "Capital": "Tokyo1", "Continent": "Asia1"},
		{"Name": "Japan2", "Capital": "Tokyo2", "Continent": "Asia2"},
		{"Name": "Japan2", "Capital": "Tokyo2", "Continent": "Asia2"},
	]
	`)
}

// Golang program to illustrate the
// concept of parsing json to struct package
main import (
	"encoding/json"
	"fmt"
)

// declaring a struct

type Country struct {
	// defining struct variables
	Name string
	Capital string
	Continent string
}

// main function
func main() {
	// defining a struct instance
	var country1 Country
	// data in JSON format which
	// is to be decoded
	Data := []byte(`
		{"Name": "Japan", "Capital": "Tokyo", "Continent": "Asia"}
	`)
		// decoding country1 struct
		// from json format
		err := json.Unmarshal(Data, &country1)
		if err != nil {
			// if error is not nil
			// print error
			fmt.Println(err) }
			// printing details of
			// decoded data
			fmt.Println("Struct is:", country1)
			fmt.Printf("%s's capital is %s and it is in %s.\n", country1.Name, country1.Capital, country1.Continent) }
:wrapper