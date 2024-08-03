package main

import (
	f "fmt"
	_ "gopkg.in/yaml.v2"
)

func yamlParser(yamlp []byte) (result []byte) {
	f.Print("arighato")
}

func main() {

	var url string
	f.Printf("\nSubmit your URL here : ")
	f.Scanf("%s\n", &url)

}
