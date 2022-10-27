
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
	_ resource.Resource                = &templateDeviceFeatureResource{}
	_ resource.ResourceWithConfigure   = &templateDeviceFeatureResource{}
	_ resource.ResourceWithImportState = &templateDeviceFeatureResource{}
)

func NewTemplateDeviceFeatureResource() resource.Resource {
	return &templateDeviceFeatureResource{}
}

// vpnSiteListsResource is the data source implementation.
type templateDeviceFeatureResource struct {
	client *sdwanAPI.APIClient
}
type templateDeviceFeatureResourceModel struct{
  TemplateId              types.String           `tfsdk:"templateId"`
  TemplateName              types.String           `tfsdk:"templateName"`
  TemplateDescription              types.String           `tfsdk:"templateDescription"`
  DeviceType              types.String           `tfsdk:"deviceType"`
  ConfigType              types.String           `tfsdk:"configType"`
  FactoryDefault              types.Bool           `tfsdk:"factoryDefault"`
  PolicyId              types.String           `tfsdk:"policyId"`
  FeatureTemplateUidRange              []String           `tfsdk:"featureTemplateUidRange"`
  ConnectionPreferenceRequired              types.Bool           `tfsdk:"connectionPreferenceRequired"`
  ConnectionPreference              types.Bool           `tfsdk:"connectionPreference"`
  GeneralTemplates              []GeneralTemplatesResource           `tfsdk:"generalTemplates"`
}

type generalTemplatesResource struct{
  TemplateId              types.String           `tfsdk:"templateId"`
  TemplateType              types.String           `tfsdk:"templateType"`
}


func (r *templateDeviceFeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templateDeviceFeatureResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_device_feature"
}

func (d *templateDeviceFeatureResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "templateId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "templateName": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "templateDescription": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "deviceType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "configType": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "factoryDefault": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "policyId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
  featureTemplateUidRange              []String           `tfsdk:"featureTemplateUidRange"`
      "connectionPreferenceRequired": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "connectionPreference": {
        Description: "",
        Computed: false,
        Type: types.Bool,
      },
      "generalTemplates": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "templateId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "templateType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},

func (d *templateDeviceFeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templateDeviceFeatureResourceModel
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
  TemplateName:              types.String{Value: vpnSiteList["key"].(string)}, 
  TemplateDescription:              types.String{Value: vpnSiteList["key"].(string)}, 
  DeviceType:              types.String{Value: vpnSiteList["key"].(string)}, 
  ConfigType:              types.String{Value: vpnSiteList["key"].(string)}, 
  FactoryDefault:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  PolicyId:              types.String{Value: vpnSiteList["key"].(string)}, 
  ConnectionPreferenceRequired:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
  ConnectionPreference:              types.Bool{Value: vpnSiteList["key"].(boolean)}, 
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