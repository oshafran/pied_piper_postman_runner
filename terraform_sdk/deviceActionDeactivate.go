
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
	_ resource.Resource                = &deviceActionDeactivateResource{}
	_ resource.ResourceWithConfigure   = &deviceActionDeactivateResource{}
	_ resource.ResourceWithImportState = &deviceActionDeactivateResource{}
)

func NewDeviceActionDeactivateResource() resource.Resource {
	return &deviceActionDeactivateResource{}
}

// vpnSiteListsResource is the data source implementation.
type deviceActionDeactivateResource struct {
	client *sdwanAPI.APIClient
}
type deviceActionDeactivateResourceModel struct{
  Id              types.String           `tfsdk:"id"`
  Action              types.String           `tfsdk:"action"`
  Devices              []DevicesResource           `tfsdk:"devices"`
  DeviceType              types.String           `tfsdk:"deviceType"`
}

type devicesResource struct{
  Version              types.String           `tfsdk:"version"`
  DeviceIp              types.String           `tfsdk:"deviceIP"`
  DeviceId              types.String           `tfsdk:"deviceId"`
}


func (r *deviceActionDeactivateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *deviceActionDeactivateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_action_deactivate"
}

func (d *deviceActionDeactivateResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "id": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "action": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "devices": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "version": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceIP": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "deviceType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
},

func (d *deviceActionDeactivateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceActionDeactivateResourceModel
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
  Action:              types.String{Value: vpnSiteList["key"].(string)}, 
  DeviceType:              types.String{Value: vpnSiteList["key"].(string)}, 
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
