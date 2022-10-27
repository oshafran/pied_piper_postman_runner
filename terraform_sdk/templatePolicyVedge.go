
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
	_ resource.Resource                = &templatePolicyVedgeResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyVedgeResource{}
	_ resource.ResourceWithImportState = &templatePolicyVedgeResource{}
)

func NewTemplatePolicyVedgeResource() resource.Resource {
	return &templatePolicyVedgeResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyVedgeResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyVedgeResourceModel struct{
  PolicyId              types.String           `tfsdk:"policyId"`
  PolicyDescription              types.String           `tfsdk:"policyDescription"`
  PolicyType              types.String           `tfsdk:"policyType"`
  PolicyName              types.String           `tfsdk:"policyName"`
  PolicyDefinition              policyDefinitionResource           `tfsdk:"policyDefinition"`
  IsPolicyActivated              types.Bool           `tfsdk:"isPolicyActivated"`
}

type policyDefinitionResource struct{
  Assembly              []AssemblyResource           `tfsdk:"assembly"`
  Settings              settingsResource           `tfsdk:"settings"`
}

type assemblyResource struct{
  DefinitionId              types.String           `tfsdk:"definitionId"`
  Type              types.String           `tfsdk:"type"`
}


type settingsResource struct{
  AppVisibility              types.Bool           `tfsdk:"appVisibility"`
  LogFrequency              types.Int64           `tfsdk:"logFrequency"`
}



func (r *templatePolicyVedgeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyVedgeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_vedge"
}

func (d *templatePolicyVedgeResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "policyId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
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
}),},
      "settings": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "appVisibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "logFrequency": {
        Description: "",
        Computed: true,
        Type: types.Int64,
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

func (d *templatePolicyVedgeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyVedgeResourceModel
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
      PolicyId:              types.String{Value: vpnSiteList["key"].(string)}, 
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
