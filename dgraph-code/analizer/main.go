package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
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
	TechAssetTechnology         string                       `json:"TechAsset.technology,omitempty"`
	TechAssetCommunicationLinks []TechAssetCommunicationLink `json:"TechAsset.communication_links,omitempty"`
	TrustBoundary               *TrustBoundary               `json:"TechAsset.trust_boundary,omitempty"`
}

type TechAssetCommunicationLink struct {
	CommunicationLinkProtocol    CommunicationLinkProtocol `json:"CommunicationLink.protocol"`
	CommunicationLinkTechAssetTo TechAsset                 `json:"CommunicationLink.tech_asset_to"`
}

type CommunicationLinkProtocol string

func (diagram *TMDiagram) GetTechAssets() []TechAsset {
	nodesMap := make(map[string]TechAsset)
	for _, techAsset := range diagram.Diagram {
		nodesMap[techAsset.TechAssetName] = techAsset
		for _, techAssetCommunicationLink := range techAsset.TechAssetCommunicationLinks {
			if _, ok := nodesMap[techAssetCommunicationLink.CommunicationLinkTechAssetTo.TechAssetName]; !ok {
				nodesMap[techAssetCommunicationLink.CommunicationLinkTechAssetTo.TechAssetName] = techAssetCommunicationLink.CommunicationLinkTechAssetTo
				for _, techAssetCommunicationLink := range techAssetCommunicationLink.CommunicationLinkTechAssetTo.TechAssetCommunicationLinks {
					if _, ok := nodesMap[techAssetCommunicationLink.CommunicationLinkTechAssetTo.TechAssetName]; !ok {
						nodesMap[techAssetCommunicationLink.CommunicationLinkTechAssetTo.TechAssetName] = techAssetCommunicationLink.CommunicationLinkTechAssetTo
					}
				}
			}
		}
	}
	var nodes []TechAsset = make([]TechAsset, 0)
	for _, node := range nodesMap {
		nodes = append(nodes, node)
	}
	return nodes
}

const (
	HTTP                CommunicationLinkProtocol = "http"
	HTTPS               CommunicationLinkProtocol = "https"
	NosqlAccessProtocol CommunicationLinkProtocol = "nosql-access-protocol"
)

func main() {
	webServer()
	// readDiagram("search-service")
	// readFullDiagram()
	// drawDiagram()
}

func readDiagram(id string) ([]byte, error) {
	client := newClient()
	q := `query Diagram($id: string) {
			trustBoundaryOrphan(func:has(TrustBoundary.name)) @filter(NOT has(TrustBoundary.parent_trust_boundary)){
				TrustBoundary.name
			}
			trustBoundary(func:has(TrustBoundary.parent_trust_boundary)){
				TrustBoundary.name
				TrustBoundary.parent_trust_boundary {
					TrustBoundary.name
				}
			}
			techAssets(func: eq(TechAsset.name, $id)) {
				TechAsset.name
				TechAsset.technology
				TechAsset.trust_boundary {
					TrustBoundary.name

				}
				TechAsset.communication_links {
					CommunicationLink.protocol
					CommunicationLink.tech_asset_to {
						TechAsset.name 
						TechAsset.trust_boundary {
							TrustBoundary.name
						}
						TechAsset.communication_links {
							CommunicationLink.protocol     
							CommunicationLink.tech_asset_to {
								TechAsset.name 
								TechAsset.trust_boundary {
									TrustBoundary.name
								}
							}
						}
					}
				}
			}
		}`
	txn := client.NewReadOnlyTxn()
	ctx := context.Background()
	vars := map[string]string{"$id": id}
	resp, err := txn.QueryWithVars(ctx, q, vars)
	if err != nil {
		log.Fatal(err)
	}

	var tmDiagram TMDiagram

	json.Unmarshal(resp.GetJson(), &tmDiagram)

	return DrawDiagramTM(&tmDiagram)
	// drawDiagram()
}

func GetTechAssets() ([]string, error) {
	client := newClient()
	q := `query Diagram {
			techAssets(func: has(TechAsset.name)) {
				TechAsset.name
			}
		}`
	txn := client.NewReadOnlyTxn()
	ctx := context.Background()
	resp, err := txn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var tmDiagram TMDiagram

	json.Unmarshal(resp.GetJson(), &tmDiagram)

	var techAssets []string = make([]string, 0)
	for _, techAsset := range tmDiagram.Diagram {
		techAssets = append(techAssets, techAsset.TechAssetName)
	}
	return techAssets, nil
}

