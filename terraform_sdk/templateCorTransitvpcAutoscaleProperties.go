
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
	_ resource.Resource                = &templateCorTransitvpcAutoscalePropertiesResource{}
	_ resource.ResourceWithConfigure   = &templateCorTransitvpcAutoscalePropertiesResource{}
	_ resource.ResourceWithImportState = &templateCorTransitvpcAutoscalePropertiesResource{}
)

func NewTemplateCorTransitvpcAutoscalePropertiesResource() resource.Resource {
	return &templateCorTransitvpcAutoscalePropertiesResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateCorTransitvpcAutoscalePropertiesResource struct {
	client *sdwanAPI.APIClient
}
type templateCorTransitvpcAutoscalePropertiesResourceModel struct{
  Id              types.String           `tfsdk:"id"`
  AccountId              types.String           `tfsdk:"accountId"`
  CloudRegion              types.String           `tfsdk:"cloudRegion"`
  CloudType              types.String           `tfsdk:"cloudType"`
  TransitVpcName              types.String           `tfsdk:"transitVpcName"`
  TransitVpcId              types.String           `tfsdk:"transitVpcId"`
  MaxHostVpcPerDevicePair              types.Int64           `tfsdk:"maxHostVpcPerDevicePair"`
}

func (r *templateCorTransitvpcAutoscalePropertiesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateCorTransitvpcAutoscalePropertiesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_cor_transitvpc_autoscale_properties"
}

func (d *templateCorTransitvpcAutoscalePropertiesResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "id": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "accountId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "cloudRegion": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "cloudType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "transitVpcName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "transitVpcId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "maxHostVpcPerDevicePair": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
},

func (d *templateCorTransitvpcAutoscalePropertiesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateCorTransitvpcAutoscalePropertiesResourceModel
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
      Id:              types.String{Value: vpnSiteList["key"].(string)}, 
  AccountId:              types.String{Value: vpnSiteList["key"].(string)}, 
  CloudRegion:              types.String{Value: vpnSiteList["key"].(string)}, 
  CloudType:              types.String{Value: vpnSiteList["key"].(string)}, 
  TransitVpcName:              types.String{Value: vpnSiteList["key"].(string)}, 
  TransitVpcId:              types.String{Value: vpnSiteList["key"].(string)}, 
  MaxHostVpcPerDevicePair:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
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
