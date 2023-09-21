package main

import (
	"fmt"
	"os"

	"github.com/jaronnie/genius"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	fileBytes, err := os.ReadFile("./swagger.yaml")
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}

	err = yaml.Unmarshal(fileBytes, &data)
	if err != nil {
		panic(err)
	}

	g := genius.New(data)
	fmt.Println(g.GetAllKeys())
}
