package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/machinebox/graphql"
	"github.com/threagile/threagile/model"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hello, World!")
	client := graphql.NewClient("http://localhost:8080/graphql")
	path := "/Users/cesar/Project/empathy/secops/threat-model/.cdktg.out/models/"
	files := readFolder(path)
	for _, file := range files {
		var modelInput model.ModelInput = readYamlFileModel(path + file)
		// fmt.Println(modelInput)
		// client := graphql.NewClient("https://nameless-brook-390012.eu-central-1.aws.cloud.dgraph.io/graphql")
		UploadDataAssets(modelInput, client)
		println("## TECH ASSETS ##")
		UploadTechAssets(modelInput, client)
		println("## TRUST BOUNDARIES ##")
		UploadConnections(modelInput, client)
		UploadTrustBoundaries(modelInput, client)
	}
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

func readFolder(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var result []string = make([]string, 0)
	for _, f := range files {
		result = append(result, f.Name())
	}
	return result
}
