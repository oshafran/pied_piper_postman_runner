
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
	_ resource.Resource                = &networkdesignResource{}
	_ resource.ResourceWithConfigure   = &networkdesignResource{}
	_ resource.ResourceWithImportState = &networkdesignResource{}
)

func NewNetworkdesignResource() resource.Resource {
	return &networkdesignResource{}
}

// vpnSiteListsResource is the data source implementation.
type networkdesignResource struct {
	client *sdwanAPI.APIClient
}
type networkdesignResourceModel struct{
  Id              types.String           `tfsdk:"id"`
  Dc              []DcResource           `tfsdk:"dc"`
  ShowDeviceProfileHelpText              types.Bool           `tfsdk:"showDeviceProfileHelpText"`
  GlobalParameters              []GlobalParametersResource           `tfsdk:"globalParameters"`
  CustomizedProfiles              []String           `tfsdk:"customizedProfiles"`
}

type dcResource struct{
  Name              types.String           `tfsdk:"name"`
  Segments              []String           `tfsdk:"segments"`
  DeviceProfiles              []DeviceProfilesResource           `tfsdk:"deviceProfiles"`
}

type deviceProfilesResource struct{
  DeviceProfileName              types.String           `tfsdk:"deviceProfileName"`
  DeviceModel              types.String           `tfsdk:"deviceModel"`
  DeviceTemplateId              types.String           `tfsdk:"deviceTemplateID"`
  DeviceProfileId              types.String           `tfsdk:"deviceProfileId"`
  Circuits              []String           `tfsdk:"circuits"`
}



type globalParametersResource struct{
  TemplateType              types.String           `tfsdk:"templateType"`
  GTemplateClass              types.String           `tfsdk:"gTemplateClass"`
  IsUsed              types.Bool           `tfsdk:"isUsed"`
}


func (r *networkdesignResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *networkdesignResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networkdesign"
}

func (d *networkdesignResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "id": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "dc": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "name": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  segments              []String           `tfsdk:"segments"`
      "deviceProfiles": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "deviceProfileName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceModel": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceTemplateID": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "deviceProfileId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  circuits              []String           `tfsdk:"circuits"`
}),},
}),},
      "showDeviceProfileHelpText": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "globalParameters": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "templateType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "gTemplateClass": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "isUsed": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
}),},
  customizedProfiles              []String           `tfsdk:"customizedProfiles"`
},

func (d *networkdesignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state networkdesignResourceModel
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
  ShowDeviceProfileHelpText:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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
