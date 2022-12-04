// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type AddDataAssetInput struct {
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Usage                    string   `json:"usage,omitempty"`
	Tags                     []string `json:"tags,omitempty"`
	Origin                   string   `json:"origin,omitempty"`
	Owner                    string   `json:"owner,omitempty"`
	Quantity                 string   `json:"quantity,omitempty"`
	Confidentiality          string   `json:"confidentiality,omitempty"`
	Integrity                string   `json:"integrity,omitempty"`
	Availability             string   `json:"availability,omitempty"`
	Justification_cia_rating string   `json:"justification_cia_rating,omitempty"`
}

// GetName returns AddDataAssetInput.Name, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetName() string { return v.Name }

// GetDescription returns AddDataAssetInput.Description, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetDescription() string { return v.Description }

// GetUsage returns AddDataAssetInput.Usage, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetUsage() string { return v.Usage }

// GetTags returns AddDataAssetInput.Tags, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetTags() []string { return v.Tags }

// GetOrigin returns AddDataAssetInput.Origin, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetOrigin() string { return v.Origin }

// GetOwner returns AddDataAssetInput.Owner, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetOwner() string { return v.Owner }

// GetQuantity returns AddDataAssetInput.Quantity, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetQuantity() string { return v.Quantity }

// GetConfidentiality returns AddDataAssetInput.Confidentiality, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetConfidentiality() string { return v.Confidentiality }

// GetIntegrity returns AddDataAssetInput.Integrity, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetIntegrity() string { return v.Integrity }

// GetAvailability returns AddDataAssetInput.Availability, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetAvailability() string { return v.Availability }

// GetJustification_cia_rating returns AddDataAssetInput.Justification_cia_rating, and is useful for accessing the field via an interface.
func (v *AddDataAssetInput) GetJustification_cia_rating() string { return v.Justification_cia_rating }

type AddTechAssetInput struct {
	Name                       string                 `json:"name,omitempty"`
	Description                string                 `json:"description,omitempty"`
	Type                       string                 `json:"type,omitempty"`
	Usage                      string                 `json:"usage,omitempty"`
	Used_as_client_by_human    bool                   `json:"used_as_client_by_human,omitempty"`
	Out_of_scope               bool                   `json:"out_of_scope,omitempty"`
	Justification_out_of_scope string                 `json:"Justification_out_of_scope,omitempty"`
	Size                       string                 `json:"size,omitempty"`
	Technology                 string                 `json:"technology,omitempty"`
	Internet                   bool                   `json:"internet,omitempty"`
	Machine                    string                 `json:"machine,omitempty"`
	Encryption                 string                 `json:"encryption,omitempty"`
	Owner                      string                 `json:"owner,omitempty"`
	Confidentiality            string                 `json:"confidentiality,omitempty"`
	Integrity                  string                 `json:"integrity,omitempty"`
	Availability               string                 `json:"availability,omitempty"`
	Justification_cia_rating   string                 `json:"justification_cia_rating,omitempty"`
	Tags                       []string               `json:"tags,omitempty"`
	Multi_tenant               bool                   `json:"multi_tenant,omitempty"`
	Redundant                  bool                   `json:"redundant,omitempty"`
	Custom_developed_parts     bool                   `json:"custom_developed_parts,omitempty"`
	Data_formats_accepted      []string               `json:"data_formats_accepted,omitempty"`
	Data_assets_processed      []DataAssetRef         `json:"data_assets_processed,omitempty"`
	Data_assets_stored         []DataAssetRef         `json:"data_assets_stored,omitempty"`
	Communication_links        []CommunicationLinkRef `json:"communication_links,omitempty"`
	Trust_boundary             *TrustBoundaryRef       `json:"trust_boundary,omitempty"`
	Runtime_environment        *SharedRunTimeRef       `json:"runtime_environment,omitempty"`
}

// GetName returns AddTechAssetInput.Name, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetName() string { return v.Name }

// GetDescription returns AddTechAssetInput.Description, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetDescription() string { return v.Description }

// GetType returns AddTechAssetInput.Type, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetType() string { return v.Type }

// GetUsage returns AddTechAssetInput.Usage, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetUsage() string { return v.Usage }

// GetUsed_as_client_by_human returns AddTechAssetInput.Used_as_client_by_human, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetUsed_as_client_by_human() bool { return v.Used_as_client_by_human }

