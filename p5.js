const fs = require("fs");
const fse = require("fs-extra");
const path = require("path");
const Converter = require("openapi-to-postmanv2");

const { publishToGitHubRepo } = require("./p4");

const main = async () => {
  openapiData = fs.readFileSync(
    path.resolve(process.env.BASE_DIR, "openapi", "test.openapi.json"),
    {
      encoding: "UTF8",
    }
  );
  const postman_dir = path.resolve(process.env.BASE_DIR, "postman");
  fse.ensureDirSync(postman_dir);
  fse.emptyDirSync(postman_dir);

  Converter.convert(
    { type: "string", data: openapiData },
    {},
    (err, conversionResult) => {
      if (!conversionResult.result) {
        console.log("Could not convert", conversionResult.reason);
      } else {
        fs.writeFileSync(
          path.resolve(postman_dir, "postman-collection-test.json"),
          JSON.stringify(conversionResult.output[0].data)
        );
        publishToGitHubRepo({
          newVersion: {
            newVersion: "1.0.0",
          },
          folder: "postman",
          git_repo: "oshafran/pied-piper-postman-collection",
        });
      }
    }
  );
};
module.exports = main;
