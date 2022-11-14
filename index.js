const newman = require("newman");
const _ = require("lodash");
const dotenv = require("dotenv");
const fs = require("fs");
const path = require("path");

const Webex = require("webex");
const p4 = require("./p4");
const p3 = require("./p3");
const p5 = require("./p5");
const p6 = require("./p6");
const ansible = require("./ansible")

const yargs = require("yargs/yargs");
const { hideBin } = require("yargs/helpers");
const argv = yargs(hideBin(process.argv))
  .command("generate-terraform", "Generates the sdwan terraform provider")
  .command("generate-ansible", "Generates the sdwan ansible modules")
  .command("generate-sdk", "Generates go and python sdks")
  .command("generate-postman", "Generates postman collection")
  .command("backwards-compatibility", "Validates schema compatibility")
  .parse();

dotenv.config();

const default_env_variables = ["VMANAGEIP", "J_USERNAME", "J_PASSWORD"];
default_env_variables.map((el) => {
  if (process.env[el] == "" || process.env[el] == undefined) {
    console.log("SETTING", el);
    process.env[el] = process.env[`${el}_DEFAULT`];
  }
});

const main = async () => {
  if (
    process.env.COLLECTIONS_DIR == undefined &&
    process.env.NODE_ENV !== "development"
  ) {
    console.log(
      "COLLECTIONS DIR ISN'T SET. USING DEFAULT /root/project/collections"
    );
    process.env.COLLECTIONS_DIR =
      process.env.NODE_ENV === "production"
        ? "/root/project/collections"
        : path.resolve(__dirname, "collections");
  } else {
    console.log(`COLLECTIONS_DIR SET TO: ${process.env.COLLECTIONS_DIR}`);
  }
  if (
    process.env.BASE_DIR == undefined &&
    process.env.NODE_ENV !== "development"
  ) {
    console.log("BASE DIR ISN'T SET. USING DEFAULT /root/project");
    process.env.BASE_DIR =
      process.env.NODE_ENV === "production"
        ? "/root/project"
        : path.resolve(__dirname);
  } else {
    console.log(`BASE_DIR SET TO: ${process.env.BASE_DIR}`);
  }
  // await login();
  if (process.env.CIRCLE_NODE_TOTAL === undefined) {
    process.env.CIRCLE_NODE_TOTAL = "2";
  }
  if (process.env.CIRCLE_NODE_INDEX === undefined) {
    process.env.CIRCLE_NODE_INDEX = "0";
  }
  console.log("CIRCLE NODE TOTAL: ", process.env.CIRCLE_NODE_TOTAL);
  console.log("CIRCLE NODE INDEX: ", process.env.CIRCLE_NODE_INDEX);
  const collection_files = fs
    .readdirSync(process.env.COLLECTIONS_DIR)
    .filter((el) => el.includes(".postman_collection.json"));
  const collections = _.chunk(
    collection_files,
    Math.ceil(collection_files.length / parseInt(process.env.CIRCLE_NODE_TOTAL))
  );
  console.log(collections);

  if (parseInt(process.env.CIRCLE_NODE_INDEX) > collections.length - 1) {
    console.log("No need for this runner");
    process.exit(0);
  }

  console.log(`RUNNER ${process.env.CIRCLE_NODE_INDEX} ACTIVE`);

  console.log(process.env);
  if (
    fs.existsSync(path.resolve(process.env.COLLECTIONS_DIR, "info.login.js"))
  ) {
    console.log("RUNNING GLOBAL LOGIN FUNCTION");
    const global_login = require(path.resolve(
      process.env.COLLECTIONS_DIR,
      "info.login.js"
    ));
    await global_login();
  }

  for (const collection of collections[process.env.CIRCLE_NODE_INDEX]) {
    const collection_name = collection.split(".postman_collection.json")[0];
    if (
      fs.existsSync(
        path.resolve(process.env.COLLECTIONS_DIR, `${collection_name}.login.js`)
      )
    ) {
      console.log("RUNNING COLLECTION SPECIFIC LOGIN FUNCTION");
      const collection_specific_login_path = path.resolve(
        process.env.COLLECTIONS_DIR,
        "collections",
        `${collection_name}.login.js`
      );
      if (fs.existsSync(collection_specific_login_path)) {
        const collection_login = require(collection_specific_login_path);
        await collection_login();
      }
    }
    let options = {
      collection: require(path.resolve(
        process.env.COLLECTIONS_DIR,
        collection
      )),
      insecure: true,
      reporter: ["json-summary"],
    };
    const env_path = path.resolve(
      process.env.COLLECTIONS_DIR,
      `${
        collection.split(".postman_collection.json")[0]
      }.postman_environment.json`
    );
    if (fs.existsSync(env_path)) {
      const env_file = require(env_path);
      options.environment = env_file;
      options.envVar = env_file.values.map((el) => {
        return {
          key: el.key,
          value:
            process.env[`${collection_name}_${el.key.toUpperCase()}`] ||
            process.env[el.key.toUpperCase()] ||
            el.value,
        };
      });
    }
    newman.run(options).on("done", async function (err, summary) {
      if (err !== null) {
        console.log(err);
      }
      summary.run.executions.map((el) => {
        if (el.assertions) {
          el.assertions.map(
            ({ error }) =>
              error &&
              fail({
                name: el.item.name,
                message: JSON.stringify(error, null, 2),
              })
          );
        }
        if (el.request.url.path.includes("j_security_check")) {
          console.log(el.assertions);
          fs.writeFileSync("test.json", JSON.stringify(el));
        }
        // console.log(el.response.code)
        if (el.response && `${el.response.code}`.startsWith("2")) {
          // console.log(el.response.stream.toString())
          //
          console.log("SUCCESS: ", el.request.url.path);
        } else {
          console.log("FAILED: ", el.request.url.path);
          // temporary removed fail
          // fail({ name: el.item.name, message: "response did not return 200" });

          // console.log("error")
        }
      });

      if (process.env.WEBEX_SUCCESS_HOOK == "true") {
        console.log("SENDING WEBEX SUCCESS HOOK");
        const result = webex.messages.create({
          text: `**Pipeline passed**\n
BUILD URL: ${process.env.CIRCLE_BUILD_URL}
BRANCH NAME: ${process.env.CIRCLE_BRANCH}
CIRCLE_JOB: ${process.env.CIRCLE_JOB}
`,
          roomId: process.env.WEBEX_ROOM_ID,
        });
        console.log(result);
      }
      await p4.main();
      await p3();
    });
  }
};

switch (argv._[0]) {
  case "generate-terraform":
    p3();
    break;
  case "generate-sdk":
    p4.main();
    break;
  case "generate-postman":
    p5();
    break;
  case "generate-ansible":
    ansible().then(() => {}).catch(() => {});
    break;
  case "backwards-compatibility":
    p6();
    break;
  case "pipeline1":
    main();
  default:
    main();
}
const webex = Webex.init({
  credentials: {
    access_token: process.env.WEBEX_TOKEN,
  },
});

const fail = ({ name, message }) => {
  console.log(`FAILED WHEN RUNNING ${name}\n${message}`);
  if (process.env.WEBEX_FAILURE_HOOK) {
    webex.messages.create({
      text: `**Error in pipeline**
BUILD URL: ${process.env.CIRCLE_BUILD_URL}
BRANCH NAME: ${process.env.CIRCLE_BRANCH}
CIRCLE_JOB: ${process.env.CIRCLE_JOB}
`,
      roomId: process.env.WEBEX_ROOM_ID,
    });
  }

  // console.log("exiting with error", el);
  process.exit(1);
};

// the easiest way to implement this login function is to transform this entire library into an npm package and export the main function. From there, we pass the login function as a paramter in the main function and call it.
