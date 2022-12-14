const replaceall = require("replaceall");
const fs = require("fs");
const path = require("path");
const {
  snakeCase,
  camelCase,
  capitalCase,
  pascalCase,
} = require("change-case");
const mockToOpenAPI = require("../mock-to-openapi");
const _ = require("lodash");

const typeMap = {
  string: "String",
  int64: "Int64",
  int32: "Int64",
  integer: "Int64",
  boolean: "Bool",
};

const buildResource = ({ schema }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          body += `  ${pascalCase(
            key
          )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          body += `  ${pascalCase(key)}              []${pascalCase(
            key
          )}Resource           \`tfsdk:"${key}"\`\n`;
          nestedResources +=
            `\ntype ${camelCase(key)}Resource struct{\n` +
            buildResource({ schema: value.items }) +
            "\n";
        }
      } else if (value.type === "object") {
        body += `  ${pascalCase(key)}              ${camelCase(
          key
        )}Resource           \`tfsdk:"${key}"\`\n`;
        nestedResources +=
          `\ntype ${camelCase(key)}Resource struct{\n` +
          buildResource({ schema: value }) +
          "\n";
      } else {
        body += `  ${pascalCase(key)}              types.${
          typeMap[value.type]
        }           \`tfsdk:"${key}"\`\n`;
      }
    }
  }

  return body + "}\n" + nestedResources;
};

const requiredAnalyzer = ({ schema, method }) => {
  if (method === "post") {
    if (schema.type === "object") {
      const computed = [];
      for (const [k, v] of Object.entries(schema.properties)) {
        computed.push(k);
      }
      schema.computed = computed;
    }
  }
  return schema;
};

const nameGenerator = ({ path }) => {
  let path_name = replaceall(
    "}",
    "",
    replaceall("{", "", replaceall("/", "_", path.replace("/", "")))
  )
    .split("_")
    .map((el) => {
      return snakeCase(el);
    })
    .join("_");
  return camelCase(path_name);
};

const resourceAnalyzer = ({ resource }) => {
  const resource_schema = {};
  for (const [key, value] of Object.entries(resource)) {
    if (value !== null) {
      if (Array.isArray(value) && value.length > 0) {
        if (typeof value[0] === "string") {
          resource_schema[key] = ["string"];
        } else {
          resource_schema[key] = [resourceAnalyzer({ resource: value[0] })];
        }
      } else if (typeof value === "object") {
        resource_schema[key] = resourceAnalyzer({ resource: value });
      } else {
        resource_schema[key] = typeof value;
      }
    }
  }
  return resource_schema;
};

const stateMapper = ({ schema }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          // body += `  ${pascalCase(
          //   key
          // )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          nestedResources +=
          `

	for _, entry := range vpnSiteList["${key}"].([]interface{}) {
 	  vpnSiteListState.${pascalCase(key)} = append(vpnSiteListState.${pascalCase(key)}, vpnSiteList${pascalCase(key)}{
${stateMapper({ schema: value.items })})
	}
`

        }
      } else if (value.type === "object") {
        // body += `  ${pascalCase(key)}              ${camelCase(
        //   key
        // )}Resource           \`tfsdk:"${key}"\`\n`;
        // nestedResources +=
        //   `\ntype ${camelCase(key)}Resource struct{\n` +
        //   buildResource({ schema: value }) +
        //   "\n";
      } else {
        body += `  ${pascalCase(key)}:              types.${
          typeMap[value.type]
        }{Value: vpnSiteList["${key}"].(${value.type})}, \n`;
      }
    }
  }

  return body + "}\n" + nestedResources;
}


const createRequestBodyMapper = ({ schema, root="plan" }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          // body += `  ${pascalCase(
          //   key
          // )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          nestedResources +=
          `

	var ${camelCase(key)} []map[string]interface{}
	for _, item := range plan.${camelCase(key)} {
		entries = append(entries, map[string]interface{}{
    // doing this will cause issues if there are multiple nested values
      ${createRequestBodyMapper({ schema: value.items, root: "item" })}
		})
	}
`

        }
      } else if (value.type === "object") {
        // body += `  ${pascalCase(key)}              ${camelCase(
        //   key
        // )}Resource           \`tfsdk:"${key}"\`\n`;
        // nestedResources +=
        //   `\ntype ${camelCase(key)}Resource struct{\n` +
        //   buildResource({ schema: value }) +
        //   "\n";
      } else {
        body += `  "${key}":        ${root}.${pascalCase(key)}.Value,\n`;
      }
    }
  }
  return nestedResources + "\n" + (root == "plan" ? `body := map[string]interface{}{\n` : '') + body + "}\n";
}


const createResponseBodyMapper = ({ schema, root="plan" }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          // body += `  ${pascalCase(
          //   key
          // )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {


        }
      } else if (value.type === "object") {
        // body += `  ${pascalCase(key)}              ${camelCase(
        //   key
        // )}Resource           \`tfsdk:"${key}"\`\n`;
        // nestedResources +=
        //   `\ntype ${camelCase(key)}Resource struct{\n` +
        //   buildResource({ schema: value }) +
        //   "\n";
      } else {
        body += `  plan.${pascalCase(key)} = types.String{Value: responseBody["${camelCase(key)}"].(string)}\n`;
      }
    }
  }
  return nestedResources + "\n" + body + "\n";
}

const resourceGenerator = ({ resource_data }) => {};

