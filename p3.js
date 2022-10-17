const {
  TerraformGenerator,
  Resource,
  map,
  fn,
} = require("terraform-generator");

const main = () => {
  const tfg = new TerraformGenerator({
    required_version: ">= 0.13",
  });
};

module.exports = main;
if (require.main === module) {
  main();
}
