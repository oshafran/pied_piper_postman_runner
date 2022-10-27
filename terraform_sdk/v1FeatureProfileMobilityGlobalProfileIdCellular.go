
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdCellularResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdCellularResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdCellularResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdCellularResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdCellularResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdCellularResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdCellularResourceModel struct{
  ParcelId              types.String           `tfsdk:"parcelId"`
  Type              types.String           `tfsdk:"type"`
  SimSlot0              simSlot0Resource           `tfsdk:"simSlot0"`
  SimSlot1              simSlot1Resource           `tfsdk:"simSlot1"`
  PrimarySlot              types.Int64           `tfsdk:"primarySlot"`
  Id              types.String           `tfsdk:"id"`
  Name              types.String           `tfsdk:"name"`
}

type simSlot0Resource struct{
  CarrierName              types.String           `tfsdk:"carrierName"`
  SlotNumber              types.Int64           `tfsdk:"slotNumber"`
  ProfileList              []ProfileListResource           `tfsdk:"profileList"`
  DataProfileIdList              []DataProfileIdListResource           `tfsdk:"dataProfileIdList"`
  AttachProfileId              types.Int64           `tfsdk:"attachProfileId"`
}

type profileListResource struct{
  Id              types.Int64           `tfsdk:"id"`
  Apn              types.String           `tfsdk:"apn"`
  PdnType              types.String           `tfsdk:"pdnType"`
  AuthMethod              types.String           `tfsdk:"authMethod"`
  UserName              types.String           `tfsdk:"userName"`
  Password              types.String           `tfsdk:"password"`
}


type dataProfileIdListResource struct{
}



type simSlot1Resource struct{
  CarrierName              types.String           `tfsdk:"carrierName"`
  SlotNumber              types.Int64           `tfsdk:"slotNumber"`
  ProfileList              []ProfileListResource           `tfsdk:"profileList"`
  DataProfileIdList              []DataProfileIdListResource           `tfsdk:"dataProfileIdList"`
  AttachProfileId              types.Int64           `tfsdk:"attachProfileId"`
}

type profileListResource struct{
  Id              types.Int64           `tfsdk:"id"`
  Apn              types.String           `tfsdk:"apn"`
  PdnType              types.String           `tfsdk:"pdnType"`
  AuthMethod              types.String           `tfsdk:"authMethod"`
  UserName              types.String           `tfsdk:"userName"`
  Password              types.String           `tfsdk:"password"`
}


type dataProfileIdListResource struct{
}



func (r *v1FeatureProfileMobilityGlobalProfileIdCellularResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdCellularResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_cellular"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdCellularResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "parcelId": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "type": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "simSlot0": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "carrierName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "slotNumber": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "profileList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "id": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "apn": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "pdnType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "authMethod": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "userName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "password": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "dataProfileIdList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

    }),},
      "attachProfileId": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
},
,
      "simSlot1": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "carrierName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "slotNumber": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "profileList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "id": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
      "apn": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "pdnType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "authMethod": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "userName": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "password": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "dataProfileIdList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

    }),},
      "attachProfileId": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
},
,
      "primarySlot": {
        Description: "",
        Computed: false,
        Type: types.Int64,
      },
      "id": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "name": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
},

func (d *v1FeatureProfileMobilityGlobalProfileIdCellularResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdCellularResourceModel
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
  Type:              types.String{Value: vpnSiteList["key"].(string)}, 
  PrimarySlot:              types.Int64{Value: vpnSiteList["key"].(integer)}, 
  Id:              types.String{Value: vpnSiteList["key"].(string)}, 
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
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
