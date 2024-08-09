package main

import (
	"AdventureQuest-cyoa/helper"
	_ "encoding/json"
	"flag"
	f "fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("story", "gophers.json", "a json file containing all the stories")
	flag.Parse()
	data, err := os.Open(*filename)
	if err != nil {
		f.Printf("Error opening file %s: ", *filename)
	}
	story, err := helper.ParseJson(data)
	if err != nil {
		f.Printf("Error Parsing filedata %s: ", *filename)
	}

	h := helper.NewHandler(story)

	log.Fatal(http.ListenAndServe(":3000", h))

}
