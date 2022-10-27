
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
	_ resource.Resource                = &templatePolicyDefinitionSecuritygroupResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyDefinitionSecuritygroupResource{}
	_ resource.ResourceWithImportState = &templatePolicyDefinitionSecuritygroupResource{}
)

func NewTemplatePolicyDefinitionSecuritygroupResource() resource.Resource {
	return &templatePolicyDefinitionSecuritygroupResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyDefinitionSecuritygroupResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyDefinitionSecuritygroupResourceModel struct{
  DefinitionId              types.String           `tfsdk:"definitionId"`
  Name              types.String           `tfsdk:"name"`
  Type              types.String           `tfsdk:"type"`
  Description              types.String           `tfsdk:"description"`
  DefaultAction              defaultActionResource           `tfsdk:"defaultAction"`
  Sequences              []SequencesResource           `tfsdk:"sequences"`
  Definition              definitionResource           `tfsdk:"definition"`
  PolicyDescription              types.String           `tfsdk:"policyDescription"`
  PolicyType              types.String           `tfsdk:"policyType"`
  PolicyName              types.String           `tfsdk:"policyName"`
  PolicyDefinition              policyDefinitionResource           `tfsdk:"policyDefinition"`
  IsPolicyActivated              types.Bool           `tfsdk:"isPolicyActivated"`
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
  Ref              types.String           `tfsdk:"ref"`
}



type actionsResource struct{
  Type              types.String           `tfsdk:"type"`
  Parameter              []ParameterResource           `tfsdk:"parameter"`
}

type parameterResource struct{
  Field              types.String           `tfsdk:"field"`
  Value              valueResource           `tfsdk:"value"`
}

type valueResource struct{
  Exclude              types.String           `tfsdk:"exclude"`
  Prepend              types.String           `tfsdk:"prepend"`
}





type definitionResource struct{
  DefaultAction              defaultActionResource           `tfsdk:"defaultAction"`
  Sequences              []String           `tfsdk:"sequences"`
  Entries              []EntriesResource           `tfsdk:"entries"`
  SignatureSet              types.String           `tfsdk:"signatureSet"`
  InspectionMode              types.String           `tfsdk:"inspectionMode"`
  SignatureWhiteList              signatureWhiteListResource           `tfsdk:"signatureWhiteList"`
  LogLevel              types.String           `tfsdk:"logLevel"`
  Logging              []String           `tfsdk:"logging"`
  TargetVpns              []String           `tfsdk:"targetVpns"`
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
  LocalDomainBypassList              localDomainBypassListResource           `tfsdk:"localDomainBypassList"`
  UmbrellaDefault              types.Bool           `tfsdk:"umbrellaDefault"`
  LocalDomainBypassEnabled              types.Bool           `tfsdk:"localDomainBypassEnabled"`
  DnsCrypt              types.Bool           `tfsdk:"dnsCrypt"`
  UmbrellaData              umbrellaDataResource           `tfsdk:"umbrellaData"`
  VpnList              types.String           `tfsdk:"vpnList"`
  SubDefinitions              []SubDefinitionsResource           `tfsdk:"subDefinitions"`
  Sites              []SitesResource           `tfsdk:"sites"`
  FlowActiveTimeout              types.Int64           `tfsdk:"flowActiveTimeout"`
  FlowInactiveTimeout              types.Int64           `tfsdk:"flowInactiveTimeout"`
  FlowSamplingInterval              types.Int64           `tfsdk:"flowSamplingInterval"`
  TemplateRefresh              types.Int64           `tfsdk:"templateRefresh"`
  Collectors              []CollectorsResource           `tfsdk:"collectors"`
  Protocol              types.String           `tfsdk:"protocol"`
  QosSchedulers              []QosSchedulersResource           `tfsdk:"qosSchedulers"`
  VpnQosSchedulers              []VpnQosSchedulersResource           `tfsdk:"vpnQosSchedulers"`
  Rules              []RulesResource           `tfsdk:"rules"`
}

type defaultActionResource struct{
  Type              types.String           `tfsdk:"type"`
}


type entriesResource struct{
  SourceZone              types.String           `tfsdk:"sourceZone"`
  DestinationZone              types.String           `tfsdk:"destinationZone"`
}


type signatureWhiteListResource struct{
  Ref              types.String           `tfsdk:"ref"`
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


type subDefinitionsResource struct{
  Name              types.String           `tfsdk:"name"`
  EqualPreference              types.Bool           `tfsdk:"equalPreference"`
  AdvertiseTloc              types.Bool           `tfsdk:"advertiseTloc"`
  Spokes              []SpokesResource           `tfsdk:"spokes"`
  TlocList              types.String           `tfsdk:"tlocList"`
}

type spokesResource struct{
  SiteList              types.String           `tfsdk:"siteList"`
  Hubs              []HubsResource           `tfsdk:"hubs"`
}

type hubsResource struct{
  SiteList              types.String           `tfsdk:"siteList"`
  Preference              types.String           `tfsdk:"preference"`
  PrefixLists              []String           `tfsdk:"prefixLists"`
  Ipv6PrefixLists              []String           `tfsdk:"ipv6PrefixLists"`
}




type sitesResource struct{
  SiteList              types.String           `tfsdk:"siteList"`
  VpnList              []String           `tfsdk:"vpnList"`
}


type collectorsResource struct{
  Vpn              types.String           `tfsdk:"vpn"`
  Address              types.String           `tfsdk:"address"`
  Port              types.Int64           `tfsdk:"port"`
  Transport              types.String           `tfsdk:"transport"`
  SourceInterface              types.String           `tfsdk:"sourceInterface"`
}


type qosSchedulersResource struct{
  Queue              types.String           `tfsdk:"queue"`
  BandwidthPercent              types.String           `tfsdk:"bandwidthPercent"`
  BufferPercent              types.String           `tfsdk:"bufferPercent"`
  Burst              types.String           `tfsdk:"burst"`
  Scheduling              types.String           `tfsdk:"scheduling"`
  Drops              types.String           `tfsdk:"drops"`
  ClassMapRef              types.String           `tfsdk:"classMapRef"`
}


type vpnQosSchedulersResource struct{
  VpnListRef              types.String           `tfsdk:"vpnListRef"`
  BandwidthRate              types.String           `tfsdk:"bandwidthRate"`
  ShapingRate              types.String           `tfsdk:"shapingRate"`
  ChildMapRef              types.String           `tfsdk:"childMapRef"`
}


type rulesResource struct{
  Class              types.String           `tfsdk:"class"`
  Plp              types.String           `tfsdk:"plp"`
  Dscp              types.String           `tfsdk:"dscp"`
  Layer2Cos              types.String           `tfsdk:"layer2Cos"`
}



type policyDefinitionResource struct{
  Assembly              []AssemblyResource           `tfsdk:"assembly"`
  Settings              settingsResource           `tfsdk:"settings"`
}

type assemblyResource struct{
  DefinitionId              types.String           `tfsdk:"definitionId"`
  Type              types.String           `tfsdk:"type"`
  Entries              []EntriesResource           `tfsdk:"entries"`
}

type entriesResource struct{
  SiteLists              []String           `tfsdk:"siteLists"`
}



type settingsResource struct{
  FlowVisibility              types.Bool           `tfsdk:"flowVisibility"`
  FlowVisibilityIPv6              types.Bool           `tfsdk:"flowVisibilityIPv6"`
  AppVisibility              types.Bool           `tfsdk:"appVisibility"`
  CloudQos              types.Bool           `tfsdk:"cloudQos"`
  CloudQosServiceSide              types.Bool           `tfsdk:"cloudQosServiceSide"`
  ImplicitAclLogging              types.Bool           `tfsdk:"implicitAclLogging"`
  AppVisibilityIPv6              types.Bool           `tfsdk:"appVisibilityIPv6"`
}



func (r *templatePolicyDefinitionSecuritygroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyDefinitionSecuritygroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_definition_securitygroup"
}

func (d *templatePolicyDefinitionSecuritygroupResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "definitionId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "name": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: false,
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
      "ref": {
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

          "field": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "value": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "exclude": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "prepend": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
}),},
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
  sequences              []String           `tfsdk:"sequences"`
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
  targetVpns              []String           `tfsdk:"targetVpns"`
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
      "umbrellaDefault": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "localDomainBypassEnabled": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
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
      "vpnList": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "subDefinitions": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "equalPreference": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "advertiseTloc": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "spokes": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "siteList": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "hubs": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "siteList": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "preference": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  prefixLists              []String           `tfsdk:"prefixLists"`
  ipv6PrefixLists              []String           `tfsdk:"ipv6PrefixLists"`
}),},
}),},
      "tlocList": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "sites": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "siteList": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  vpnList              []String           `tfsdk:"vpnList"`
}),},
      "flowActiveTimeout": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "flowInactiveTimeout": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "flowSamplingInterval": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "templateRefresh": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "collectors": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "vpn": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "address": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "port": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "transport": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sourceInterface": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "protocol": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "qosSchedulers": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "queue": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "bandwidthPercent": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "bufferPercent": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "burst": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "scheduling": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "drops": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "classMapRef": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "vpnQosSchedulers": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "vpnListRef": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "bandwidthRate": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "shapingRate": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "childMapRef": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "rules": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "class": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "plp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "dscp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "layer2Cos": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},
,
      "policyDescription": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "policyType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "policyName": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "policyDefinition": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "assembly": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "definitionId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

      siteLists              []String           `tfsdk:"siteLists"`
}),},
}),},
      "settings": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "flowVisibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "flowVisibilityIPv6": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "appVisibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "cloudQos": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "cloudQosServiceSide": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "implicitAclLogging": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "appVisibilityIPv6": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
},
,
},
,
      "isPolicyActivated": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
},

func (d *templatePolicyDefinitionSecuritygroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyDefinitionSecuritygroupResourceModel
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
      DefinitionId:              types.String{Value: vpnSiteList["key"].(string)}, 
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyDescription:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyType:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyName:              types.String{Value: vpnSiteList["key"].(string)}, 
  IsPolicyActivated:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
