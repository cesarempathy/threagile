package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
	"github.com/threagile/threagile/model"
)

func mapDataAssetToModelInput(model model.ModelInput) map[string]AddDataAssetInput {
	var dataAssetInput map[string]AddDataAssetInput = make(map[string]AddDataAssetInput)

	for _, dataAsset := range model.Data_assets {

		dataAssetInput[dataAsset.ID] = AddDataAssetInput{
			Name:                     dataAsset.ID,
			Description:              dataAsset.Description,
			Owner:                    dataAsset.Owner,
			Confidentiality:          dataAsset.Confidentiality,
			Integrity:                dataAsset.Integrity,
			Availability:             dataAsset.Availability,
			Justification_cia_rating: dataAsset.Justification_cia_rating,
			Tags:                     dataAsset.Tags,
			Usage:                    dataAsset.Usage,
			Origin:                   dataAsset.Origin,
			Quantity:                 dataAsset.Quantity,
		}

	}
	return dataAssetInput
}
func UploadDataAssets(model model.ModelInput, client *graphql.Client) map[string]string {
	ctx := context.Background()
	dataAssets := mapDataAssetToModelInput(model)
	dataAssetsId := make(map[string]string, 0)
	print("## DATA ASSETS ##")
	reqDataAssets := graphql.NewRequest(`
	mutation createDataAsset($input: [AddDataAssetInput!]!) {
		addDataAsset(input: $input, upsert: true) {
			dataAsset{
				id
			}
		}
	}
`)

	for key, dataAsset := range dataAssets {
		reqDataAssets.Var("input", dataAsset)
		reqDataAssets.Header.Set("Cache-Control", "no-cache")
		var respDataAsset interface{}
		if err := client.Run(ctx, reqDataAssets, &respDataAsset); err != nil {
			log.Fatal(err)
		}
		dataAssetsId[key] = respDataAsset.(map[string]interface{})["addDataAsset"].(map[string]interface{})["dataAsset"].([]interface{})[0].(map[string]interface{})["id"].(string)

	}
	return dataAssetsId
}