// GetOut_of_scope returns AddTechAssetInput.Out_of_scope, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetOut_of_scope() bool { return v.Out_of_scope }

// GetJustification_out_of_scope returns AddTechAssetInput.Justification_out_of_scope, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetJustification_out_of_scope() string {
	return v.Justification_out_of_scope
}

// GetSize returns AddTechAssetInput.Size, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetSize() string { return v.Size }

// GetTechnology returns AddTechAssetInput.Technology, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetTechnology() string { return v.Technology }

// GetInternet returns AddTechAssetInput.Internet, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetInternet() bool { return v.Internet }

// GetMachine returns AddTechAssetInput.Machine, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetMachine() string { return v.Machine }

// GetEncryption returns AddTechAssetInput.Encryption, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetEncryption() string { return v.Encryption }

// GetOwner returns AddTechAssetInput.Owner, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetOwner() string { return v.Owner }

// GetConfidentiality returns AddTechAssetInput.Confidentiality, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetConfidentiality() string { return v.Confidentiality }

// GetIntegrity returns AddTechAssetInput.Integrity, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetIntegrity() string { return v.Integrity }

// GetAvailability returns AddTechAssetInput.Availability, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetAvailability() string { return v.Availability }

// GetJustification_cia_rating returns AddTechAssetInput.Justification_cia_rating, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetJustification_cia_rating() string { return v.Justification_cia_rating }

// GetTags returns AddTechAssetInput.Tags, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetTags() []string { return v.Tags }

// GetMulti_tenant returns AddTechAssetInput.Multi_tenant, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetMulti_tenant() bool { return v.Multi_tenant }

// GetRedundant returns AddTechAssetInput.Redundant, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetRedundant() bool { return v.Redundant }

// GetCustom_developed_parts returns AddTechAssetInput.Custom_developed_parts, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetCustom_developed_parts() bool { return v.Custom_developed_parts }

// GetData_formats_accepted returns AddTechAssetInput.Data_formats_accepted, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetData_formats_accepted() []string { return v.Data_formats_accepted }

// GetData_assets_processed returns AddTechAssetInput.Data_assets_processed, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetData_assets_processed() []DataAssetRef { return v.Data_assets_processed }

// GetData_assets_stored returns AddTechAssetInput.Data_assets_stored, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetData_assets_stored() []DataAssetRef { return v.Data_assets_stored }

// GetCommunication_links returns AddTechAssetInput.Communication_links, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetCommunication_links() []CommunicationLinkRef {
	return v.Communication_links
}

// GetTrust_boundary returns AddTechAssetInput.Trust_boundary, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetTrust_boundary() *TrustBoundaryRef { return v.Trust_boundary }

// GetRuntime_environment returns AddTechAssetInput.Runtime_environment, and is useful for accessing the field via an interface.
func (v *AddTechAssetInput) GetRuntime_environment() *SharedRunTimeRef { return v.Runtime_environment }

type CommunicationLinkRef struct {
	Id                   string         `json:"id,omitempty"`
	Description          string         `json:"description,omitempty"`
	Protocol             string         `json:"protocol,omitempty"`
	Authentication       string         `json:"authentication,omitempty"`
	Authorization        string         `json:"authorization,omitempty"`
	Tags                 []string       `json:"tags,omitempty"`
	Vpn                  bool           `json:"vpn,omitempty"`
	Ip_filtered          bool           `json:"ip_filtered,omitempty"`
	Readonly             bool           `json:"readonly,omitempty"`
	Usage                string         `json:"usage,omitempty"`
	Data_assets_sent     []DataAssetRef `json:"data_assets_sent,omitempty"`
	Data_assets_received []DataAssetRef `json:"data_assets_received,omitempty"`
	Tech_asset_from      *TechAssetRef   `json:"tech_asset_from,omitempty"`
	Tech_asset_to        *TechAssetRef   `json:"tech_asset_to,omitempty"`
}

// GetId returns CommunicationLinkRef.Id, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetId() string { return v.Id }

// GetDescription returns CommunicationLinkRef.Description, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetDescription() string { return v.Description }

// GetProtocol returns CommunicationLinkRef.Protocol, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetProtocol() string { return v.Protocol }

// GetAuthentication returns CommunicationLinkRef.Authentication, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetAuthentication() string { return v.Authentication }

// GetAuthorization returns CommunicationLinkRef.Authorization, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetAuthorization() string { return v.Authorization }