const createGenerator = ({ schema, path_name, response_schema }) => {
  return `


func (r *vpnSiteListsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vpnSiteListResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan

  ${createRequestBodyMapper({ schema })}

  bodyStringed, _ := json.Marshal(&body)
  
	_, response, err := r.client.ConfigurationPolicyVPNListBuilderApi.CreatePolicyList39(context.Background()).XXSRFTOKEN(token).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling \`ConfigurationPolicyVPNListBuilderApi.CreatePolicyList39\`: %v\n", err)
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

${createResponseBodyMapper({ schema })}
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
`
}

const updateGenerator =({ schema, path_name }) => {
  return `
func (r *${camelCase(path_name)}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	return
}
  `
}

const deleteGenerator = ({ schema, path_name }) => {
  return `
func (r *${camelCase(path_name)}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
		fmt.Fprintf(os.Stderr, "Error when calling \`ConfigurationPolicyVPNListBuilderApi.DeletePolicyList39\`: %v\n", err)
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
`
}

const readGenerator = ({ schema, path_name }) => {
  return `
func (d *${camelCase(path_name)}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ${camelCase(path_name)}ResourceModel
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
    ${stateMapper({ schema })}


	state = vpnSiteListState

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
`
}

const schemaGenerator = ({ schema, path_name, custom_ending }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          // body += `  ${camelCase(
          //   key
          // )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          body +=
            `      "${key}": {\n        Description:"",\n        Computed: true,\n        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
\n` +
            "    " +
            schemaGenerator({ schema: value.items, path_name, custom_ending: "})," }) +
            "},\n";
        }
      } else if (value.type === "object") {
        body +=
          `      "${key}": {\n        Description:"",\n        Computed: true,\n        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
\n` +
          "    " +
          schemaGenerator({ schema: value, path_name }) +
          ",\n";
      } else {
        body += `      "${key}": {\n        Description: "",\n        Computed: ${
          schema?.computed?.includes(key) ? "false" : "true"
        },\n        Type: types.${typeMap[value.type]},\n      },\n`;
      }
    }
  }

  // console.log(schema);

  return body + (custom_ending ? custom_ending : "},\n") + nestedResources;
};

const generateEndpoint = ({ open_api_spec }) => {
  let count = 0;
  for (const [url, path_data] of Object.entries(open_api_spec.paths)) {
    console.log(url);
    const path_name = nameGenerator({ path: url });
    let endpoint = `
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
	_ resource.Resource                = &${path_name}Resource{}
	_ resource.ResourceWithConfigure   = &${path_name}Resource{}
	_ resource.ResourceWithImportState = &${path_name}Resource{}
)

func New${pascalCase(path_name)}Resource() resource.Resource {
	return &${path_name}Resource{}
}

// vpnSiteListsResource is the data source implementation.
type ${path_name}Resource struct {
	client *sdwanAPI.APIClient
}
`;
    let current_schema = {};
    for (const [method, method_data] of Object.entries(path_data)) {
      // if (method === "get") {
      if (
        method_data?.responses["200"]?.content?.["application/json"]?.[
          "schema"
        ]?.["$ref"]
      ) {
        count++;
      }
      if (
        method_data?.responses["200"]?.content?.["application/json"]?.[
          "examples"
        ]
      ) {
        for (const [k, v] of Object.entries(
          method_data?.responses["200"]?.content?.["application/json"]?.[
            "examples"
          ]
        )) {
          if (
            typeof v.value === "object" &&
            !_.isEmpty(v.value) &&
            v.value?.data === undefined
          ) {
            let spec = mockToOpenAPI(v.value);
            spec = requiredAnalyzer({ schema: spec, method });

            current_schema = _.merge(current_schema, spec);

            // console.log(v.value);
            count++;
            // }
          }
        }
        //
        if (
          method_data?.requestBody?.content?.["application/json"]?.["examples"]
        ) {
          for (const [k, v] of Object.entries(
            method_data?.requestBody.content?.["application/json"]?.["examples"]
          )) {
            if (
              typeof v.value === "object" &&
              !_.isEmpty(v.value) &&
              v.value?.data === undefined
            ) {
              let spec = mockToOpenAPI(v.value);
              spec = requiredAnalyzer({ schema: spec, method });

              function customizer(objValue, srcValue) {
                if (_.isArray(objValue)) {
                  return _.uniq(objValue.concat(srcValue));
                }
              }
              current_schema = _.mergeWith(current_schema, spec, customizer);

              // console.log(v.value);
              count++;
              // }
            }
          }
        }
      }
    }
    if (!_.isEmpty(current_schema)) {
      // console.log("current_schema: ", current_schema);
      endpoint +=
        `type ${camelCase(path_name)}ResourceModel struct{\n` +
        buildResource({ schema: current_schema });

      endpoint += `
func (r *${camelCase(
        path_name
      )}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Metadata returns the data source type name.
func (d *${camelCase(
        path_name
      )}Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_${snakeCase(path_name)}"
}
`;
      // console.log(endpoint);
      endpoint += `
func (d *${camelCase(
        path_name
      )}Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
`;
      endpoint += schemaGenerator({ schema: current_schema, path_name });

      endpoint += readGenerator({ schema: current_schema, path_name })
      endpoint += createGenerator({ schema: current_schema, path_name })
      endpoint += updateGenerator({ schema: current_schema, path_name })
      endpoint += deleteGenerator({ schema: current_schema, path_name })
      fs.writeFileSync(
        path.resolve(
          __dirname,
          "../terraform_sdk/",
          `${camelCase(path_name)}.go`
        ),
        endpoint
      );
      // process.exit(0);
    }
  }
  console.log(count);
};

module.exports =  generateEndpoint;
