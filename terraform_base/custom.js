const replaceall = require("replaceall");
const fs = require("fs");
const path = require("path");
const fse = require("fs-extra");
YAML = require("yamljs");

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
  float64: "Float64",
  float32: "Float64",
  float: "Float64",
  integer: "Int64",
  boolean: "Bool",
};

const typeMapGo = {
  string: "string",
  int64: "int64",
  int32: "int64",
  float: "float64",
  integer: "int",
  boolean: "bool",
};

const buildResource = ({ schema, prefix = "" }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      console.log(key);
      if (value.type === "array") {
        if (value.items.type === "string") {
          body += `  ${pascalCase(
            key
          )}              []String           \`tfsdk:"${snakeCase(key)}"\`\n`;
        } else {
          body += `  ${pascalCase(key)}              []${camelCase(
            prefix ? prefix + "_" + key : key
          )}Resource           \`tfsdk:"${snakeCase(key)}"\`\n`;
          nestedResources +=
            `\ntype ${camelCase(
              prefix ? prefix + "_" + key : key
            )}Resource struct{\n` +
            buildResource({
              schema: value.items,

              prefix: prefix ? prefix + "_" + key : key,
            }) +
            "\n";
        }
      } else if (value.type === "object") {
        body += `  ${pascalCase(key)}              ${camelCase(
          prefix ? prefix + "_" + key : key
        )}Resource           \`tfsdk:"${snakeCase(key)}"\`\n`;
        nestedResources +=
          `\ntype ${camelCase(
            prefix ? prefix + "_" + key : key
          )}Resource struct{\n` +
          buildResource({
            schema: value,
            prefix: prefix ? prefix + "_" + key : key,
          }) +
          "\n";
      } else {
        body += `  ${pascalCase(key)}              types.${
          typeMap[value.type]
        }           \`tfsdk:"${snakeCase(key)}"\`\n`;
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

const rootStateMapper = ({ schema }) => {
  return `vpnSiteListState := vpnSiteListResourceModel{}`;
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
          nestedResources += `

	for _, entry := range vpnSiteList["${key}"].([]interface{}) {
 	  vpnSiteListState.${pascalCase(key)} = append(vpnSiteListState.${pascalCase(
            key
          )}, vpnSiteList${pascalCase(key)}{
${stateMapper({ schema: value.items })})
	}
`;
        }
      } else if (value.type === "object") {
        body += `state${camelCase(key)}Resource :=  ${camelCase(key)}Resource {

          ${stateMapper({ schema: value })}
        }`;
        nestedResources += `vpnSiteListState.${pascalCase(
          key
        )} = state${camelCase(key)}Resource;`;
      } else {
        body += `  ${pascalCase(key)}:              types.${
          typeMap[value.type]
        }{Value: vpnSiteList["${key}"].(${typeMapGo[value.type]})}, \n`;
      }
    }
  }

  return body + "\n" + nestedResources;
};

const createRequestBodyMapper = ({ schema, root = "plan", ending = "" }) => {
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
          nestedResources += `

	var ${camelCase(key)} []map[string]interface{}
	for _, item := range plan.${
    value["$map"]
      ? value["$map"]
          .split(".")
          .map((el) => pascalCase(el))
          .join(".")
      : pascalCase(key)
  } {
		${camelCase(key)} = append(${camelCase(key)}, map[string]interface{}{
    // doing this will cause issues if there are multiple nested values
${createRequestBodyMapper({ schema: value.items, root: "item", ending: "," })}
		)
	}

  body["${key}"] = ${camelCase(key)}
`;
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
        body += `  "${key}":        ${root}.${
          value["$map"]
            ? value["$map"]
                .split(".")
                .map((el) => pascalCase(el))
                .join(".")
            : pascalCase(key)
        }.Value,\n`;
      }
    }
  }
  return (
    (root == "plan" ? `body := map[string]interface{}{\n` : "") +
    body +
    "}" +
    ending +
    "\n" +
    nestedResources +
    "\n"
  );
};

const createResponseBodyMapper = ({ schema, root = "plan" }) => {
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
        body += `  plan.${
          value["$map"]
            ? value["$map"]
                .split(".")
                .map((el) => pascalCase(el))
                .join(".")
            : pascalCase(key)
        } = types.String{Value: responseBody["${camelCase(key)}"].(string)}\n`;
      }
    }
  }
  return nestedResources + "\n" + body + "\n";
};

const resourceGenerator = ({ resource_data }) => {};