func readFullDiagram() ([]byte, error) {
	client := newClient()
	q := `query Diagram {
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
				TechAsset.technology
				TechAsset.trust_boundary {
					TrustBoundary.name

				}
				TechAsset.communication_links {
					CommunicationLink.protocol
					CommunicationLink.tech_asset_to {
						TechAsset.name 
						TechAsset.trust_boundary {
							TrustBoundary.name
							
						}
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

	return DrawDiagramTM(&tmDiagram)
	// drawDiagram()
}

func webServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("id") {
			tempate := `<html>
			<head>
				<title>Templated Application</title>
			</head>
			<body>
				<h1>Templated </h1>
				Id: <b>{{.Id}}<p />
				Name: <b> {{.Name}}</b>
			</body>
		</html>`
			fmt.Println("no id")
			return
		}
		id := r.URL.Query().Get("id")

		data, err := readDiagram(id)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		// w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		// w.Header().Set("Pragma", "no-cache")
		// w.Header().Set("Expires", "0")
		w.Write(data)
	})
	http.HandleFunc("/full", func(w http.ResponseWriter, r *http.Request) {

		data, err := readFullDiagram()
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		// w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		// w.Header().Set("Pragma", "no-cache")
		// w.Header().Set("Expires", "0")
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
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

func DrawDiagramTM(diagram *TMDiagram) ([]byte, error) {
	ruler, _ := textmeasure.NewRuler()
	_, graph, _ := d2lib.Compile(context.Background(), "", &d2lib.CompileOptions{
		Layout: d2dagrelayout.Layout,
		Ruler:  ruler,

		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	graph, trustBMap, _ := CreateTrustBoundary(graph, diagram)
	graph, _ = CreateTechAssets(graph, trustBMap, diagram)
	fmt.Print(d2format.Format(graph.AST))
	g, _, _ := d2lib.Compile(context.Background(), d2format.Format(graph.AST), &d2lib.CompileOptions{
		Layout:  d2elklayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	return d2svg.Render(g)

}

func CreateTrustBoundary(graph *d2graph.Graph, diagram *TMDiagram) (*d2graph.Graph, map[string]string, error) {
	// Create a shape with the ID, "meow"
	trustBoundaryMap := make(map[string]string)
	for _, trustBoundary := range diagram.TrustBoundaryOrphan {
		graph, _, _ = d2oracle.Create(graph, trustBoundary.Name)
	}
	for _, trustBoundary := range diagram.TrustBoundary {

		graph, _, _ = d2oracle.Create(graph, trustBoundary.Name)
		trustBoundaryMap[trustBoundary.Name] = trustBoundary.Parent.Name
		id := CreateTrustBoundaryId(trustBoundary.Name, trustBoundaryMap)
		graph, _ = d2oracle.Move(graph, trustBoundary.Name, id)

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
	for _, techAsset := range diagram.GetTechAssets() {
		if techAsset.TrustBoundary != nil {
			id := fmt.Sprintf("%s.%s", CreateTrustBoundaryId(techAsset.TrustBoundary.Name, trustBMap), techAsset.TechAssetName)
			mapIds[techAsset.TechAssetName] = id
		}

	}
	fmt.Println(mapIds)
	for _, techAsset := range diagram.GetTechAssets() {
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
	var err error
	for _, techAsset := range diagram.GetTechAssets() {
		if ok := ChooseIcon(techAsset); ok != "" {
			graph, err = SetIcon(graph, techAsset.TechAssetName, ok)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}

		fmt.Println(techAsset.TechAssetName, mapIds[techAsset.TechAssetName])
		if mapIds[techAsset.TechAssetName] != "" {
			graph, err = d2oracle.Move(graph, techAsset.TechAssetName, mapIds[techAsset.TechAssetName])
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	// fmt.Print(d2format.Format(graph.AST))
	return graph, nil
}

func SetIcon(graph *d2graph.Graph, id string, icon string) (*d2graph.Graph, error) {
	// Create a shape with the ID, "meow"
	// shape := "image"
	// var err error
	// graph, err = d2oracle.Set(graph, id+".icon", nil, &icon)
	// if err != nil {
	// 	return nil, err
	// }
	return d2oracle.Set(graph, id+".icon", nil, &icon)
}

func ChooseIcon(techAsset TechAsset) string {

	if isMongo(techAsset.TechAssetName, techAsset.TechAssetTechnology) {
		return "https://img.icons8.com/color/512/mongodb.png"
	}
	if isElastic(techAsset.TechAssetName, techAsset.TechAssetTechnology) {
		return "https://img.icons8.com/color/512/elasticsearch.png"
	}
	if isS3(techAsset.TechAssetName, techAsset.TechAssetTechnology) {
		return "https://img.icons8.com/color/512/amazon-s3.png"
	}
	if isRepo(techAsset.TechAssetName, techAsset.TechAssetTechnology) {
		return "https://img.icons8.com/color/512/github.png"
	}
	if isKinesis(techAsset.TechAssetName, techAsset.TechAssetTechnology) {
		return "https://www.pngrepo.com/png/353453/180/aws-kinesis.png"
	}

	return ""
}

func isMongo(techAssetName, tech string) bool {
	return strings.Contains(techAssetName, "mongo") && tech == "database"
}

func isElastic(techAssetName, tech string) bool {
	return strings.Contains(techAssetName, "elasticsearch") && tech == "search-index"
}

func isS3(techAssetName, tech string) bool {
	return strings.Contains(techAssetName, "s3") || tech == "block-storage"
}

func isKinesis(techAssetName, tech string) bool {
	return strings.Contains(techAssetName, "kinesis") || tech == "message-queue"
}

func isRepo(techAssetName, tech string) bool {
	return tech == "sourcecode-repository"
}
