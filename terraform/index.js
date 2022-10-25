const fs = require("fs");
const path = require("path");
const generateEndpoint = require("./endpoint");

const generateMain = () => {

  const data =  `
package main

import (
	"context"
	"terraform-provider-sdwan/sdwan"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name hashicups
func main() {
	providerserver.Serve(context.Background(), sdwan.New, providerserver.ServeOpts{
		Address: "hashicorp.com/edu/sdwan",
	})
}
`;
  fs.writeFileSync(path.resolve(__dirname, "../", "terraform_sdk", "main.go"), data)
}


const main = () => {
  // generateMain();
  // fs.mkdirSync(path.resolve(__dirname, "../", "terraform_sdk/sdwan"));
  const open_api_spec = JSON.parse(fs.readFileSync(path.resolve(__dirname, "../openapi/test.openapi.json")))

  generateEndpoint({ open_api_spec })

}

module.exports = main;