const createGenerator = ({ path_name, steps, schema }) => {
  return `

func (r *${camelCase(
    path_name
  )}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}
func (r *${camelCase(
    path_name
  )}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vpnSiteListResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan



  ${steps
    .map((el) => {
      let body = `
{
      
${createRequestBodyMapper({
  schema: schemaBuilder({ keys: el.requestBody, schema, reverse: true }),
})}

  bodyStringed, _ := json.Marshal(&body)
`;
      if (el.preScript) {
        body += `body = ${pascalCase(el.preScript)}(body, plan)//pre-script\n`;
      }
      body += `  _, response, err := r.client.${el.sdkPackage}.${
        el.sdkFunction
      }(context.Background())${
        el.xsrftoken ? ".XXSRFTOKEN(token)" : ""
      }.Body(body).Execute()
  if err != nil {
	  fmt.Fprintf(os.Stderr, "Error when calling \`ConfigurationPolicyVPNListBuilderApi.CreatePolicyList39\`: %v\\n", err)
	  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\\n", r)
  }\n`;
      if (el.postScript) {
        body += `  //post-script`;
      }
      body += `
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

${createResponseBodyMapper({
  schema: schemaBuilder({ keys: el.responseBody, schema, reverse: true }),
})}
}`;
      return body;
    })
    .join("\n")}
  
	


	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
`;
};

const updateGenerator = ({ schema, path_name }) => {
  return `
func (r *${camelCase(
    path_name
  )}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	return
}
  `;
};

const deleteGenerator = ({ steps, path_name }) => {
  return `
func (r *${camelCase(
    path_name
  )}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state vpnSiteListResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Delete existing order
${steps
  .map(
    ({ sdkPackage, sdkFunction, id }) =>
      `{
_, err := r.client.${sdkPackage}.${sdkFunction}(context.Background(), state.${id[
        "$path"
      ]
        .split(".")
        .map((el) => pascalCase(el))
        .join(".")}.Value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling \`ConfigurationPolicyVPNListBuilderApi.DeletePolicyList39\`: %v\\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\\n", r)
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
  )
  .join("\n")}

}
`;
};

const readGenerator = ({ steps, path_name, schema }) => {
  return `
func (d *${camelCase(
    path_name
  )}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ${camelCase(path_name)}ResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
//
 vpnSiteListState := vpnSiteListResourceModel{}
  ${steps
    .map(({ sdkPackage, sdkFunction, id, responseBody }) => {
      return `
{
	_, r, err := d.client.${sdkPackage}.${sdkFunction}(context.Background(), state.${id[
        "$path"
      ]
        .split(".")
        .map((el) => pascalCase(el))
        .join(".")}.Value).Execute()
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

${stateMapper({
  schema: schemaBuilder({
    schema,
    keys: responseBody,
  }),
})}
}`;
    })
    .join("\n")}




	state = vpnSiteListState

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
`;
};

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
            `      "${snakeCase(
              key
            )}": {\n        Description:"",\n        Computed: false,\n         Optional: true,\n       Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
\n` +
            "    " +
            schemaGenerator({
              schema: value.items,
              path_name,
              custom_ending: "}),",
            }) +
            "},\n";
        }
      } else if (value.type === "object") {
        body +=
          `      "${snakeCase(
            key
          )}": {\n        Description:"",\n         Optional: true,\n       Computed: false,\n        Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
\n` +
          "    " +
          schemaGenerator({
            schema: value,
            path_name,

            custom_ending: "}),",
          }) +
          "},\n";
      } else {
        body += `      "${snakeCase(
          key
        )}": {\n        Description: "",\n        Computed: ${
          value?.computed ? "true" : "false"
        },\n      Required: ${
          value?.computed ? "false" : "true"
        },\n        Type: types.${typeMap[value.type]}Type,\n      },\n`;
      }
    }
  }

  // console.log(schema);

  return body + (custom_ending ? custom_ending : "},\n") + nestedResources;
};

const schemaFetcher = ({ path, schema }) => {
  if (schema[path[0]].type === "object" && schema[path[0]].properties) {
    const new_path = [...path];
    new_path.shift();
    return schemaFetcher({
      schema: schema[path[0]].properties,
      path: new_path,
    });
  } else if (path.length === 1) {
    return schema[path[0]];
  } else {
    throw new Error("Shit");
  }
};

const nestedBlankSchema = ({ schema, path, data }) => {
  if (path.length == 1) {
    return { [path[0]]: data };
  } else if (schema[path[0]].type == "object") {
    const new_path = [...path];
    new_path.shift();
    console.log("DATA: ", data);
    return {
      [path[0]]: {
        type: schema[path[0]].type,
        computed: data.computed || false,
        properties: nestedBlankSchema({
          schema: schema[path[0]].properties,
          path: new_path,
          data,
        }),
      },
    };
  } else {
    throw new Error("shit");
  }
};

