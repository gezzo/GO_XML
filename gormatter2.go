package main

import (
	"fmt"
	"html/template"
	"os"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Greetings string            `yaml:"Greetings"`
	Pourpose  string            `yaml:"Pourpose"`
	Persons   map[string]Person `yaml:"Persons"`
}

type Person struct {
	Name       string `yaml:"Name"`
	FamilyName string `yaml:"FamilyName"`
	Gender     string `yaml:"Gender"`
	Age        int    `yaml:"Age"`
}

func main() {
	// Read the yaml file
	data, err := os.ReadFile("test.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var dataIngestion Data

	// unmarshal the data into the go struct
	err = yaml.Unmarshal([]byte(data), &dataIngestion)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(dataIngestion)

	xmlTemplate, err := template.ParseFiles("template.xml")
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
