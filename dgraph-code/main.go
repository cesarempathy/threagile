package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/threagile/threagile/model"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hello, World!")
	var modelInput model.ModelInput = readYamlFileModel("../demo/example/threagile.yaml")
	fmt.Println(modelInput)
}

func readYamlFileModel(path string) model.ModelInput {
	var modelInput model.ModelInput
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &modelInput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return modelInput
}
