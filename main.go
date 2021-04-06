package main

import (
	"log"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	Diagram, err := yaml.ReadFile("diagram.yaml")
	if err != nil {
		log.Fatal(err)
	}
	state, err := Diagram.Count("diagram.from")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(state)
}
