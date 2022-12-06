package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type TMDiagram struct {
	Diagram             []TechAsset     `json:"techAssets"`
	TrustBoundary       []TrustBoundary `json:"trustBoundary"`
	TrustBoundaryOrphan []TrustBoundary `json:"trustBoundaryOrphan"`
}

type TrustBoundary struct {
	Name   string         `json:"TrustBoundary.name"`
	Parent *TrustBoundary `json:"TrustBoundary.parent_trust_boundary,omitempty"`
}

type TechAsset struct {
	TechAssetName               string                       `json:"TechAsset.name"`
	TechAssetCommunicationLinks []TechAssetCommunicationLink `json:"TechAsset.communication_links,omitempty"`
	TrustBoundary               *TrustBoundary               `json:"TechAsset.trust_boundary,omitempty"`
}

type TechAssetCommunicationLink struct {
	CommunicationLinkProtocol    CommunicationLinkProtocol `json:"CommunicationLink.protocol"`
	CommunicationLinkTechAssetTo TechAsset                 `json:"CommunicationLink.tech_asset_to"`
}

type CommunicationLinkProtocol string

const (
	HTTP                CommunicationLinkProtocol = "http"
	HTTPS               CommunicationLinkProtocol = "https"
	NosqlAccessProtocol CommunicationLinkProtocol = "nosql-access-protocol"
)

func main() {
	readDiagram()

	// drawDiagram()
}

func readDiagram() {
	client := newClient()
	q := `query {
			trustBoundaryOrphan(func:has(TrustBoundary.name)) @filter(NOT has(TrustBoundary.parent_trust_boundary)){
				TrustBoundary.name
			}
			trustBoundary(func:has(TrustBoundary.parent_trust_boundary)){
				TrustBoundary.name
				TrustBoundary.parent_trust_boundary {
					TrustBoundary.name
				}
			}
			techAssets(func: has(TechAsset.name)) {
				TechAsset.name
				TechAsset.trust_boundary {
					TrustBoundary.name

				}
				TechAsset.communication_links {
					CommunicationLink.protocol
					CommunicationLink.tech_asset_to {
						TechAsset.name
					}
				}
			}
		}`
	txn := client.NewReadOnlyTxn()
	ctx := context.Background()
	resp, err := txn.Query(ctx, q)
	if err != nil {
		log.Fatal(err)
	}

	var tmDiagram TMDiagram

	json.Unmarshal(resp.GetJson(), &tmDiagram)

	DrawDiagramTM(&tmDiagram)
	// drawDiagram()
}

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func drawDiagram() {
	ruler, _ := textmeasure.NewRuler()
	_, graph, _ := d2lib.Compile(context.Background(), "x -> y", &d2lib.CompileOptions{
		Layout:  d2dagrelayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})

	// Create a shape with the ID, "meow"
	graph, _, _ = d2oracle.Create(graph, "x->y")
	// Style the shape green
	color := "green"
	graph, _ = d2oracle.Set(graph, "meow.style.fill", nil, &color)
	// Create a shape with the ID, "cat"
	graph, _, _ = d2oracle.Create(graph, "cat")

	// Move the shape "meow" inside the container "cat"
	graph, _ = d2oracle.Move(graph, "meow", "cat.meow")
	// Prints formatted D2 script
	fmt.Print(d2format.Format(graph.AST))
}

func DrawDiagramTM(diagram *TMDiagram) {
	ruler, _ := textmeasure.NewRuler()
	_, graph, _ := d2lib.Compile(context.Background(), "", &d2lib.CompileOptions{
		Layout:  d2dagrelayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	graph, trustBMap, _ := CreateTrustBoundary(graph, diagram)
	graph, _ = CreateTechAssets(graph, trustBMap, diagram)
	fmt.Print(d2format.Format(graph.AST))

}

func CreateTrustBoundary(graph *d2graph.Graph, diagram *TMDiagram) (*d2graph.Graph, map[string]string, error) {
	// Create a shape with the ID, "meow"
	trustBoundaryMap := make(map[string]string)
	for _, trustBoundary := range diagram.TrustBoundaryOrphan {
		graph, _, _ = d2oracle.Create(graph, trustBoundary.Name)
	}
	for _, trustBoundary := range diagram.TrustBoundary {
		graph, _, _ = d2oracle.Create(graph, trustBoundary.Name)
		graph, _ = d2oracle.Move(graph, trustBoundary.Name, trustBoundary.Parent.Name+"."+trustBoundary.Name)
		trustBoundaryMap[trustBoundary.Name] = trustBoundary.Parent.Name
	}
	// fmt.Print(d2format.Format(graph.AST))
	return graph, trustBoundaryMap, nil
}

func CreateTrustBoundaryId(id string, m map[string]string) string {
	if value, ok := m[id]; !ok {
		return id
	} else {
		return CreateTrustBoundaryId(value, m) + "." + id
	}
}

func CreateTechAssets(graph *d2graph.Graph, trustBMap map[string]string, diagram *TMDiagram) (*d2graph.Graph, error) {
	// Create a shape with the ID, "meow"
	mapIds := make(map[string]string)
	for _, techAsset := range diagram.Diagram {
		if techAsset.TrustBoundary != nil {
			id := fmt.Sprintf("%s.%s", CreateTrustBoundaryId(techAsset.TrustBoundary.Name, trustBMap), techAsset.TechAssetName)
			mapIds[techAsset.TechAssetName] = id
		}
	}
	for _, techAsset := range diagram.Diagram {
		// graph, _, _ = d2oracle.Create(graph, techAsset.TechAssetName)
		for _, communicationLink := range techAsset.TechAssetCommunicationLinks {
			link := fmt.Sprintf("%s->%s", techAsset.TechAssetName, communicationLink.CommunicationLinkTechAssetTo.TechAssetName)
			// fmt.Println(link)
			var linkKey string
			graph, linkKey, _ = d2oracle.Create(graph, link)
			color := "green"
			d2oracle.Set(graph, linkKey+".style.stroke", nil, &color)
			label := fmt.Sprintf("%s", communicationLink.CommunicationLinkProtocol)
			d2oracle.Set(graph, linkKey, nil, &label)
		}

	}
	for _, techAsset := range diagram.Diagram {
		if techAsset.TrustBoundary != nil {
			graph, _ = d2oracle.Move(graph, techAsset.TechAssetName, mapIds[techAsset.TechAssetName])
		}
	}
	// fmt.Print(d2format.Format(graph.AST))
	return graph, nil
}
