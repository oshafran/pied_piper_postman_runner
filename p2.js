const fs = require("fs");
const path = require("path");
const postmanToOpenApi = require("postman-to-openapi");
const _ = require("lodash");
const { execSync } = require("child_process");
const dotenv = require("dotenv").config();
const axios = require("axios");
const { parseInt } = require("lodash");
const exec = require("child_process").exec;
const fse = require("fs-extra");
const replaceAll = require("replaceall");
const incrementVersion = (gitCommit = "fix: whatever", currentVersion) => {
  const [major, patch, minor] = currentVersion.split(".");
  if (gitCommit.startsWith("fix:")) {
    return {
      update: true,
      newVersion: `${major}.${patch}.${parseInt(minor) + 1}`,
    };
  } else if (gitCommit.startsWith("feat:")) {
    return {
      update: true,
      newVersion: `${major}.${parseInt(patch) + 1}.0`,
    };
  } else if (gitCommit.startsWith("BREAKING CHANGE: ")) {
    return {
      update: true,
      newVersion: `${parseInt(major) + 1}.0.0`,
    };
  } else {
    return {
      update: false,
      newVersion: currentVersion,
    };
  }
};

const genericResponseType = ({ result }) => {
  result = replaceAll(
    `          content:
            application/json: {}`,
    `          content:
             text/plain:
               schema:
                 type: string
                 example: pong
    `,
    result
  );
  console.log(result);
  return result;
};

const publishToGitHubRepo = ({ newVersion, git_repo, folder }) => {
  execSync(`cd ${process.env.BASE_DIR}/${folder} && git init`);
  execSync(`cd ${process.env.BASE_DIR}/${folder} && git branch -m main`);
  execSync(`cd ${process.env.BASE_DIR}/${folder} && git add .`);
  execSync(
    `cd ${process.env.BASE_DIR}/${folder} && git commit -m "${newVersion.newVersion}"`
  );
  console.log(
    `https://${process.env.GITHUB_USERNAME}:${process.env.GITHUB_TOKEN}@github.com/${git_repo}`
  );
  execSync(
    `cd ${process.env.BASE_DIR}/${folder} && git remote add origin https://${process.env.GITHUB_USERNAME}:${process.env.GITHUB_TOKEN}@github.com/${git_repo}`
  );
  execSync(
    `cd ${process.env.BASE_DIR}/${folder} && git push --set-upstream -f origin main`
  );
  execSync(
    `cd ${process.env.BASE_DIR}/${folder} && git tag v${newVersion.newVersion}`
  );
  execSync(
    `cd ${process.env.BASE_DIR}/${folder} && git push -f origin v${newVersion.newVersion}`
  );
};

const publishGoPackage = async () => {
  console.log(process.env.BASE_DIR);
  if (fs.existsSync(path.resolve(process.env.BASE_DIR, "sdk_go"))) {
    fse.emptyDirSync(path.resolve(process.env.BASE_DIR, "sdk_go"));
    fs.rmdirSync(path.resolve(process.env.BASE_DIR, "sdk_go"), {
      force: true,
    });
  }

  const package_name =
    `${process.env.GIT_USER_ID_GO}/${process.env.GIT_REPO_ID_GO}` ||
    "pied-piper-openapi-client-go";

  let newVersion = { newVersion: "1.0.0", update: true };

  try {
    // rename to pied-piper-sdwan-client
    const data = await axios.get(
      `https://api.github.com/repos/${package_name}/tags`
    );
    if (data.length === 0) {
      throw new Error("No existing tags exist");
    }
    // this will fail and crash if the package does not exist. Too bad
    const version = data.data[0].name.split("v")[1];

    newVersion = incrementVersion(process.env.GIT_COMMIT_DESC, version);
  } catch (e) {
    console.log(e);
    console.log("COULD NOT FIND EXISTING PACKAGE. RELEASING 1.0.0");
  }

  console.log(newVersion);
  if (newVersion.update) {
    execSync(
      `openapi-generator-cli generate --git-repo-id ${process.env.GIT_REPO_ID_GO} --git-user-id ${process.env.GIT_USER_ID_GO} -i ${process.env.BASE_DIR}/openapi/test.openapi.yaml -g go -o ${process.env.BASE_DIR}/sdk_go --additional-properties="projectName=${package_name},packageVersion=${newVersion.newVersion}"`
    );

    publishToGitHubRepo({
      newVersion,
      git_repo: package_name,
      folder: "sdk_go",
    });

    // there are actually is a way of modifying through using the package but it takes extra work and so for now whatever
  } else {
    console.log(
      "skipping as git commit doesn't include (fix, feat or BREAKING)"
    );
  }
};

