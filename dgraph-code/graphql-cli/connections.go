package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
	"github.com/threagile/threagile/model"
)

func UploadConnections(model model.ModelInput, client *graphql.Client) {
	ctx := context.Background()
	techAssets := mapTechAssetToModelInput(model)
	// techAssetsId := make(map[string]string, 0)
	print("## TECH ASSETS ##")
	reqConnections := graphql.NewRequest(`
	mutation MyMutation($name: StringHashFilter,$communication_links: [CommunicationLinkRef]) {
		updateTechAsset(input: {filter: {name: $name}, set: {communication_links: $communication_links}}) {
		  techAsset {
			name
		  }
		}
	  }
		   
	`)
	for _, techAsset := range techAssets {
		var searchTerm struct {
			Op string `json:"eq"`
		}
		searchTerm.Op = techAsset.Name
		fmt.Println(searchTerm)
		reqConnections.Var("name", searchTerm)
		reqConnections.Var("communication_links", techAsset.Communication_links)
		reqConnections.Header.Set("Cache-Control", "no-cache")
		// fmt.Println(reqConnections)
		var respTechAsset interface{}
		if err := client.Run(ctx, reqConnections, &respTechAsset); err != nil {
			log.Fatal(err)
		}
		// techAssetsId[key] = respTechAsset.(map[string]interface{})["updateTechAsset"].(map[string]interface{})["techAsset"].([]interface{})[0].(map[string]interface{})["name"].(string)
		fmt.Println(respTechAsset)
	}
	// return techAssetsId
}
