
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
	_ resource.Resource                = &templateDeviceConfigExportcsvResource{}
	_ resource.ResourceWithConfigure   = &templateDeviceConfigExportcsvResource{}
	_ resource.ResourceWithImportState = &templateDeviceConfigExportcsvResource{}
)

func NewTemplateDeviceConfigExportcsvResource() resource.Resource {
	return &templateDeviceConfigExportcsvResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateDeviceConfigExportcsvResource struct {
	client *sdwanAPI.APIClient
}
type templateDeviceConfigExportcsvResourceModel struct{
  TemplateId              types.String           `tfsdk:"templateId"`
  IsEdited              types.Bool           `tfsdk:"isEdited"`
  IsMasterEdited              types.Bool           `tfsdk:"isMasterEdited"`
}

func (r *templateDeviceConfigExportcsvResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateDeviceConfigExportcsvResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_device_config_exportcsv"
}

func (d *templateDeviceConfigExportcsvResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "templateId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "isEdited": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "isMasterEdited": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
},

func (d *templateDeviceConfigExportcsvResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateDeviceConfigExportcsvResourceModel
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
      TemplateId:              types.String{Value: vpnSiteList["key"].(string)}, 
  IsEdited:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  IsMasterEdited:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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