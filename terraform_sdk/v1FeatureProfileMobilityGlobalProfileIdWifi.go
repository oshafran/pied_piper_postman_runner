
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
	_ resource.Resource                = &v1FeatureProfileMobilityGlobalProfileIdWifiResource{}
	_ resource.ResourceWithConfigure   = &v1FeatureProfileMobilityGlobalProfileIdWifiResource{}
	_ resource.ResourceWithImportState = &v1FeatureProfileMobilityGlobalProfileIdWifiResource{}
)

func NewV1FeatureProfileMobilityGlobalProfileIdWifiResource() resource.Resource {
	return &v1FeatureProfileMobilityGlobalProfileIdWifiResource{}
}

// vpnSiteListsResource is the data source implementation.
type v1FeatureProfileMobilityGlobalProfileIdWifiResource struct {
	client *sdwanAPI.APIClient
}
type v1FeatureProfileMobilityGlobalProfileIdWifiResourceModel struct{
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
  SsidConfigList              []SsidConfigListResource           `tfsdk:"ssidConfigList"`
  GuestWifi              guestWifiResource           `tfsdk:"guestWifi"`
  CorporateWifi              corporateWifiResource           `tfsdk:"corporateWifi"`
  AdvancedRadioSetting              advancedRadioSettingResource           `tfsdk:"advancedRadioSetting"`
  ParcelType              types.String           `tfsdk:"parcelType"`
  Subparcels              []String           `tfsdk:"subparcels"`
}

type ssidConfigListResource struct{
  Ssid              types.String           `tfsdk:"ssid"`
  Visibility              types.Bool           `tfsdk:"visibility"`
  SecurityAuthType              types.String           `tfsdk:"securityAuthType"`
  WpaPskKey              types.String           `tfsdk:"wpaPskKey"`
  QosSettings              types.String           `tfsdk:"qosSettings"`
}


type guestWifiResource struct{
  Ssid              types.String           `tfsdk:"ssid"`
  Visibility              types.Bool           `tfsdk:"visibility"`
  SecurityAuthType              types.String           `tfsdk:"securityAuthType"`
  WpaPskKey              types.String           `tfsdk:"wpaPskKey"`
  EncryptionType              types.String           `tfsdk:"encryptionType"`
  WpaEncryptionMode              types.String           `tfsdk:"wpaEncryptionMode"`
  ValidityPeriod              types.String           `tfsdk:"validityPeriod"`
}


type corporateWifiResource struct{
  Ssid              types.String           `tfsdk:"ssid"`
  Visibility              types.Bool           `tfsdk:"visibility"`
  SecurityAuthType              types.String           `tfsdk:"securityAuthType"`
  WpaPskKey              types.String           `tfsdk:"wpaPskKey"`
  EncryptionType              types.String           `tfsdk:"encryptionType"`
  WpaEncryptionMode              types.String           `tfsdk:"wpaEncryptionMode"`
  ValidityPeriod              types.String           `tfsdk:"validityPeriod"`
  CorporateWlan              types.Bool           `tfsdk:"corporateWlan"`
}


type advancedRadioSettingResource struct{
  CountryRegionSettings              countryRegionSettingsResource           `tfsdk:"countryRegionSettings"`
  ChannelPowerSettings              channelPowerSettingsResource           `tfsdk:"channelPowerSettings"`
}

type countryRegionSettingsResource struct{
  CountryRegion              types.String           `tfsdk:"countryRegion"`
  RegulatoryDomain              types.String           `tfsdk:"regulatoryDomain"`
}


type channelPowerSettingsResource struct{
  RadioBand2Dot4Ghz              radioBand2Dot4GhzResource           `tfsdk:"radioBand2Dot4Ghz"`
  RadioBand5Ghz              radioBand5GhzResource           `tfsdk:"radioBand5Ghz"`
}

type radioBand2Dot4GhzResource struct{
  Band              types.String           `tfsdk:"band"`
  Channel              types.String           `tfsdk:"channel"`
  TransmitPower              types.String           `tfsdk:"transmitPower"`
  ChannelWidth              types.String           `tfsdk:"channelWidth"`
}


type radioBand5GhzResource struct{
  Band              types.String           `tfsdk:"band"`
  Channel              types.String           `tfsdk:"channel"`
  TransmitPower              types.String           `tfsdk:"transmitPower"`
  ChannelWidth              types.String           `tfsdk:"channelWidth"`
}






func (r *v1FeatureProfileMobilityGlobalProfileIdWifiResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *v1FeatureProfileMobilityGlobalProfileIdWifiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_feature_profile_mobility_global_profile_id_wifi"
}

func (d *v1FeatureProfileMobilityGlobalProfileIdWifiResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

          "ssidConfigList": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ssid": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "visibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "securityAuthType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wpaPskKey": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "qosSettings": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
      "guestWifi": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ssid": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "visibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "securityAuthType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wpaPskKey": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "encryptionType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wpaEncryptionMode": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "validityPeriod": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "corporateWifi": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "ssid": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "visibility": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
      "securityAuthType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wpaPskKey": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "encryptionType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "wpaEncryptionMode": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "validityPeriod": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "corporateWlan": {
        Description: "",
        Computed: true,
        Type: types.Bool,
      },
},
,
      "advancedRadioSetting": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "countryRegionSettings": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "countryRegion": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "regulatoryDomain": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "channelPowerSettings": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "radioBand2Dot4Ghz": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "band": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "channel": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "transmitPower": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "channelWidth": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
      "radioBand5Ghz": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "band": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "channel": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "transmitPower": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "channelWidth": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
},
,
},
,
},
,
      "parcelType": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
  subparcels              []String           `tfsdk:"subparcels"`
}),},
}),},
},

func (d *v1FeatureProfileMobilityGlobalProfileIdWifiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state v1FeatureProfileMobilityGlobalProfileIdWifiResourceModel
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
