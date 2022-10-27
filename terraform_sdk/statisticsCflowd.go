
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
	_ resource.Resource                = &statisticsCflowdResource{}
	_ resource.ResourceWithConfigure   = &statisticsCflowdResource{}
	_ resource.ResourceWithImportState = &statisticsCflowdResource{}
)

func NewStatisticsCflowdResource() resource.Resource {
	return &statisticsCflowdResource{}
}

// vpnSiteListsResource is the data source implementation.
type statisticsCflowdResource struct {
	client *sdwanAPI.APIClient
}
type statisticsCflowdResourceModel struct{
  Query              queryResource           `tfsdk:"query"`
  Aggregation              aggregationResource           `tfsdk:"aggregation"`
}

type queryResource struct{
  Condition              types.String           `tfsdk:"condition"`
  Rules              []RulesResource           `tfsdk:"rules"`
}

type rulesResource struct{
  Value              []String           `tfsdk:"value"`
  Field              types.String           `tfsdk:"field"`
  Type              types.String           `tfsdk:"type"`
  Operator              types.String           `tfsdk:"operator"`
}



type aggregationResource struct{
  Field              []FieldResource           `tfsdk:"field"`
  Histogram              histogramResource           `tfsdk:"histogram"`
  Metrics              []MetricsResource           `tfsdk:"metrics"`
}

type fieldResource struct{
  Property              types.String           `tfsdk:"property"`
  Sequence              types.Int64           `tfsdk:"sequence"`
}


type histogramResource struct{
  Property              types.String           `tfsdk:"property"`
  Type              types.String           `tfsdk:"type"`
  Interval              types.Int64           `tfsdk:"interval"`
  Order              types.String           `tfsdk:"order"`
}


type metricsResource struct{
  Property              types.String           `tfsdk:"property"`
  Type              types.String           `tfsdk:"type"`
}



func (r *statisticsCflowdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *statisticsCflowdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_statistics_cflowd"
}

func (d *statisticsCflowdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

      value              []String           `tfsdk:"value"`
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
      "operator": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},
,
      "aggregation": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "field": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "property": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "sequence": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
}),},
      "histogram": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "property": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "interval": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "order": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "metrics": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "property": {
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
},
,
},

func (d *statisticsCflowdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state statisticsCflowdResourceModel
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
