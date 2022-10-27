
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
	_ resource.Resource                = &templatePolicyVsmartPolicyIdResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyVsmartPolicyIdResource{}
	_ resource.ResourceWithImportState = &templatePolicyVsmartPolicyIdResource{}
)

func NewTemplatePolicyVsmartPolicyIdResource() resource.Resource {
	return &templatePolicyVsmartPolicyIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyVsmartPolicyIdResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyVsmartPolicyIdResourceModel struct{
  PolicyId              types.String           `tfsdk:"policyId"`
  PolicyState              types.String           `tfsdk:"policyState"`
  PolicyVersion              types.String           `tfsdk:"policyVersion"`
  LastUpdatedBy              types.String           `tfsdk:"lastUpdatedBy"`
  PolicyName              types.String           `tfsdk:"policyName"`
  PolicyDefinition              types.String           `tfsdk:"policyDefinition"`
  CreatedOn              types.Int64           `tfsdk:"createdOn"`
  IsPolicyActivated              types.Bool           `tfsdk:"isPolicyActivated"`
  PolicyDescription              types.String           `tfsdk:"policyDescription"`
  Rid              types.Int64           `tfsdk:"@rid"`
  CreatedBy              types.String           `tfsdk:"createdBy"`
  PolicyType              types.String           `tfsdk:"policyType"`
  LastUpdatedOn              types.Int64           `tfsdk:"lastUpdatedOn"`
  IsEdited              types.Bool           `tfsdk:"isEdited"`
}

func (r *templatePolicyVsmartPolicyIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyVsmartPolicyIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_vsmart_policy_id"
}

func (d *templatePolicyVsmartPolicyIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "policyId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyState": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyVersion": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "lastUpdatedBy": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyDefinition": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "createdOn": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "isPolicyActivated": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "policyDescription": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "@rid": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "createdBy": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "policyType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "lastUpdatedOn": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "isEdited": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
},

func (d *templatePolicyVsmartPolicyIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyVsmartPolicyIdResourceModel
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
  PolicyState:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyVersion:              types.String{Value: vpnSiteList["key"].(string)}, 
  LastUpdatedBy:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyName:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyDefinition:              types.String{Value: vpnSiteList["key"].(string)}, 
  CreatedOn:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  IsPolicyActivated:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  PolicyDescription:              types.String{Value: vpnSiteList["key"].(string)}, 
  Rid:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  CreatedBy:              types.String{Value: vpnSiteList["key"].(string)}, 
  PolicyType:              types.String{Value: vpnSiteList["key"].(string)}, 
  LastUpdatedOn:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  IsEdited:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
