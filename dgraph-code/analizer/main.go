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
	Test []struct {
		TechAssetName               string `json:"TechAsset.name"`
		TechAssetCommunicationLinks []struct {
			CommunicationLinkProtocol    string `json:"CommunicationLink.protocol"`
			CommunicationLinkTechAssetTo struct {
				TechAssetName string `json:"TechAsset.name"`
			} `json:"CommunicationLink.tech_asset_to"`
		} `json:"TechAsset.communication_links,omitempty"`
	} `json:"diagram"`
}

func main() {
	fmt.Println("Hello, World!")
	// readDiagram()

	drawDiagram()
}

func readDiagram() {
	client := newClient()
	q := `query {
		diagram(func: has(TechAsset.name)) {
			TechAsset.name
		TechAsset.communication_links {
		  CommunicationLink.protocol
				CommunicationLink.tech_asset_to{
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

	fmt.Println(string(resp.GetJson()))
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
	graph, _, _ = d2oracle.Create(graph, "x")
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

func CreateTrustBoundary(graph *d2graph.Graph, id string) (*d2graph.Graph, error) {
	// Create a shape with the ID, "meow"
	graph, _, _ = d2oracle.Create(graph, id)
	// Style the shape green
	color := "green"
	graph, _ = d2oracle.Set(graph, "meow.style.fill", nil, &color)
	// Create a shape with the ID, "cat"
	graph, _, _ = d2oracle.Create(graph, "cat")
	// Move the shape "meow" inside the container "cat"
	graph, _ = d2oracle.Move(graph, "meow", "cat.meow")
	// Prints formatted D2 script
	fmt.Print(d2format.Format(graph.AST))
	return graph, nil
}