// GetTags returns CommunicationLinkRef.Tags, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetTags() []string { return v.Tags }

// GetVpn returns CommunicationLinkRef.Vpn, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetVpn() bool { return v.Vpn }

// GetIp_filtered returns CommunicationLinkRef.Ip_filtered, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetIp_filtered() bool { return v.Ip_filtered }

// GetReadonly returns CommunicationLinkRef.Readonly, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetReadonly() bool { return v.Readonly }

// GetUsage returns CommunicationLinkRef.Usage, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetUsage() string { return v.Usage }

// GetData_assets_sent returns CommunicationLinkRef.Data_assets_sent, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetData_assets_sent() []DataAssetRef { return v.Data_assets_sent }

// GetData_assets_received returns CommunicationLinkRef.Data_assets_received, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetData_assets_received() []DataAssetRef {
	return v.Data_assets_received
}

// GetTech_asset_from returns CommunicationLinkRef.Tech_asset_from, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetTech_asset_from() *TechAssetRef { return v.Tech_asset_from }

// GetTech_asset_to returns CommunicationLinkRef.Tech_asset_to, and is useful for accessing the field via an interface.
func (v *CommunicationLinkRef) GetTech_asset_to() *TechAssetRef { return v.Tech_asset_to }

type DataAssetRef struct {
	Id                       string   `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Usage                    string   `json:"usage,omitempty"`
	Tags                     []string `json:"tags,omitempty"`
	Origin                   string   `json:"origin,omitempty"`
	Owner                    string   `json:"owner,omitempty"`
	Quantity                 string   `json:"quantity,omitempty"`
	Confidentiality          string   `json:"confidentiality,omitempty"`
	Integrity                string   `json:"integrity,omitempty"`
	Availability             string   `json:"availability,omitempty"`
	Justification_cia_rating string   `json:"justification_cia_rating,omitempty"`
}

// GetId returns DataAssetRef.Id, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetId() string { return v.Id }

// GetName returns DataAssetRef.Name, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetName() string { return v.Name }

// GetDescription returns DataAssetRef.Description, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetDescription() string { return v.Description }

// GetUsage returns DataAssetRef.Usage, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetUsage() string { return v.Usage }

// GetTags returns DataAssetRef.Tags, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetTags() []string { return v.Tags }

// GetOrigin returns DataAssetRef.Origin, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetOrigin() string { return v.Origin }

// GetOwner returns DataAssetRef.Owner, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetOwner() string { return v.Owner }

// GetQuantity returns DataAssetRef.Quantity, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetQuantity() string { return v.Quantity }

// GetConfidentiality returns DataAssetRef.Confidentiality, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetConfidentiality() string { return v.Confidentiality }

// GetIntegrity returns DataAssetRef.Integrity, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetIntegrity() string { return v.Integrity }

// GetAvailability returns DataAssetRef.Availability, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetAvailability() string { return v.Availability }

// GetJustification_cia_rating returns DataAssetRef.Justification_cia_rating, and is useful for accessing the field via an interface.
func (v *DataAssetRef) GetJustification_cia_rating() string { return v.Justification_cia_rating }

type SharedRunTimeRef struct {
	Id            string         `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Owner         string         `json:"owner,omitempty"`
	Justification string         `json:"justification,omitempty"`
	Tech_assets   []TechAssetRef `json:"tech_assets,omitempty"`
}

// GetId returns SharedRunTimeRef.Id, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetId() string { return v.Id }

// GetName returns SharedRunTimeRef.Name, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetName() string { return v.Name }

// GetDescription returns SharedRunTimeRef.Description, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetDescription() string { return v.Description }

// GetOwner returns SharedRunTimeRef.Owner, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetOwner() string { return v.Owner }

// GetJustification returns SharedRunTimeRef.Justification, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetJustification() string { return v.Justification }

// GetTech_assets returns SharedRunTimeRef.Tech_assets, and is useful for accessing the field via an interface.
func (v *SharedRunTimeRef) GetTech_assets() []TechAssetRef { return v.Tech_assets }

