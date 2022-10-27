
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
	_ resource.Resource                = &templateDeviceConfigAttachBootStrapResource{}
	_ resource.ResourceWithConfigure   = &templateDeviceConfigAttachBootStrapResource{}
	_ resource.ResourceWithImportState = &templateDeviceConfigAttachBootStrapResource{}
)

func NewTemplateDeviceConfigAttachBootStrapResource() resource.Resource {
	return &templateDeviceConfigAttachBootStrapResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateDeviceConfigAttachBootStrapResource struct {
	client *sdwanAPI.APIClient
}
type templateDeviceConfigAttachBootStrapResourceModel struct{
  DeviceTemplateList              []DeviceTemplateListResource           `tfsdk:"deviceTemplateList"`
}

type deviceTemplateListResource struct{
  TemplateId              types.String           `tfsdk:"templateId"`
  Device              []DeviceResource           `tfsdk:"device"`
  IsEdited              types.Bool           `tfsdk:"isEdited"`
  IsMasterEdited              types.Bool           `tfsdk:"isMasterEdited"`
  IsDraftDisabled              types.Bool           `tfsdk:"isDraftDisabled"`
}

type deviceResource struct{
  CsvStatus              types.String           `tfsdk:"csv-status"`
  CsvDeviceId              types.String           `tfsdk:"csv-deviceId"`
  CsvDeviceIp              types.String           `tfsdk:"csv-deviceIP"`
  CsvHostName              types.String           `tfsdk:"csv-host-name"`
  SystemHostName              types.String           `tfsdk:"//system/host-name"`
  SystemSystemIp              types.String           `tfsdk:"//system/system-ip"`
  SystemSiteId              types.String           `tfsdk:"//system/site-id"`
  CsvTemplateId              types.String           `tfsdk:"csv-templateId"`
  Selected              types.String           `tfsdk:"selected"`
}



func (r *templateDeviceConfigAttachBootStrapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateDeviceConfigAttachBootStrapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_device_config_attach_boot_strap"
}

func (d *templateDeviceConfigAttachBootStrapResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "deviceTemplateList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "templateId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "device": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "csv-status": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "csv-deviceId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "csv-deviceIP": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "csv-host-name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "//system/host-name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "//system/system-ip": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "//system/site-id": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "csv-templateId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "selected": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "isEdited": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "isMasterEdited": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "isDraftDisabled": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
}),},
},

func (d *templateDeviceConfigAttachBootStrapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateDeviceConfigAttachBootStrapResourceModel
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