const publishPythonPackage = async () => {
  if (fs.existsSync(path.resolve(process.env.BASE_DIR, "sdk_python"))) {
    fse.emptyDirSync(path.resolve(process.env.BASE_DIR, "sdk_python"));
    fs.rmdirSync(path.resolve(process.env.BASE_DIR, "sdk_python"), {
      force: true,
    });
  }

  const package_name =
    process.env.PYPI_PACKAGE_NAME || "pied-piper-openapi-client";

  let newVersion = { newVersion: "1.0.0", update: true };

  try {
    // rename to pied-piper-sdwan-client
    const data = await axios.get(`https://pypi.org/pypi/${package_name}/json`);
    // this will fail and crash if the package does not exist. Too bad
    const version = data.data.info.version;

    newVersion = incrementVersion(process.env.GIT_COMMIT_DESC, version);
  } catch (e) {
    console.log("COULD NOT FIND EXISTING PACKAGE. RELEASING 1.0.0");
  }

  console.log(newVersion);
  if (newVersion.update) {
    execSync(
      `openapi-generator-cli generate -i ${process.env.BASE_DIR}/openapi/test.openapi.yaml -g python --git-repo-id ${process.env.GIT_REPO_ID_PYTHON} --git-user-id ${process.env.GIT_USER_ID_PYTHON} -o ${process.env.BASE_DIR}/sdk_python --additional-properties="projectName=${package_name},packageVersion=${newVersion.newVersion}"`
    );
    // const rest_file = replaceAll(
    //   "ssl.CERT_REQUIRED",
    //   "ssl.CERT_NONE",
    //   String(
    //     fs.readFileSync(
    //       path.resolve(
    //         process.env.BASE_DIR,
    //         "sdk_python",
    //         "openapi_client",
    //         "rest.py"
    //       )
    //     )
    //   )
    // );
    // fs.writeFileSync(
    //   path.resolve(
    //     process.env.BASE_DIR,
    //     "sdk_python",
    //     "openapi_client",
    //     "rest.py"
    //   ),
    //   rest_file
    // );

    publishToGitHubRepo({
      newVersion,
      git_repo: `${process.env.GIT_USER_ID_PYTHON}/${process.env.GIT_REPO_ID_PYTHON}`,
      folder: "sdk_python",
    });
    execSync(
      `cd ${process.env.BASE_DIR}/sdk_python && python setup.py sdist bdist_wheel`
    );
    execSync(
      `TWINE_PASSWORD="${process.env.TWINE_PASSWORD}" TWINE_USERNAME=${process.env.TWINE_USERNAME} python -m twine upload --skip-existing -r pypi ${process.env.BASE_DIR}/sdk_python/dist/*`
    );

    // there are actually is a way of modifying through using the package but it takes extra work and so for now whatever
  } else {
    console.log(
      "skipping as git commit doesn't include (fix, feat or BREAKING)"
    );
  }
};

const main = async () => {
  const collection_files = fs
    .readdirSync(process.env.COLLECTIONS_DIR)
    .filter((el) => el.includes(".postman_collection.json"));

  if (!fs.existsSync(path.resolve(process.env.BASE_DIR, "openapi"))) {
    fs.mkdirSync(path.resolve(process.env.BASE_DIR, "openapi"));
  }

  fs.writeFileSync(
    path.resolve(process.env.BASE_DIR, "openapitools.json"),
    `{
  "$schema": "./node_modules/@openapitools/openapi-generator-cli/config.schema.json",
  "spaces": 2,
  "generator-cli": {
    "version": "6.1.0"
  }
}`
  );

  for (const collection_file of collection_files) {
    const collection_name = collection_file.split(
      ".postman_collection.json"
    )[0];
    const outputFile = path.resolve(
      process.env.BASE_DIR,
      "openapi",
      `${collection_file.split(".postman_collection.json")[0]}.openapi.json`
    );
    let additionalVars = {};
    const env_path = path.resolve(
      process.env.COLLECTIONS_DIR,
      `${collection_name}.postman_environment.json`
    );
    const env_variables_to_include =
      process.env.ENV_VARIABLES_TO_INCLUDE ? process.env.ENV_VARIABLES_TO_INCLUDE.split(",") : [];
    if (fs.existsSync(env_path)) {
      const env_file = require(env_path);
      additionalVars = env_file.values.map((el) => {
        //
        if (env_variables_to_include.includes(el.key)) {
          return {
            [el.key]:
              process.env[`${collection_name}_${el.key.toUpperCase()}`] ||
              process.env[el.key.toUpperCase()] ||
              el.value,
          };
        }
        return {
          [el.key]: "",
        };
      });
    }
    additionalVars = Object.assign({}, ...additionalVars);
    console.log(additionalVars);
    let result = await postmanToOpenApi(
      path.resolve(process.env.COLLECTIONS_DIR, collection_file),
      outputFile,
      {
        defaultTag: "General",
        additionalVars,
        replaceVars: true,
      }
    );

    result = genericResponseType({ result });

    fs.writeFileSync(
      `${process.env.BASE_DIR}/openapi/${collection_name}.openapi.yaml`,
      result
    );

    // execSync(
    //   "openapi-generator-cli generate -i openapi/test.openapi.yaml -g go -o ${folder}"
    // );

    // process.exit(0)
    try {
      await publishGoPackage();
      await publishPythonPackage();
    } catch (e) {
      console.log(e);
    }
  }
};
module.exports = main;
if (require.main === module) {
  main();
}