type TechAssetRef struct {
	Id                         string                 `json:"id,omitempty"`
	Name                       string                 `json:"name"`
	Description                string                 `json:"description"`
	Type                       string                 `json:"type"`
	Usage                      string                 `json:"usage"`
	Used_as_client_by_human    bool                   `json:"used_as_client_by_human"`
	Out_of_scope               bool                   `json:"out_of_scope"`
	Justification_out_of_scope string                 `json:"Justification_out_of_scope"`
	Size                       string                 `json:"size"`
	Technology                 string                 `json:"technology"`
	Internet                   bool                   `json:"internet"`
	Machine                    string                 `json:"machine"`
	Encryption                 string                 `json:"encryption"`
	Owner                      string                 `json:"owner"`
	Confidentiality            string                 `json:"confidentiality"`
	Integrity                  string                 `json:"integrity"`
	Availability               string                 `json:"availability"`
	Justification_cia_rating   string                 `json:"justification_cia_rating"`
	Tags                       []string               `json:"tags"`
	Multi_tenant               bool                   `json:"multi_tenant"`
	Redundant                  bool                   `json:"redundant"`
	Custom_developed_parts     bool                   `json:"custom_developed_parts"`
	Data_formats_accepted      []string               `json:"data_formats_accepted"`
	Data_assets_processed      []DataAssetRef         `json:"data_assets_processed"`
	Data_assets_stored         []DataAssetRef         `json:"data_assets_stored"`
	Communication_links        []CommunicationLinkRef `json:"communication_links"`
	Trust_boundary             *TrustBoundaryRef       `json:"trust_boundary,omitempty"`
	Runtime_environment        *SharedRunTimeRef       `json:"runtime_environment,omitempty"`
}

// GetId returns TechAssetRef.Id, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetId() string { return v.Id }

// GetName returns TechAssetRef.Name, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetName() string { return v.Name }

// GetDescription returns TechAssetRef.Description, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetDescription() string { return v.Description }

// GetType returns TechAssetRef.Type, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetType() string { return v.Type }

// GetUsage returns TechAssetRef.Usage, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetUsage() string { return v.Usage }

// GetUsed_as_client_by_human returns TechAssetRef.Used_as_client_by_human, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetUsed_as_client_by_human() bool { return v.Used_as_client_by_human }

// GetOut_of_scope returns TechAssetRef.Out_of_scope, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetOut_of_scope() bool { return v.Out_of_scope }

// GetJustification_out_of_scope returns TechAssetRef.Justification_out_of_scope, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetJustification_out_of_scope() string { return v.Justification_out_of_scope }

// GetSize returns TechAssetRef.Size, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetSize() string { return v.Size }

// GetTechnology returns TechAssetRef.Technology, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetTechnology() string { return v.Technology }

// GetInternet returns TechAssetRef.Internet, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetInternet() bool { return v.Internet }

// GetMachine returns TechAssetRef.Machine, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetMachine() string { return v.Machine }

// GetEncryption returns TechAssetRef.Encryption, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetEncryption() string { return v.Encryption }

// GetOwner returns TechAssetRef.Owner, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetOwner() string { return v.Owner }

// GetConfidentiality returns TechAssetRef.Confidentiality, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetConfidentiality() string { return v.Confidentiality }

// GetIntegrity returns TechAssetRef.Integrity, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetIntegrity() string { return v.Integrity }

// GetAvailability returns TechAssetRef.Availability, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetAvailability() string { return v.Availability }

// GetJustification_cia_rating returns TechAssetRef.Justification_cia_rating, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetJustification_cia_rating() string { return v.Justification_cia_rating }

// GetTags returns TechAssetRef.Tags, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetTags() []string { return v.Tags }

// GetMulti_tenant returns TechAssetRef.Multi_tenant, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetMulti_tenant() bool { return v.Multi_tenant }

// GetRedundant returns TechAssetRef.Redundant, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetRedundant() bool { return v.Redundant }

// GetCustom_developed_parts returns TechAssetRef.Custom_developed_parts, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetCustom_developed_parts() bool { return v.Custom_developed_parts }

// GetData_formats_accepted returns TechAssetRef.Data_formats_accepted, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetData_formats_accepted() []string { return v.Data_formats_accepted }

// GetData_assets_processed returns TechAssetRef.Data_assets_processed, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetData_assets_processed() []DataAssetRef { return v.Data_assets_processed }

// GetData_assets_stored returns TechAssetRef.Data_assets_stored, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetData_assets_stored() []DataAssetRef { return v.Data_assets_stored }

// GetCommunication_links returns TechAssetRef.Communication_links, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetCommunication_links() []CommunicationLinkRef { return v.Communication_links }

// GetTrust_boundary returns TechAssetRef.Trust_boundary, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetTrust_boundary() *TrustBoundaryRef { return v.Trust_boundary }

