const fs = require("fs");
const fse = require("fs-extra");
const path = require("path");
const generateEndpoint = require("./endpoint");
const generateCustom = require("./custom");
const { execSync } = require("child_process");
const { publishToGitHubRepo } = require("../p4");

const generateMain = () => {
  // fse.emptyDirSync(path.resolve(__dirname, "../terraform_sdk"))

  const data = `
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
  fs.writeFileSync(
    path.resolve(__dirname, "../", "terraform_sdk", "main.go"),
    data
  );
};

const main = () => {
  if (
    process.env.TERRAFORM_DIR == undefined &&
    process.env.NODE_ENV !== "development"
  ) {
    console.log(
      "COLLECTIONS DIR ISN'T SET. USING DEFAULT /root/project/collections"
    );
    process.env.TERRAFORM_DIR =
      process.env.NODE_ENV === "production"
        ? "/root/project/collections"
        : path.resolve(__dirname, "collections");
  } else {
    console.log(`TERRAFORM_DIR SET TO: ${process.env.TERRAFORM_DIR}`);
  }

  fse.emptyDirSync(path.resolve(process.env.BASE_DIR, "terraform_sdk"));
  // execSync(
  //   `cd terraform_sdk && git clone git@github.com:oshafran/pied-piper-terraform-provider-gen.git .`
  // );
  // generateCustom();
  // generateMain();
  // fs.mkdirSync(path.resolve(__dirname, "../", "terraform_sdk/sdwan"));
  const open_api_spec = JSON.parse(
    fs.readFileSync(
      path.resolve(process.env.BASE_DIR, "openapi/test.openapi.json")
    )
  );

  generateMain();
  generateCustom();
  fs.copyFileSync(
    path.resolve(__dirname, "provider.go"),
    path.resolve(process.env.BASE_DIR, "terraform_sdk/sdwan/provider.go")
  );
  execSync(`cd terraform_sdk && go mod init terraform-provider-sdwan`);
  execSync(`cd terraform_sdk && go mod tidy`);
  // only needed for dev
  if (process.env.NODE_ENV === "development") {
    execSync(`cd terraform_sdk && go install .`);
  }
  publishToGitHubRepo({
    newVersion: {
      newVersion: "1.0.0",
    },
    folder: "terraform_sdk",
    git_repo: "oshafran/pied-piper-terraform-provider-gen",
  });
  // execSync(`cd terraform_sdk && git add . && git commit -m "test" && git push`)
  // generateEndpoint({ open_api_spec })
};

module.exports = main;
