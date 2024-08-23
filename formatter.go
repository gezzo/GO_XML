package main

import (
	"fmt"
	"html/template"
	"os"

	"gopkg.in/yaml.v3"
)

type Dataingestion struct {
	Element1 string `yaml:"Element1"`
	Element2 string `yaml:"Element2"`
	Element3 string `yaml:"Element3"`
	Element4 string `yaml:"Element4"`
	Element5 string `yaml:"Element5"`
}

func main() {
	// Read the yaml file
	data, err := os.ReadFile("start.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var dataIngestion Dataingestion

	// unmarshal the data into the go struct
	err = yaml.Unmarshal(data, &dataIngestion)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(dataIngestion)

	xmlTemplate, err := template.ParseFiles("template_example.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	newXml, _ := os.Create("generatedXml.xml")

	err = xmlTemplate.Execute(newXml, dataIngestion)
	if err != nil {
		fmt.Println(err)
		return
	}

}
