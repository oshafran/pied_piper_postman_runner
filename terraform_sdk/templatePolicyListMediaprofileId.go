
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
	_ resource.Resource                = &templatePolicyListMediaprofileIdResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyListMediaprofileIdResource{}
	_ resource.ResourceWithImportState = &templatePolicyListMediaprofileIdResource{}
)

func NewTemplatePolicyListMediaprofileIdResource() resource.Resource {
	return &templatePolicyListMediaprofileIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyListMediaprofileIdResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyListMediaprofileIdResourceModel struct{
  ListId              types.String           `tfsdk:"listId"`
  Name              types.String           `tfsdk:"name"`
  Type              types.String           `tfsdk:"type"`
  Description              types.String           `tfsdk:"description"`
  Entries              []EntriesResource           `tfsdk:"entries"`
  LastUpdated              types.Int64           `tfsdk:"lastUpdated"`
  Owner              types.String           `tfsdk:"owner"`
  ReadOnly              types.Bool           `tfsdk:"readOnly"`
  Version              types.String           `tfsdk:"version"`
  InfoTag              types.String           `tfsdk:"infoTag"`
  ReferenceCount              types.Int64           `tfsdk:"referenceCount"`
  References              []String           `tfsdk:"references"`
  IsActivatedByVsmart              types.Bool           `tfsdk:"isActivatedByVsmart"`
  MasterTemplatesAffected              []String           `tfsdk:"masterTemplatesAffected"`
}

type entriesResource struct{
  Community              types.String           `tfsdk:"community"`
}


func (r *templatePolicyListMediaprofileIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyListMediaprofileIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_list_mediaprofile_id"
}

func (d *templatePolicyListMediaprofileIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "listId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
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
      "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "community": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "lastUpdated": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "owner": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "readOnly": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "version": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "infoTag": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "referenceCount": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
  references              []String           `tfsdk:"references"`
      "isActivatedByVsmart": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
  masterTemplatesAffected              []String           `tfsdk:"masterTemplatesAffected"`
},

func (d *templatePolicyListMediaprofileIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyListMediaprofileIdResourceModel
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
      ListId:              types.String{Value: vpnSiteList["key"].(string)}, 
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
  LastUpdated:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  Owner:              types.String{Value: vpnSiteList["key"].(string)}, 
  ReadOnly:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  Version:              types.String{Value: vpnSiteList["key"].(string)}, 
  InfoTag:              types.String{Value: vpnSiteList["key"].(string)}, 
  ReferenceCount:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  IsActivatedByVsmart:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
