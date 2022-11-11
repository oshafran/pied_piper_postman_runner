const terraformGenerator = require("./terraform_base");

const main = () => {
  terraformGenerator()
};

module.exports = main;
if (require.main === module) {
  main();
}
