
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResourceModel struct{
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
  EthernetInterfaceList              []EthernetInterfaceListResource           `tfsdk:"ethernetInterfaceList"`
  ParcelType              types.String           `tfsdk:"parcelType"`
  Subparcels              []String           `tfsdk:"subparcels"`
}

type ethernetInterfaceListResource struct{
  InterfaceName              types.String           `tfsdk:"interfaceName"`
  PortType              types.String           `tfsdk:"portType"`
  WanConfiguration              types.String           `tfsdk:"wanConfiguration"`
  IpAssignment              types.String           `tfsdk:"ipAssignment"`
  StaticIpAddress              types.String           `tfsdk:"staticIpAddress"`
  StaticIpAddressSubnetMask              types.String           `tfsdk:"staticIpAddressSubnetMask"`
  StaticRouteIp              types.String           `tfsdk:"staticRouteIp"`
}




func (r *v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_ethernet_ethernet_id"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
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
      "family": {
        Description: "",
        Computed: true,
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

          "ethernetInterfaceList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "interfaceName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "portType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wanConfiguration": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "ipAssignment": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "staticIpAddress": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "staticIpAddressSubnetMask": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "staticRouteIp": {
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

func (d *v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdEthernetEthernetIdResourceModel
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
