package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
	"github.com/threagile/threagile/model"
)

func UploadTrustBoundaries(model model.ModelInput, client *graphql.Client) {
	ctx := context.Background()
	trustBoundaries := mapTrustBoundaryToModelInput(model)
	// trustBoundariesId := make(map[string]string, 0)
	println("## TRUST BOUNDARIES ##")
	reqTrustBoundaries := graphql.NewRequest(`
	mutation createTrustBoundary($input: [AddTrustBoundaryInput!]!) {
		addTrustBoundary(input: $input, upsert: true) {
			trustBoundary{
				name
			}
		}
	}
		   
	`)
	for _, trustBoundary := range trustBoundaries {
		reqTrustBoundaries.Var("input", trustBoundary)
		reqTrustBoundaries.Header.Set("Cache-Control", "no-cache")
		var respDataAsset interface{}
		bla, _ := json.Marshal(trustBoundary)
		fmt.Println(string(bla))
		if err := client.Run(ctx, reqTrustBoundaries, &respDataAsset); err != nil {
			log.Fatal(err)
		}
		// dataAssetsId[key] = respDataAsset.(map[string]interface{})["addDataAsset"].(map[string]interface{})["dataAsset"].([]interface{})[0].(map[string]interface{})["id"].(string)

	}
	// return trustBoundariesId
}

func mapTrustBoundaryToModelInput(model model.ModelInput) []TrustBoundaryRef {
	var trustBoundaryInput []TrustBoundaryRef

	for _, trustBoundary := range model.Trust_boundaries {
		nestedTB := make([]TrustBoundaryRef, 0)
		tecAssets := make([]TechAssetRef, 0)
		for _, nested := range trustBoundary.Trust_boundaries_nested {
			nestedTB = append(nestedTB, TrustBoundaryRef{
				Name: nested,
			})
		}
		for _, techAsset := range trustBoundary.Technical_assets_inside {
			tecAssets = append(tecAssets, TechAssetRef{
				Name: techAsset,
			})
		}
		trustBoundaryInput = append(trustBoundaryInput, TrustBoundaryRef{
			Name:                    trustBoundary.ID,
			Description:             trustBoundary.Description,
			Type:                    trustBoundary.Type,
			Nested_trust_boundaries: nestedTB,
			Tech_assets:             tecAssets,
		})
	}
	return trustBoundaryInput
}
