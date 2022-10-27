
package sdwan

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	sdwanAPI "github.com/oshafran/pied-piper-openapi-client-go"
)

var token string;
var (
	_ resource.Resource                = &templatePolicyDefinitionSrstphoneprofileIdResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyDefinitionSrstphoneprofileIdResource{}
	_ resource.ResourceWithImportState = &templatePolicyDefinitionSrstphoneprofileIdResource{}
)

func NewTemplatePolicyDefinitionSrstphoneprofileIdResource() resource.Resource {
	return &templatePolicyDefinitionSrstphoneprofileIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyDefinitionSrstphoneprofileIdResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyDefinitionSrstphoneprofileIdResourceModel struct{
  MasterTemplatesAffected              []String           `tfsdk:"masterTemplatesAffected"`
  Name              types.String           `tfsdk:"name"`
  Type              types.String           `tfsdk:"type"`
  Description              types.String           `tfsdk:"description"`
  DefaultAction              defaultActionResource           `tfsdk:"defaultAction"`
  Sequences              []SequencesResource           `tfsdk:"sequences"`
  Definition              definitionResource           `tfsdk:"definition"`
}

type defaultActionResource struct{
  Type              types.String           `tfsdk:"type"`
}


type sequencesResource struct{
  SequenceId              types.Int64           `tfsdk:"sequenceId"`
  SequenceName              types.String           `tfsdk:"sequenceName"`
  BaseAction              types.String           `tfsdk:"baseAction"`
  SequenceType              types.String           `tfsdk:"sequenceType"`
  SequenceIpType              types.String           `tfsdk:"sequenceIpType"`
  Match              matchResource           `tfsdk:"match"`
  Actions              []ActionsResource           `tfsdk:"actions"`
}

type matchResource struct{
  Entries              []EntriesResource           `tfsdk:"entries"`
}

type entriesResource struct{
  Field              types.String           `tfsdk:"field"`
  Value              types.String           `tfsdk:"value"`
}



type actionsResource struct{
  Type              types.String           `tfsdk:"type"`
  Parameter              parameterResource           `tfsdk:"parameter"`
}

type parameterResource struct{
  Ref              types.String           `tfsdk:"ref"`
}




type definitionResource struct{
  DefaultAction              defaultActionResource           `tfsdk:"defaultAction"`
  Sequences              []SequencesResource           `tfsdk:"sequences"`
  Entries              []EntriesResource           `tfsdk:"entries"`
  SignatureSet              types.String           `tfsdk:"signatureSet"`
  InspectionMode              types.String           `tfsdk:"inspectionMode"`
  SignatureWhiteList              signatureWhiteListResource           `tfsdk:"signatureWhiteList"`
  LogLevel              types.String           `tfsdk:"logLevel"`
  Logging              []String           `tfsdk:"logging"`
  TargetVpns              []TargetVpnsResource           `tfsdk:"targetVpns"`
  WebCategoriesAction              types.String           `tfsdk:"webCategoriesAction"`
  WebCategories              []String           `tfsdk:"webCategories"`
  WebReputation              types.String           `tfsdk:"webReputation"`
  UrlWhiteList              urlWhiteListResource           `tfsdk:"urlWhiteList"`
  UrlBlackList              urlBlackListResource           `tfsdk:"urlBlackList"`
  BlockPageAction              types.String           `tfsdk:"blockPageAction"`
  BlockPageContents              types.String           `tfsdk:"blockPageContents"`
  EnableAlerts              types.Bool           `tfsdk:"enableAlerts"`
  Alerts              []String           `tfsdk:"alerts"`
  MatchAllVpn              types.Bool           `tfsdk:"matchAllVpn"`
  FileReputationCloudServer              types.String           `tfsdk:"fileReputationCloudServer"`
  FileReputationEstServer              types.String           `tfsdk:"fileReputationEstServer"`
  FileReputationAlert              types.String           `tfsdk:"fileReputationAlert"`
  FileAnalysisCloudServer              types.String           `tfsdk:"fileAnalysisCloudServer"`
  FileAnalysisFileTypes              []String           `tfsdk:"fileAnalysisFileTypes"`
  FileAnalysisAlert              types.String           `tfsdk:"fileAnalysisAlert"`
  FileAnalysisEnabled              types.Bool           `tfsdk:"fileAnalysisEnabled"`
  LocalDomainBypassList              localDomainBypassListResource           `tfsdk:"localDomainBypassList"`
  DnsCrypt              types.Bool           `tfsdk:"dnsCrypt"`
  UmbrellaData              umbrellaDataResource           `tfsdk:"umbrellaData"`
}

type defaultActionResource struct{
  Type              types.String           `tfsdk:"type"`
}


type sequencesResource struct{
  SequenceId              types.Int64           `tfsdk:"sequenceId"`
  SequenceName              types.String           `tfsdk:"sequenceName"`
  BaseAction              types.String           `tfsdk:"baseAction"`
  SequenceType              types.String           `tfsdk:"sequenceType"`
  Match              matchResource           `tfsdk:"match"`
  Actions              []String           `tfsdk:"actions"`
}

type matchResource struct{
  Entries              []EntriesResource           `tfsdk:"entries"`
}

type entriesResource struct{
  Field              types.String           `tfsdk:"field"`
  Ref              types.String           `tfsdk:"ref"`
}




type entriesResource struct{
  SourceZone              types.String           `tfsdk:"sourceZone"`
  DestinationZone              types.String           `tfsdk:"destinationZone"`
}


type signatureWhiteListResource struct{
  Ref              types.String           `tfsdk:"ref"`
}


type targetVpnsResource struct{
  Vpns              []String           `tfsdk:"vpns"`
  UmbrellaDefault              types.Bool           `tfsdk:"umbrellaDefault"`
  DnsServerIp              types.String           `tfsdk:"dnsServerIP"`
  LocalDomainBypassEnabled              types.Bool           `tfsdk:"localDomainBypassEnabled"`
}


type urlWhiteListResource struct{
  Ref              types.String           `tfsdk:"ref"`
}


type urlBlackListResource struct{
  Ref              types.String           `tfsdk:"ref"`
}


type localDomainBypassListResource struct{
  Ref              types.String           `tfsdk:"ref"`
}


type umbrellaDataResource struct{
  Ref              types.String           `tfsdk:"ref"`
}



func (r *templatePolicyDefinitionSrstphoneprofileIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyDefinitionSrstphoneprofileIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_definition_srstphoneprofile_id"
}

func (d *templatePolicyDefinitionSrstphoneprofileIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
  masterTemplatesAffected              []String           `tfsdk:"masterTemplatesAffected"`
      "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "defaultAction": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "sequences": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "sequenceId": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "sequenceName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "baseAction": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sequenceType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sequenceIpType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "match": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "field": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "value": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},
,
      "actions": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "parameter": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
}),},
}),},
      "definition": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "defaultAction": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "sequences": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "sequenceId": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "sequenceName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "baseAction": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sequenceType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "match": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "field": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},
