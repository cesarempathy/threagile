package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
	"github.com/threagile/threagile/model"
)

func UploadTechAssets(model model.ModelInput, client *graphql.Client) map[string]string {
	ctx := context.Background()
	techAssets := mapTechAssetToModelInput(model)
	techAssetsId := make(map[string]string, 0)
	print("## TECH ASSETS ##")
	reqTechAssets := graphql.NewRequest(`
	mutation createTechAsset($input: [AddTechAssetInput!]!) {
		addTechAsset(input: $input, upsert: true) {
			techAsset{
				id
			}
		}
	}
`)

	for key, techAsset := range techAssets {
		reqTechAssets.Var("input", techAsset)
		reqTechAssets.Header.Set("Cache-Control", "no-cache")
		var respTechAsset interface{}
		// bla, _ := json.Marshal(techAsset)
		// fmt.Println(string(bla))
		if err := client.Run(ctx, reqTechAssets, &respTechAsset); err != nil {
			log.Fatal(err)
		}
		techAssetsId[key] = respTechAsset.(map[string]interface{})["addTechAsset"].(map[string]interface{})["techAsset"].([]interface{})[0].(map[string]interface{})["id"].(string)

	}
	return techAssetsId
}

func mapTechAssetToModelInput(model model.ModelInput) map[string]AddTechAssetInput {
	var techAssetInput map[string]AddTechAssetInput = make(map[string]AddTechAssetInput)
	// var connectionsMap = make(map[string][]string)
	for key, techAsset := range model.Technical_assets {

		var dataAssetsProcessed []DataAssetRef = make([]DataAssetRef, 0)
		for _, data := range techAsset.Data_assets_processed {
			dataAssetsProcessed = append(dataAssetsProcessed, DataAssetRef{
				Name: data,
			})
		}
		var dataAssetsStored []DataAssetRef = make([]DataAssetRef, 0)
		for _, data := range techAsset.Data_assets_stored {
			dataAssetsStored = append(dataAssetsStored, DataAssetRef{
				Name: data,
			})
		}

		techAssetInput[key] = AddTechAssetInput{
			Name:                       techAsset.ID,
			Description:                techAsset.Description,
			Type:                       techAsset.Type,
			Usage:                      techAsset.Usage,
			Used_as_client_by_human:    techAsset.Used_as_client_by_human,
			Out_of_scope:               techAsset.Out_of_scope,
			Technology:                 techAsset.Technology,
			Justification_out_of_scope: techAsset.Justification_out_of_scope,
			Size:                       techAsset.Size,
			Internet:                   techAsset.Internet,
			Machine:                    techAsset.Machine,
			Encryption:                 techAsset.Encryption,
			Owner:                      techAsset.Owner,
			Confidentiality:            techAsset.Confidentiality,
			Integrity:                  techAsset.Integrity,
			Availability:               techAsset.Availability,
			Justification_cia_rating:   techAsset.Justification_cia_rating,
			Tags:                       techAsset.Tags,
			Multi_tenant:               techAsset.Multi_tenant,
			Redundant:                  techAsset.Redundant,
			Custom_developed_parts:     techAsset.Custom_developed_parts,
			Data_formats_accepted:      techAsset.Data_formats_accepted,
			Data_assets_processed:      dataAssetsProcessed,
			Data_assets_stored:         dataAssetsStored,
			Trust_boundary:             nil,
			Runtime_environment:        nil,
		}

	}
	return techAssetInput
}
