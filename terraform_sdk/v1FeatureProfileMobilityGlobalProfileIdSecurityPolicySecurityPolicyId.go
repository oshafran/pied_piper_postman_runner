
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResourceModel struct{
  Type              types.String           `tfsdk:"type"`
  PolicyName              types.String           `tfsdk:"policyName"`
  DefaultAction              types.String           `tfsdk:"defaultAction"`
  PolicyRules              []PolicyRulesResource           `tfsdk:"policyRules"`
}

type policyRulesResource struct{
  ProtocolType              []String           `tfsdk:"protocolType"`
  SourceIp              types.String           `tfsdk:"sourceIp"`
  DestIp              types.String           `tfsdk:"destIp"`
  SourcePort              types.Int64           `tfsdk:"sourcePort"`
  DestPort              types.Int64           `tfsdk:"destPort"`
  Action              types.String           `tfsdk:"action"`
}


func (r *v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_security_policy_security_policy_id"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "defaultAction": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyRules": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

      protocolType              []String           `tfsdk:"protocolType"`
      "sourceIp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "destIp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sourcePort": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "destPort": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "action": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},

func (d *v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdSecurityPolicySecurityPolicyIdResourceModel
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
      Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyName:              types.String{Value: vpnSiteList["key"].(string)}, 
  DefaultAction:              types.String{Value: vpnSiteList["key"].(string)}, 
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
