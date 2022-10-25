const terraformGenerator = require("./terraform");

const main = () => {
  terraformGenerator()
};

module.exports = main;
if (require.main === module) {
  main();
}
