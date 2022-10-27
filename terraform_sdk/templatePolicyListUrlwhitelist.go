
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
	_ resource.Resource                = &templatePolicyListUrlwhitelistResource{}
	_ resource.ResourceWithConfigure   = &templatePolicyListUrlwhitelistResource{}
	_ resource.ResourceWithImportState = &templatePolicyListUrlwhitelistResource{}
)

func NewTemplatePolicyListUrlwhitelistResource() resource.Resource {
	return &templatePolicyListUrlwhitelistResource{}
}

// vpnSiteListsResource is the data source implementation.
type templatePolicyListUrlwhitelistResource struct {
	client *sdwanAPI.APIClient
}
type templatePolicyListUrlwhitelistResourceModel struct{
  ListId              types.String           `tfsdk:"listId"`
  Name              types.String           `tfsdk:"name"`
  Description              types.String           `tfsdk:"description"`
  Type              types.String           `tfsdk:"type"`
  Entries              []EntriesResource           `tfsdk:"entries"`
}

type entriesResource struct{
  Vpn              types.String           `tfsdk:"vpn"`
  GeneratorId              types.String           `tfsdk:"generatorId"`
  SignatureId              types.String           `tfsdk:"signatureId"`
  Pattern              types.String           `tfsdk:"pattern"`
  Region              types.String           `tfsdk:"region"`
  Apikey              types.String           `tfsdk:"apikey"`
  NameServer              types.String           `tfsdk:"nameServer"`
  ApiKey              types.String           `tfsdk:"apiKey"`
  Secret              types.String           `tfsdk:"secret"`
  UmbOrgId              types.String           `tfsdk:"umbOrgId"`
  Token              types.String           `tfsdk:"token"`
  AppFamily              types.String           `tfsdk:"appFamily"`
  App              types.String           `tfsdk:"app"`
  Color              types.String           `tfsdk:"color"`
  IpPrefix              types.String           `tfsdk:"ipPrefix"`
  Burst              types.String           `tfsdk:"burst"`
  Exceed              types.String           `tfsdk:"exceed"`
  Rate              types.String           `tfsdk:"rate"`
  Le              types.String           `tfsdk:"le"`
  SiteId              types.String           `tfsdk:"siteId"`
  Queue              types.String           `tfsdk:"queue"`
  Map              []MapResource           `tfsdk:"map"`
  ForwardingClass              types.String           `tfsdk:"forwardingClass"`
  Latency              types.String           `tfsdk:"latency"`
  Loss              types.String           `tfsdk:"loss"`
  Jitter              types.String           `tfsdk:"jitter"`
  AppProbeClass              types.String           `tfsdk:"appProbeClass"`
  Tloc              types.String           `tfsdk:"tloc"`
  Encap              types.String           `tfsdk:"encap"`
  Preference              types.String           `tfsdk:"preference"`
  AsPath              types.String           `tfsdk:"asPath"`
  Community              types.String           `tfsdk:"community"`
  RemoteDest              types.String           `tfsdk:"remoteDest"`
  Source              types.String           `tfsdk:"source"`
}

type mapResource struct{
  Color              types.String           `tfsdk:"color"`
  Dscp              types.Int64           `tfsdk:"dscp"`
}



func (r *templatePolicyListUrlwhitelistResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *templatePolicyListUrlwhitelistResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template_policy_list_urlwhitelist"
}

func (d *templatePolicyListUrlwhitelistResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "listId": {
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
      "type": {
        Description: "",
        Computed: false,
        Type: types.String,
      },
      "entries": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "vpn": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "generatorId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "signatureId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "pattern": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "region": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "apikey": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "nameServer": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "apiKey": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "secret": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "umbOrgId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "token": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "appFamily": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "app": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "color": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "ipPrefix": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "burst": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "exceed": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "rate": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "le": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "siteId": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "queue": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "map": {
        Description:"",
        Computed: true,
        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

          "color": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "dscp": {
        Description: "",
        Computed: true,
        Type: types.Int64,
      },
}),},
      "forwardingClass": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "latency": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "loss": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "jitter": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "appProbeClass": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "tloc": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "encap": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "preference": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "asPath": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "community": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "remoteDest": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
      "source": {
        Description: "",
        Computed: true,
        Type: types.String,
      },
}),},
},

func (d *templatePolicyListUrlwhitelistResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state templatePolicyListUrlwhitelistResourceModel
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
      ListId:              types.String{Value: vpnSiteList["key"].(string)}, 
  Name:              types.String{Value: vpnSiteList["key"].(string)}, 
  Description:              types.String{Value: vpnSiteList["key"].(string)}, 
  Type:              types.String{Value: vpnSiteList["key"].(string)}, 
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
