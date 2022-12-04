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
	var modelInput model.ModelInput = readYamlFileModel("../../demo/example/threagile.yaml")
	// fmt.Println(modelInput)
	client := graphql.NewClient("https://nameless-brook-390012.eu-central-1.aws.cloud.dgraph.io/graphql")
	// getTechAsset(context, graphqlClient, "threagile/threagile")
	UploadDataAssets(modelInput, client)
	println("## TECH ASSETS ##")
	UploadTechAssets(modelInput, client)
	// UploadConnections(modelInput, client)
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
