
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
	_ resource.Resource                = &mdpOnboardNmsIdResource{}
	_ resource.ResourceWithConfigure   = &mdpOnboardNmsIdResource{}
	_ resource.ResourceWithImportState = &mdpOnboardNmsIdResource{}
)

func NewMdpOnboardNmsIdResource() resource.Resource {
	return &mdpOnboardNmsIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type mdpOnboardNmsIdResource struct {
	client *sdwanAPI.APIClient
}
type mdpOnboardNmsIdResourceModel struct{
  NmsId              types.String           `tfsdk:"nmsId"`
  ControllerUuid              types.String           `tfsdk:"controllerUUID"`
  Name              types.String           `tfsdk:"name"`
  Description              types.String           `tfsdk:"description"`
  RequestStream              types.String           `tfsdk:"requestStream"`
  ResponseStream              types.String           `tfsdk:"responseStream"`
  Otp              types.String           `tfsdk:"otp"`
  Uri              types.String           `tfsdk:"uri"`
}

func (r *mdpOnboardNmsIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *mdpOnboardNmsIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mdp_onboard_nms_id"
}

func (d *mdpOnboardNmsIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "nmsId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "controllerUUID": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "requestStream": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "responseStream": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "otp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "uri": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},

func (d *mdpOnboardNmsIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state mdpOnboardNmsIdResourceModel
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
      NmsId:              types.String{Value: vpnSiteList["key"].(string)}, 
  ControllerUuid:              types.String{Value: vpnSiteList["key"].(string)}, 
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
  RequestStream:              types.String{Value: vpnSiteList["key"].(string)}, 
  ResponseStream:              types.String{Value: vpnSiteList["key"].(string)}, 
  Otp:              types.String{Value: vpnSiteList["key"].(string)}, 
  Uri:              types.String{Value: vpnSiteList["key"].(string)}, 
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
