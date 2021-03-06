package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title:"Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title:"Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title:"Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqualine Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data1, err1 := json.MarshalIndent(movies, "", "  ")
	if err1 != nil {
		log.Fatalf("JSON Marshaling failed: %s", err1)
	}
	fmt.Printf("%s\n", data1)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
