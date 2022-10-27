
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
	_ resource.Resource                = &alarmsResource{}
	_ resource.ResourceWithConfigure   = &alarmsResource{}
	_ resource.ResourceWithImportState = &alarmsResource{}
)

func NewAlarmsResource() resource.Resource {
	return &alarmsResource{}
}

// vpnSiteListsResource is the data source implementation.
type alarmsResource struct {
	client *sdwanAPI.APIClient
}
type alarmsResourceModel struct{
  Query              queryResource           `tfsdk:"query"`
  Fields              []String           `tfsdk:"fields"`
}

type queryResource struct{
  Condition              types.String           `tfsdk:"condition"`
  Rules              []RulesResource           `tfsdk:"rules"`
}

type rulesResource struct{
  Field              types.String           `tfsdk:"field"`
  Type              types.String           `tfsdk:"type"`
  Value              []String           `tfsdk:"value"`
  Operator              types.String           `tfsdk:"operator"`
}



func (r *alarmsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *alarmsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_alarms"
}

func (d *alarmsResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "query": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "condition": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "rules": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "field": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  value              []String           `tfsdk:"value"`
      "operator": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},
,
  fields              []String           `tfsdk:"fields"`
},

func (d *alarmsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state alarmsResourceModel
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
