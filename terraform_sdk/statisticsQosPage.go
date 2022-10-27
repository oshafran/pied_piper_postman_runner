
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
	_ resource.Resource                = &statisticsQosPageResource{}
	_ resource.ResourceWithConfigure   = &statisticsQosPageResource{}
	_ resource.ResourceWithImportState = &statisticsQosPageResource{}
)

func NewStatisticsQosPageResource() resource.Resource {
	return &statisticsQosPageResource{}
}

// vpnSiteListsResource is the data source implementation.
type statisticsQosPageResource struct {
	client *sdwanAPI.APIClient
}
type statisticsQosPageResourceModel struct{
  Query              queryResource           `tfsdk:"query"`
  Size              types.Int64           `tfsdk:"size"`
  Sort              []SortResource           `tfsdk:"sort"`
  Fields              []String           `tfsdk:"fields"`
}

type queryResource struct{
  Field              types.String           `tfsdk:"field"`
  Type              types.String           `tfsdk:"type"`
  Value              []String           `tfsdk:"value"`
  Operator              types.String           `tfsdk:"operator"`
}


type sortResource struct{
  Field              types.String           `tfsdk:"field"`
  Type              types.String           `tfsdk:"type"`
  Order              types.String           `tfsdk:"order"`
}


func (r *statisticsQosPageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *statisticsQosPageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_statistics_qos_page"
}

func (d *statisticsQosPageResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
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
      "size": {
        Description: "",
        Computed: false,
        Type: types.Int64,
      },
      "sort": {
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
      "order": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
  fields              []String           `tfsdk:"fields"`
},

func (d *statisticsQosPageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state statisticsQosPageResourceModel
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
      Size:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
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
