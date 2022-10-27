
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
	_ resource.Resource                = &v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource{}
)

func NewV1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource() resource.Resource {
	return &v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResourceModel struct{
  ParcelId              types.String           `tfsdk:"parcelId"`
}

func (r *v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_sdwan_transport_transport_id_wan_vpn_vpn_id_interface_ethernet_ethernet_id"
}

func (d *v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "parcelId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},

func (d *v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileSdwanTransportTransportIdWanVpnVpnIdInterfaceEthernetEthernetIdResourceModel
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
      ParcelId:              types.String{Value: vpnSiteList["key"].(string)}, 
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
