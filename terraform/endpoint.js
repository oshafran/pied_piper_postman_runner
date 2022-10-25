const replaceall = require("replaceall");
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
          body += `  ${camelCase(
            key
          )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          body += `  ${camelCase(key)}              []${pascalCase(
            key
          )}Resource           \`tfsdk:"${key}"\`\n`;
          nestedResources +=
            `\ntype ${camelCase(key)}Resource struct{\n` +
            buildResource({ schema: value.items }) +
            "\n";
        }
      } else if (value.type === "object") {
        body += `  ${camelCase(key)}              ${camelCase(
          key
        )}Resource           \`tfsdk:"${key}"\`\n`;
        nestedResources +=
          `\ntype ${camelCase(key)}Resource struct{\n` +
          buildResource({ schema: value }) +
          "\n";
      } else {
        body += `  ${camelCase(key)}              types.${
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

const resourceGenerator = ({ resource_data }) => {};

const schemaGenerator = ({ schema, path_name }) => {
  let body = ``;
  let nestedResources = "";
  if (schema.type === "object") {
    for (const [key, value] of Object.entries(schema.properties)) {
      if (value.type === "array") {
        if (value.items.type === "string") {
          body += `  ${camelCase(
            key
          )}              []String           \`tfsdk:"${key}"\`\n`;
        } else {
          // body += `  ${camelCase(key)}              []${pascalCase(
          //   key
          // )}Resource           \`tfsdk:"${key}"\`\n`;
          // nestedResources +=
          //   `\ntype ${camelCase(key)}Resource struct{\n` +
          //   buildResource({ schema: value.items }) +
          //   "\n";
        }
      } else if (value.type === "object") {
        body += `      "${key}": {\n        Description:"",\n        Computed: true,\n        Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
\n`;
        nestedResources += "    " +
          schemaGenerator({ schema: value, path_name }) + ",\n";
      } else {
        body += `      "${key}": {\n        Description: "",\n        Computed: ${schema?.computed?.includes(key) ? "false" : "true"},\n        Type: types.${
          typeMap[value.type]
        }\n      },\n`;
      }
    }
  }

  // console.log(schema);

  return body + "}\n" + nestedResources;
};

const generateEndpoint = ({ open_api_spec }) => {
  let count = 0;
  for (const [path, path_data] of Object.entries(open_api_spec.paths)) {
    console.log(path);
    const path_name = nameGenerator({ path });
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
      data = `
func (d *${camelCase(
        path_name
      )}Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Fetches the list of coffees.",
		Attributes: map[string]tfsdk.Attribute{
`;
      console.log(schemaGenerator({ schema: current_schema, path_name }));

      // process.exit(0);
    }
  }
  console.log(count);
};

module.exports = generateEndpoint;
