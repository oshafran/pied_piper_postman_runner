
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
	_ resource.Resource                = &templateDeviceConfigAttachcloudxResource{}
	_ resource.ResourceWithConfigure   = &templateDeviceConfigAttachcloudxResource{}
	_ resource.ResourceWithImportState = &templateDeviceConfigAttachcloudxResource{}
)

func NewTemplateDeviceConfigAttachcloudxResource() resource.Resource {
	return &templateDeviceConfigAttachcloudxResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateDeviceConfigAttachcloudxResource struct {
	client *sdwanAPI.APIClient
}
type templateDeviceConfigAttachcloudxResourceModel struct{
  Id              types.String           `tfsdk:"id"`
  IsEdited              types.Bool           `tfsdk:"isEdited"`
  SiteList              []SiteListResource           `tfsdk:"siteList"`
  SiteType              types.String           `tfsdk:"siteType"`
}

type siteListResource struct{
}


func (r *templateDeviceConfigAttachcloudxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateDeviceConfigAttachcloudxResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_device_config_attachcloudx"
}

func (d *templateDeviceConfigAttachcloudxResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "id": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "isEdited": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "siteList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

    }),},
      "siteType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
},

func (d *templateDeviceConfigAttachcloudxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateDeviceConfigAttachcloudxResourceModel
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
  IsEdited:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  SiteType:              types.String{Value: vpnSiteList["key"].(string)}, 
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