const schemaBuilder = ({ schema, keys, type = "object", reverse = false }) => {
  let new_schema = {};
  for (const key of keys) {
    if (typeof key == "string") {
      new_schema[key] = schema[key];
    } else if (
      typeof key == "object" &&
      Object.keys(key).length == 1 &&
      Object.keys(key[Object.keys(key)])[0] == "$path"
    ) {
      const new_key = Object.keys(key)[0];
      if (reverse) {
        new_schema[new_key] = schemaFetcher({
          path: key[Object.keys(key)[0]]["$path"].split("."),
          schema,
        });
        new_schema[new_key]["$map"] = key[Object.keys(key)[0]]["$path"];
      } else {
        const data = schemaFetcher({
          path: key[new_key]["$path"].split("."),
          schema,
        });

        new_schema = _.merge(
          new_schema,
          nestedBlankSchema({
            schema,
            path: key[new_key]["$path"].split("."),
            data,
          })
        );
      }

      // new_schema[new_key] = schemaFetcher({ path: key[Object.keys(key)[0]]["$path"].split("."), schema });
    } else if (typeof key == "object") {
      const new_key = Object.keys(key)[0];

      console.log("NEW KEY:", new_key);
      console.log(key);
      console.log(
        schemaFetcher({
          path: key[new_key]["$path"]?.split(".") || [new_key],
          schema,
        })
      );

      new_schema[new_key] = schemaBuilder({
        schema: schema[new_key].items.properties,
        keys: key[new_key],
        type: schema[new_key].type,
      });
    }
  }
  if (type === "array") {
    return {
      type,
      items: {
        type: "object",
        properties: new_schema,
      },
    };
  }
  return { type, properties: new_schema };
};

const generateEndpoint = ({ path_name, api }) => {
  const scripts = fs.readFileSync(path.resolve(process.env.BASE_DIR, "terraform", "scripts.go"));
  let endpoint = `
package sdwan

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	sdwanAPI "github.com/oshafran/pied-piper-openapi-client-go"
)

${scripts}

var token string;
var (
	_ resource.Resource                = &${path_name}Resource{}
	_ resource.ResourceWithConfigure   = &${path_name}Resource{}
	_ resource.ResourceWithImportState = &${path_name}Resource{}
)

func New${pascalCase(path_name)}Resource() resource.Resource {
	return &${path_name}Resource{}
}

func (d *${path_name}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdwanAPI.APIClient)
}

// vpnSiteListsResource is the data source implementation.
type ${path_name}Resource struct {
	client *sdwanAPI.APIClient
}
`;

  // console.log("current_schema: ", current_schema);
  endpoint +=
    `type ${camelCase(path_name)}ResourceModel struct{\n` +
    buildResource({ schema: api.schema });

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
  endpoint += schemaGenerator({ schema: api.schema, path_name }) + "}, nil\n}";

  // this is temporarily half-assed and will only work for nested objects. Don't have the time to put in a proper solution sadly
  endpoint += readGenerator({
    schema: api.schema.properties,
    path_name,
    steps: api.read.steps,
  });

  endpoint += createGenerator({
    path_name,
    steps: api.create.steps,
    schema: api.schema.properties,
  });

  // endpoint += updateGenerator({ schema: current_schema, path_name });
  endpoint += deleteGenerator({ steps: api.delete.steps, path_name });

  fs.writeFileSync(
    path.resolve(
      process.env.BASE_DIR,
      "terraform_sdk/sdwan",
      `${camelCase(path_name)}.go`
    ),
    endpoint
  );

  // if (!fse.existsSync(path.resolve(__dirname, "../terraform_sdk/scripts"))) {
  //   fs.mkdirSync(path.resolve(__dirname, "../terraform_sdk/scripts"));
  // }
  //
  // fs.copyFileSync(
  //   path.resolve(__dirname, "scripts.go"),
  //   path.resolve(__dirname, "../terraform_sdk/scripts/scripts.go")
  // );

  // process.exit(0);
  // process.exit(0);
};

const main = () => {
  const data = YAML.load(path.resolve(process.env.BASE_DIR, "terraform", "schema.yaml")).endpoints;

  fse.ensureDirSync(path.resolve(process.env.BASE_DIR, "terraform_sdk/sdwan"));
  fse.emptyDirSync(path.resolve(process.env.BASE_DIR, "terraform_sdk/sdwan"));

  // const data = yamlToJson(fs.readFileSync(path.resolve(__dirname, "./schema.yaml")).toString(), {})
  for (const api of data) {
    generateEndpoint({ path_name: api.name, api });
  }
};

module.exports = main;
