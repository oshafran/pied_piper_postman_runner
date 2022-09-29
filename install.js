const fs = require("fs");
const path = require("path");

const execSync = require("child_process").execSync;
const main = async () => {
  if (fs.existsSync(path.resolve(__dirname, "collections", "packages.json"))) {
    const extra_packages = JSON.parse(
      fs.readFileSync(path.resolve(__dirname, "collections", "packages.json"))
    );

    let package_json = JSON.parse(
      fs.readFileSync(path.resolve(__dirname, "package.json"))
    );
    package_json = {
      ...package_json,
      dependencies: {
        ...package_json.dependencies,
        ...extra_packages
      }
    }

    console.log("installing extra dependencies")
    execSync(`yarn install`);
  }
};

if (require.main === module) {
  main();
}