// GetRuntime_environment returns TechAssetRef.Runtime_environment, and is useful for accessing the field via an interface.
func (v *TechAssetRef) GetRuntime_environment() *SharedRunTimeRef { return v.Runtime_environment }

type TrustBoundaryRef struct {
	Id                      string             `json:"id,omitempty"`
	Name                    string             `json:"name,omitempty"`
	Description             string             `json:"description,omitempty"`
	Owner                   string             `json:"owner,omitempty"`
	Justification           string             `json:"justification,omitempty"`
	Tech_assets             []TechAssetRef     `json:"tech_assets,omitempty"`
	Nested_trust_boundaries []TrustBoundaryRef `json:"nested_trust_boundaries,omitempty"`
	Parent_trust_boundary   *TrustBoundaryRef   `json:"parent_trust_boundary,omitempty"`
}

// GetId returns TrustBoundaryRef.Id, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetId() string { return v.Id }

// GetName returns TrustBoundaryRef.Name, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetName() string { return v.Name }

// GetDescription returns TrustBoundaryRef.Description, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetDescription() string { return v.Description }

// GetOwner returns TrustBoundaryRef.Owner, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetOwner() string { return v.Owner }

// GetJustification returns TrustBoundaryRef.Justification, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetJustification() string { return v.Justification }

// GetTech_assets returns TrustBoundaryRef.Tech_assets, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetTech_assets() []TechAssetRef { return v.Tech_assets }

// GetNested_trust_boundaries returns TrustBoundaryRef.Nested_trust_boundaries, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetNested_trust_boundaries() []TrustBoundaryRef {
	return v.Nested_trust_boundaries
}

// GetParent_trust_boundary returns TrustBoundaryRef.Parent_trust_boundary, and is useful for accessing the field via an interface.
func (v *TrustBoundaryRef) GetParent_trust_boundary() *TrustBoundaryRef {
	return v.Parent_trust_boundary
}

// __createDataAssetInput is used internally by genqlient
type __createDataAssetInput struct {
	Input []AddDataAssetInput `json:"input,omitempty"`
}

// GetInput returns __createDataAssetInput.Input, and is useful for accessing the field via an interface.
func (v *__createDataAssetInput) GetInput() []AddDataAssetInput { return v.Input }

// __createTechAssetInput is used internally by genqlient
type __createTechAssetInput struct {
	Input []AddTechAssetInput `json:"input,omitempty"`
}

// GetInput returns __createTechAssetInput.Input, and is useful for accessing the field via an interface.
func (v *__createTechAssetInput) GetInput() []AddTechAssetInput { return v.Input }

// __getDataAssetInput is used internally by genqlient
type __getDataAssetInput struct {
	Id string `json:"id,omitempty"`
}

// GetId returns __getDataAssetInput.Id, and is useful for accessing the field via an interface.
func (v *__getDataAssetInput) GetId() string { return v.Id }

// __getTechAssetInput is used internally by genqlient
type __getTechAssetInput struct {
	Id string `json:"id,omitempty"`
}

// GetId returns __getTechAssetInput.Id, and is useful for accessing the field via an interface.
func (v *__getTechAssetInput) GetId() string { return v.Id }

// createDataAssetAddDataAssetAddDataAssetPayload includes the requested fields of the GraphQL type AddDataAssetPayload.
type createDataAssetAddDataAssetAddDataAssetPayload struct {
	NumUids int `json:"numUids,omitempty"`
}

// GetNumUids returns createDataAssetAddDataAssetAddDataAssetPayload.NumUids, and is useful for accessing the field via an interface.
func (v *createDataAssetAddDataAssetAddDataAssetPayload) GetNumUids() int { return v.NumUids }

// createDataAssetResponse is returned by createDataAsset on success.
type createDataAssetResponse struct {
	AddDataAsset createDataAssetAddDataAssetAddDataAssetPayload `json:"addDataAsset,omitempty"`
}

// GetAddDataAsset returns createDataAssetResponse.AddDataAsset, and is useful for accessing the field via an interface.
func (v *createDataAssetResponse) GetAddDataAsset() createDataAssetAddDataAssetAddDataAssetPayload {
	return v.AddDataAsset
}

// createTechAssetAddTechAssetAddTechAssetPayload includes the requested fields of the GraphQL type AddTechAssetPayload.
type createTechAssetAddTechAssetAddTechAssetPayload struct {
	NumUids int `json:"numUids,omitempty"`
}

