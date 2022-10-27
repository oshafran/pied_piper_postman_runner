
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
	_ resource.Resource                = &certificateDeviceInvalidateResource{}
	_ resource.ResourceWithConfigure   = &certificateDeviceInvalidateResource{}
	_ resource.ResourceWithImportState = &certificateDeviceInvalidateResource{}
)

func NewCertificateDeviceInvalidateResource() resource.Resource {
	return &certificateDeviceInvalidateResource{}
}

// vpnSiteListsResource is the data source implementation.
type certificateDeviceInvalidateResource struct {
	client *sdwanAPI.APIClient
}
type certificateDeviceInvalidateResourceModel struct{
  ChassisNumber              types.String           `tfsdk:"chassisNumber"`
  DeviceIp              types.String           `tfsdk:"deviceIp"`
}

func (r *certificateDeviceInvalidateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *certificateDeviceInvalidateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_device_invalidate"
}

func (d *certificateDeviceInvalidateResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "chassisNumber": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "deviceIp": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
},

func (d *certificateDeviceInvalidateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state certificateDeviceInvalidateResourceModel
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
      ChassisNumber:              types.String{Value: vpnSiteList["key"].(string)}, 
  DeviceIp:              types.String{Value: vpnSiteList["key"].(string)}, 
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
