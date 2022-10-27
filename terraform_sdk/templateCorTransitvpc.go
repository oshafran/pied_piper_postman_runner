
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
	_ resource.Resource                = &templateCorTransitvpcResource{}
	_ resource.ResourceWithConfigure   = &templateCorTransitvpcResource{}
	_ resource.ResourceWithImportState = &templateCorTransitvpcResource{}
)

func NewTemplateCorTransitvpcResource() resource.Resource {
	return &templateCorTransitvpcResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateCorTransitvpcResource struct {
	client *sdwanAPI.APIClient
}
type templateCorTransitvpcResourceModel struct{
  Id              types.String           `tfsdk:"id"`
  AccountId              types.String           `tfsdk:"accountId"`
  CloudRegion              types.String           `tfsdk:"cloudRegion"`
  CloudType              types.String           `tfsdk:"cloudType"`
  TransitVpcName              types.String           `tfsdk:"transitVpcName"`
  TransitVpcSize              types.String           `tfsdk:"transitVpcSize"`
  AmiId              types.String           `tfsdk:"amiId"`
  DeviceModelType              types.String           `tfsdk:"deviceModelType"`
  TransitVpcSubnet              types.String           `tfsdk:"transitVpcSubnet"`
  MaxHostVpcPerDevicePair              types.Int64           `tfsdk:"maxHostVpcPerDevicePair"`
  DevicePairList              []DevicePairListResource           `tfsdk:"devicePairList"`
}

type devicePairListResource struct{
  DeviceList              []DeviceListResource           `tfsdk:"deviceList"`
  DevicePairId              types.String           `tfsdk:"devicePairId"`
  IsPrimary              types.Bool           `tfsdk:"isPrimary"`
}

type deviceListResource struct{
  Uuid              types.String           `tfsdk:"uuid"`
  Preference              types.String           `tfsdk:"preference"`
}



func (r *templateCorTransitvpcResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateCorTransitvpcResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_cor_transitvpc"
}

func (d *templateCorTransitvpcResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "id": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "accountId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "cloudRegion": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "cloudType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "transitVpcName": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "transitVpcSize": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "amiId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "deviceModelType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "transitVpcSubnet": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "maxHostVpcPerDevicePair": {
        Description: "",
        Computed: false,
        Type: types.Int64,
      },
      "devicePairList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "deviceList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "uuid": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "preference": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "devicePairId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "isPrimary": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
}),},
},

func (d *templateCorTransitvpcResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateCorTransitvpcResourceModel
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
  TransitVpcSize:              types.String{Value: vpnSiteList["key"].(string)}, 
  AmiId:              types.String{Value: vpnSiteList["key"].(string)}, 
  DeviceModelType:              types.String{Value: vpnSiteList["key"].(string)}, 
  TransitVpcSubnet:              types.String{Value: vpnSiteList["key"].(string)}, 
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
