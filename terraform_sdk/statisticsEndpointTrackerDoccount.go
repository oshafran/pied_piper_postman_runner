
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
	_ resource.Resource                = &statisticsEndpointTrackerDoccountResource{}
	_ resource.ResourceWithConfigure   = &statisticsEndpointTrackerDoccountResource{}
	_ resource.ResourceWithImportState = &statisticsEndpointTrackerDoccountResource{}
)

func NewStatisticsEndpointTrackerDoccountResource() resource.Resource {
	return &statisticsEndpointTrackerDoccountResource{}
}

// vpnSiteListsResource is the data source implementation.
type statisticsEndpointTrackerDoccountResource struct {
	client *sdwanAPI.APIClient
}
type statisticsEndpointTrackerDoccountResourceModel struct{
  Count              types.Int64           `tfsdk:"count"`
  Query              queryResource           `tfsdk:"query"`
  Fields              []String           `tfsdk:"fields"`
}

type queryResource struct{
  Field              types.String           `tfsdk:"field"`
  Type              types.String           `tfsdk:"type"`
  Value              []String           `tfsdk:"value"`
  Operator              types.String           `tfsdk:"operator"`
}


func (r *statisticsEndpointTrackerDoccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *statisticsEndpointTrackerDoccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_statistics_endpoint_tracker_doccount"
}

func (d *statisticsEndpointTrackerDoccountResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "count": {
        Description: "",
        Computed: false,
        Type: types.Int64,
      },
      "query": {
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
},
,
  fields              []String           `tfsdk:"fields"`
},

func (d *statisticsEndpointTrackerDoccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state statisticsEndpointTrackerDoccountResourceModel
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
      Count:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
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