,
  actions              []String           `tfsdk:"actions"`
}),},
      "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "sourceZone": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "destinationZone": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "signatureSet": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "inspectionMode": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "signatureWhiteList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "logLevel": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  logging              []String           `tfsdk:"logging"`
      "targetVpns": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

      vpns              []String           `tfsdk:"vpns"`
      "umbrellaDefault": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "dnsServerIP": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "localDomainBypassEnabled": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
}),},
      "webCategoriesAction": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  webCategories              []String           `tfsdk:"webCategories"`
      "webReputation": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "urlWhiteList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "urlBlackList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "blockPageAction": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "blockPageContents": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "enableAlerts": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
  alerts              []String           `tfsdk:"alerts"`
      "matchAllVpn": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "fileReputationCloudServer": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "fileReputationEstServer": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "fileReputationAlert": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "fileAnalysisCloudServer": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  fileAnalysisFileTypes              []String           `tfsdk:"fileAnalysisFileTypes"`
      "fileAnalysisAlert": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "fileAnalysisEnabled": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "localDomainBypassList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "dnsCrypt": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "umbrellaData": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ref": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
},
,
},

func (d *templatePolicyDefinitionSrstphoneprofileIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyDefinitionSrstphoneprofileIdResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups

	_, r, err := d.client.ConfigurationPolicyVPNListBuilderApi.GetListsById39(context.Background(), state.ListID.Value).Execute()
	dataStr, err := ioutil.ReadAll(r.Body)
  fmt.Println(string(dataStr))
	data := map[string]interface{}{}
	json.Unmarshal(dataStr, &data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HashiCups Coffees",
			err.Error(),
		)
		return
	}

		resp.Diagnostics.AddWarning(
			"test",
			string(dataStr),
		)
	// Map response body to model

	vpnSiteList := data

	vpnSiteListState := vpnSiteListResourceModel{
      Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
}

	}

	// for _, entry := range vpnSiteList["entries"].([]interface{}) {
	// 	vpnSiteListState.Entries = append(vpnSiteListState.Entries, vpnSiteListEntries{
	// 		VPN: types.String{Value: entry["vpn"].(string)},
	// 	})
	// }

	// for _, references := range vpnSiteList["references"].([]map[string]interface{}) {
	// 	vpnSiteListState.References = append(vpnSiteListState.References, vpnSiteListReference{
	// 		ID:   types.String{Value: references["id"].(string)},
	// 		Type: types.String{Value: ""},
	// 	})
	// }

	state = vpnSiteListState

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