// GetNumUids returns createTechAssetAddTechAssetAddTechAssetPayload.NumUids, and is useful for accessing the field via an interface.
func (v *createTechAssetAddTechAssetAddTechAssetPayload) GetNumUids() int { return v.NumUids }

// createTechAssetResponse is returned by createTechAsset on success.
type createTechAssetResponse struct {
	AddTechAsset createTechAssetAddTechAssetAddTechAssetPayload `json:"addTechAsset,omitempty"`
}

// GetAddTechAsset returns createTechAssetResponse.AddTechAsset, and is useful for accessing the field via an interface.
func (v *createTechAssetResponse) GetAddTechAsset() createTechAssetAddTechAssetAddTechAssetPayload {
	return v.AddTechAsset
}

// getDataAssetGetDataAsset includes the requested fields of the GraphQL type DataAsset.
type getDataAssetGetDataAsset struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetId returns getDataAssetGetDataAsset.Id, and is useful for accessing the field via an interface.
func (v *getDataAssetGetDataAsset) GetId() string { return v.Id }

// GetDescription returns getDataAssetGetDataAsset.Description, and is useful for accessing the field via an interface.
func (v *getDataAssetGetDataAsset) GetDescription() string { return v.Description }

// getDataAssetResponse is returned by getDataAsset on success.
type getDataAssetResponse struct {
	GetDataAsset getDataAssetGetDataAsset `json:"getDataAsset,omitempty"`
}

// GetGetDataAsset returns getDataAssetResponse.GetDataAsset, and is useful for accessing the field via an interface.
func (v *getDataAssetResponse) GetGetDataAsset() getDataAssetGetDataAsset { return v.GetDataAsset }

// getTechAssetGetTechAsset includes the requested fields of the GraphQL type TechAsset.
type getTechAssetGetTechAsset struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetId returns getTechAssetGetTechAsset.Id, and is useful for accessing the field via an interface.
func (v *getTechAssetGetTechAsset) GetId() string { return v.Id }

// GetDescription returns getTechAssetGetTechAsset.Description, and is useful for accessing the field via an interface.
func (v *getTechAssetGetTechAsset) GetDescription() string { return v.Description }

// getTechAssetResponse is returned by getTechAsset on success.
type getTechAssetResponse struct {
	GetTechAsset getTechAssetGetTechAsset `json:"getTechAsset,omitempty"`
}

// GetGetTechAsset returns getTechAssetResponse.GetTechAsset, and is useful for accessing the field via an interface.
func (v *getTechAssetResponse) GetGetTechAsset() getTechAssetGetTechAsset { return v.GetTechAsset }

func createDataAsset(
	ctx context.Context,
	client graphql.Client,
	input []AddDataAssetInput,
) (*createDataAssetResponse, error) {
	req := &graphql.Request{
		OpName: "createDataAsset",
		Query: `
mutation createDataAsset ($input: [AddDataAssetInput!]!) {
	addDataAsset(input: $input) {
		numUids
	}
}
`,
		Variables: &__createDataAssetInput{
			Input: input,
		},
	}
	var err error

	var data createDataAssetResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func createTechAsset(
	ctx context.Context,
	client graphql.Client,
	input []AddTechAssetInput,
) (*createTechAssetResponse, error) {
	req := &graphql.Request{
		OpName: "createTechAsset",
		Query: `
mutation createTechAsset ($input: [AddTechAssetInput!]!) {
	addTechAsset(input: $input) {
		numUids
	}
}
`,
		Variables: &__createTechAssetInput{
			Input: input,
		},
	}
	var err error

	var data createTechAssetResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func getDataAsset(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*getDataAssetResponse, error) {
	req := &graphql.Request{
		OpName: "getDataAsset",
		Query: `
query getDataAsset ($id: ID!) {
	getDataAsset(id: $id) {
		id
		description
	}
}
`,
		Variables: &__getDataAssetInput{
			Id: id,
		},
	}
	var err error

	var data getDataAssetResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func getTechAsset(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*getTechAssetResponse, error) {
	req := &graphql.Request{
		OpName: "getTechAsset",
		Query: `
query getTechAsset ($id: ID!) {
	getTechAsset(id: $id) {
		id
		description
	}
}
`,
		Variables: &__getTechAssetInput{
			Id: id,
		},
	}
	var err error

	var data getTechAssetResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
