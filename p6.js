const fs = require("fs");
const path = require("path");


const dotenv = require("dotenv").config();
const login = require("./utils/login");

const main = async () => {
  const { instance, x_xsrf_token, cookie } = await login();
  const base_url = `https://${process.env.VMANAGEIP}:${process.env.PORT}/dataservice`;
  const openapi_spec = JSON.parse(
    fs.readFileSync(
      path.resolve(process.env.BASE_DIR, "openapi", "test.openapi.json")
    )
  );
  // console.log(openapi_spec.paths);
  for (const [path, { get }] of Object.entries(openapi_spec.paths)) {
    if (get && !path.includes("{")) {
      console.log(path);
      try {
        const data = await instance.get(base_url + path, {
          headers: {
            "X-XSRF-TOKEN": x_xsrf_token,
            Cookie: cookie,
          },
        });
        // console.log(data.data);
      } catch (e) {
        if (e.response.status === 404) {
          throw new Error("PATH DOES NOT EXIST");
        }
        console.log(e.response.status);
        console.log(e.response.data.error);
      }
    }
  }
};

module.exports = main;
if (require.main === module) {
  main();
}
