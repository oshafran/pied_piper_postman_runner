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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &vpnSiteListsResource{}
	_ resource.ResourceWithConfigure   = &vpnSiteListsResource{}
	_ resource.ResourceWithImportState = &vpnSiteListsResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewVPNSiteListResource() resource.Resource {
	return &vpnSiteListsResource{}
}

// vpnSiteListsResource is the data source implementation.
type vpnSiteListsResource struct {
	client *sdwanAPI.APIClient
}

type vpnSiteListResourceModel struct {
	ListID              types.String           `tfsdk:"list_id"`
	Name                types.String           `tfsdk:"name"`
	Description         types.String           `tfsdk:"description"`
	LastUpdated         types.Float64          `tfsdk:"last_updated"`
	Owner               types.String           `tfsdk:"owner"`
	ReferenceCount      types.String            `tfsdk:"reference_count"`
	Type                types.String           `tfsdk:"type"`
	Entries             []vpnSiteListEntries   `tfsdk:"entries"`
	ReadOnly            types.Bool             `tfsdk:"read_only"`
	Version             types.String            `tfsdk:"version"`
	InfoTag             types.String           `tfsdk:"info_tag"`
	IsActivatedByVsmart types.Bool             `tfsdk:"is_activated_by_vsmart"`
	// References          *[]vpnSiteListReference `tfsdk:"references"`
}
type vpnSiteListReference struct {
	ID   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type vpnSiteListEntries struct {
	VPN types.String `tfsdk:"vpn"`
}

func (r *vpnSiteListsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *vpnSiteListsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_site_list"
}

// GetSchema defines the schema for the data source.
func (d *vpnSiteListsResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
      "list_id": {
				Description: "Placeholder identifier attribute.",
				Computed:    true,
				Type:        types.StringType,
			},
			"name": {
				Description: "Product name of the coffee.",
				Type:        types.StringType,
				Computed:    false,

        Required: true,
			},
			"description": {
				Description: "Product description of the coffee.",
				Type:        types.StringType,
				Computed:    false,
        Required: true,
			},
			"type": {
				Description: "Product description of the coffee.",
				Type:        types.StringType,
				Computed:    false,

        Required: true,
			},
			"owner": {
				Description: "Product description of the coffee.",
				Type:        types.StringType,
				Computed:    true,
			},
			"last_updated": {
				Description: "",
				Type:        types.Float64Type,
				Computed:    true,
			},
			"reference_count": {
				Description: "",
				Type:        types.StringType,
				Computed:    true,
			},
			"read_only": {
				Description: "",
				Type:        types.BoolType,
				Computed:    true,
			},
			"version": {
				Description: "",
				Type:        types.StringType,
				Computed:    true,
			},
			"info_tag": {
				Description: "",
				Type:        types.StringType,
				Computed:    true,
			},
			"is_activated_by_vsmart": {
				Description: "",
				Type:        types.BoolType,
				Computed:    true,
			},
			"entries": {
				Description: "",
				Computed:    false,
        Required: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"vpn": {
						Description: "Fun tagline for the coffee.",
						Type:        types.StringType,
						Computed:    false,
            Required: true,
					},
				}),
			},
			// "references": {
			// 	Description: "",
			// 	Computed:    true,
   //      Required: false,
			// 	Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
			// 		"id": {
			// 			Description: "Fun tagline for the coffee.",
			// 			Type:        types.StringType,
			// 			Computed:    true,
   //          
			// 		},
			// 		"type": {
			// 			Description: "Fun tagline for the coffee.",
			// 			Type:        types.StringType,
			// 			Computed:    true,
			// 		},
			// 	}),
			// },
		},
	}, nil
}

// Configure adds the provider configured client to the data source.
func (d *vpnSiteListsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdwanAPI.APIClient)
}

// Read refreshes the Terraform state with the latest data.
func (d *vpnSiteListsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vpnSiteListResourceModel
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
		Name:                types.String{Value: vpnSiteList["name"].(string)},
		Description:         types.String{Value: vpnSiteList["description"].(string)},
		ListID:              types.String{Value: vpnSiteList["listId"].(string)},
		LastUpdated:         types.Float64{Value: vpnSiteList["lastUpdated"].(float64)},
		Owner:               types.String{Value: vpnSiteList["owner"].(string)},
		ReferenceCount:      types.String{Value: strconv.Itoa(int(vpnSiteList["referenceCount"].(float64)))},
		Type:                types.String{Value: ""},
		ReadOnly:            types.Bool{Value: vpnSiteList["readOnly"].(bool)},
		Version:             types.String{Value: vpnSiteList["version"].(string)},
		InfoTag:             types.String{Value: ""},
		IsActivatedByVsmart: types.Bool{Value: vpnSiteList["isActivatedByVsmart"].(bool)},
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

func (r *vpnSiteListsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vpnSiteListResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var entries []map[string]interface{}
	for _, item := range plan.Entries {
		entries = append(entries, map[string]interface{}{
			"vpn": item.VPN.Value,
		})
	}
	body := map[string]interface{}{
		"name":        plan.Name.Value,
		"type":        plan.Type.Value,
		"entries":     entries,
	}

  bodyStringed, _ := json.Marshal(&body)
  
	_, response, err := r.client.ConfigurationPolicyVPNListBuilderApi.CreatePolicyList39(context.Background()).XXSRFTOKEN(token).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVPNListBuilderApi.CreatePolicyList39``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseBodyString, _ := ioutil.ReadAll(response.Body)
	// Create new order
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error() + string(responseBodyString) + string(bodyStringed),

		)
		return
	}

  resp.Diagnostics.AddWarning("Response body string", string(responseBodyString))

	responseBody := map[string]interface{}{}


  fmt.Println(string(responseBodyString))

	err = json.Unmarshal(responseBodyString, &responseBody)

	// Map response body to schema and populate Computed attribute values
	plan.ListID = types.String{Value: responseBody["listId"].(string)}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *vpnSiteListsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	return
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vpnSiteListsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state vpnSiteListResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Delete existing order
	_, err := r.client.ConfigurationPolicyVPNListBuilderApi.DeletePolicyList39(context.Background(), state.ListID.Value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVPNListBuilderApi.DeletePolicyList39``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}

