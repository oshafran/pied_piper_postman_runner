
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResourceModel struct{
  Type              types.String           `tfsdk:"type"`
  DnsSettings              types.String           `tfsdk:"DNSSettings"`
  NtpInherit              types.Bool           `tfsdk:"NTPInherit"`
  NtpSettings              []String           `tfsdk:"NTPSettings"`
  DhcpPool              dhcpPoolResource           `tfsdk:"DHCPPool"`
  DhcpOptions              []DhcpOptionsResource           `tfsdk:"DHCPOptions"`
  NatRules              []NatRulesResource           `tfsdk:"NATRules"`
}

type dhcpPoolResource struct{
  PoolNetwork              types.String           `tfsdk:"poolNetwork"`
  LeaseTime              types.String           `tfsdk:"leaseTime"`
  StartAddress              types.String           `tfsdk:"startAddress"`
  EndAddress              types.String           `tfsdk:"endAddress"`
}


type dhcpOptionsResource struct{
  Type              types.Int64           `tfsdk:"type"`
  DataType              types.String           `tfsdk:"dataType"`
  DhcpValue              types.String           `tfsdk:"dhcpValue"`
}


type natRulesResource struct{
  Protocol              types.String           `tfsdk:"protocol"`
  OutPort              types.Int64           `tfsdk:"outPort"`
  InPort              types.Int64           `tfsdk:"inPort"`
  InsideIp              types.String           `tfsdk:"insideIp"`
  Description              types.String           `tfsdk:"description"`
}


func (r *v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_network_protocol_network_protocol_id"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "type": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "DNSSettings": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "NTPInherit": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
  ntpSettings              []String           `tfsdk:"NTPSettings"`
      "DHCPPool": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "poolNetwork": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "leaseTime": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "startAddress": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "endAddress": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "DHCPOptions": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "type": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "dataType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "dhcpValue": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "NATRules": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "protocol": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "outPort": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "inPort": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "insideIp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},

func (d *v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdNetworkProtocolNetworkProtocolIdResourceModel
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
      Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  DnsSettings:              types.String{Value: vpnSiteList["key"].(string)}, 
  NtpInherit:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
