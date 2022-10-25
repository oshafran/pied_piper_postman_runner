const fs = require("fs")
const Converter = require("openapi-to-postmanv2");

const main = async () => {
  openapiData = fs.readFileSync(`${process.env.OPENAPI_DIR}/test.openapi.json`, {
    encoding: "UTF8",
  });
  Converter.convert(
    { type: "string", data: openapiData },
    {},
    (err, conversionResult) => {
      if (!conversionResult.result) {
        console.log("Could not convert", conversionResult.reason);
      } else {
        console.log(
          "The collection object is: ",
          fs.writeFileSync(
            "postman-collection-test.json",
            JSON.stringify(conversionResult.output[0].data)
          )
        );
      }
    }
  );
};
module.exports = main;
