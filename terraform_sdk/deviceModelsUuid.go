
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
	_ resource.Resource                = &deviceModelsUuidResource{}
	_ resource.ResourceWithConfigure   = &deviceModelsUuidResource{}
	_ resource.ResourceWithImportState = &deviceModelsUuidResource{}
)

func NewDeviceModelsUuidResource() resource.Resource {
	return &deviceModelsUuidResource{}
}

// vpnSiteListsResource is the data source implementation.
type deviceModelsUuidResource struct {
	client *sdwanAPI.APIClient
}
type deviceModelsUuidResourceModel struct{
  Name              types.String           `tfsdk:"name"`
  DisplayName              types.String           `tfsdk:"displayName"`
  DeviceType              types.String           `tfsdk:"deviceType"`
  IsCliSupported              types.Bool           `tfsdk:"isCliSupported"`
  TemplateSupported              types.Bool           `tfsdk:"templateSupported"`
  DeviceClass              types.String           `tfsdk:"deviceClass"`
  CpuCountAttribute              cpuCountAttributeResource           `tfsdk:"cpuCountAttribute"`
  OnboardCert              types.Bool           `tfsdk:"onboardCert"`
}

type cpuCountAttributeResource struct{
  Enable              types.Bool           `tfsdk:"enable"`
  AttributeField              types.String           `tfsdk:"attributeField"`
}


func (r *deviceModelsUuidResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *deviceModelsUuidResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_models_uuid"
}

func (d *deviceModelsUuidResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "displayName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "isCliSupported": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "templateSupported": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "deviceClass": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "cpuCountAttribute": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "enable": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "attributeField": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "onboardCert": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
},

func (d *deviceModelsUuidResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceModelsUuidResourceModel
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
  DisplayName:              types.String{Value: vpnSiteList["key"].(string)}, 
  DeviceType:              types.String{Value: vpnSiteList["key"].(string)}, 
  IsCliSupported:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  TemplateSupported:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  DeviceClass:              types.String{Value: vpnSiteList["key"].(string)}, 
  OnboardCert:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
