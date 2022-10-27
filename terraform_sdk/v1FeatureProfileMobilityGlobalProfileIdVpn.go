
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdVpnResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdVpnResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdVpnResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdVpnResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdVpnResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdVpnResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdVpnResourceModel struct{
  ParcelId              types.String           `tfsdk:"parcelId"`
  Name              types.String           `tfsdk:"name"`
  Description              types.String           `tfsdk:"description"`
  Family              types.String           `tfsdk:"family"`
  Featureprofiles              []FeatureprofilesResource           `tfsdk:"featureprofiles"`
}

type featureprofilesResource struct{
  ProfileName              types.String           `tfsdk:"profileName"`
  Description              types.String           `tfsdk:"description"`
  ProfileType              types.String           `tfsdk:"profileType"`
  Parcels              []ParcelsResource           `tfsdk:"parcels"`
}

type parcelsResource struct{
  SiteToSiteVpn              siteToSiteVpnResource           `tfsdk:"siteToSiteVpn"`
  IpSecPolicy              ipSecPolicyResource           `tfsdk:"ipSecPolicy"`
  Id              types.String           `tfsdk:"id"`
  Name              types.String           `tfsdk:"name"`
  Variables              []VariablesResource           `tfsdk:"variables"`
  ParcelType              types.String           `tfsdk:"parcelType"`
  Subparcels              []String           `tfsdk:"subparcels"`
}

type siteToSiteVpnResource struct{
  Name              types.String           `tfsdk:"name"`
  RemotePrivateSubnets              types.String           `tfsdk:"remotePrivateSubnets"`
  PreSharedSecret              types.String           `tfsdk:"preSharedSecret"`
  RemotePublicIp              types.String           `tfsdk:"remotePublicIp"`
  TunnelDnsAddress              types.String           `tfsdk:"tunnelDnsAddress"`
  LocalInterface              types.String           `tfsdk:"localInterface"`
  LocalPrivateSubnet              types.String           `tfsdk:"localPrivateSubnet"`
}


type ipSecPolicyResource struct{
  Preset              types.String           `tfsdk:"preset"`
  IkePhase1              ikePhase1Resource           `tfsdk:"ikePhase1"`
  IkePhase2CipherSuite              types.String           `tfsdk:"ikePhase2CipherSuite"`
}

type ikePhase1Resource struct{
  CipherSuite              types.String           `tfsdk:"cipherSuite"`
  IkeVersion              types.Int64           `tfsdk:"ikeVersion"`
  DiffeHellmanGroup              types.String           `tfsdk:"diffeHellmanGroup"`
  RekeyTimer              types.Int64           `tfsdk:"rekeyTimer"`
}



type variablesResource struct{
  VarName              types.String           `tfsdk:"varName"`
  JsonPath              types.String           `tfsdk:"jsonPath"`
}




func (r *v1FeatureProfileMobilityGlobalProfileIdVpnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdVpnResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_vpn"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdVpnResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "parcelId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "name": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "family": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "featureprofiles": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "profileName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "description": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "profileType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "parcels": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "siteToSiteVpn": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "remotePrivateSubnets": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "preSharedSecret": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "remotePublicIp": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "tunnelDnsAddress": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "localInterface": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "localPrivateSubnet": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "ipSecPolicy": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "preset": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "ikePhase1": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "cipherSuite": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "ikeVersion": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "diffeHellmanGroup": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "rekeyTimer": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
},
,
      "ikePhase2CipherSuite": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "id": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "variables": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "varName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "jsonPath": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "parcelType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  subparcels              []String           `tfsdk:"subparcels"`
}),},
}),},
},

func (d *v1FeatureProfileMobilityGlobalProfileIdVpnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdVpnResourceModel
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
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
  Family:              types.String{Value: vpnSiteList["key"].(string)}, 
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
