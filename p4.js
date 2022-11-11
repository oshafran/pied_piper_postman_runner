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
      `openapi-generator-cli generate --skip-validate-spec --git-repo-id ${process.env.GIT_REPO_ID_GO} --git-user-id ${process.env.GIT_USER_ID_GO} -i ${process.env.OPENAPI_DIR}/test.openapi.json -g go -o ${process.env.BASE_DIR}/sdk_go --additional-properties="projectName=${package_name},packageVersion=${newVersion.newVersion},structPrefix=true"`
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
      `openapi-generator-cli generate --skip-validate-spec -i ${process.env.OPENAPI_DIR}/test.openapi.json -g python --git-repo-id ${process.env.GIT_REPO_ID_PYTHON} --git-user-id ${process.env.GIT_USER_ID_PYTHON} -o ${process.env.BASE_DIR}/sdk_python --additional-properties="projectName=${package_name},packageVersion=${newVersion.newVersion}"`,
      { stdio: 'ignore' }
    );
    const rest_file = replaceAll(
      `        try:
            # perform request and return response
            response_data = self.request(`,
      `        try:
            # perform request and return response
            header_params['Accept'] = '*/*'
            response_data = self.request(`,
      String(
        fs.readFileSync(
          path.resolve(
            process.env.BASE_DIR,
            "sdk_python",
            "openapi_client",
            "api_client.py"
          )
        )
      )
    );
    fs.writeFileSync(
      path.resolve(
        process.env.BASE_DIR,
        "sdk_python",
        "openapi_client",
        "api_client.py"
      ),
      rest_file
    );

    // publishToGitHubRepo({
    //   newVersion,
    //   git_repo: `${process.env.GIT_USER_ID_PYTHON}/${process.env.GIT_REPO_ID_PYTHON}`,
    //   folder: "sdk_python",
    // });
    execSync(
      `cd ${process.env.BASE_DIR}/sdk_python && python setup.py sdist bdist_wheel`
    );
    execSync(
      `TWINE_PASSWORD="${process.env.TWINE_PASSWORD}" TWINE_USERNAME=${process.env.TWINE_USERNAME} python -m twine upload --skip-existing -r pypi ${process.env.BASE_DIR}/sdk_python/dist/*`
    );
    console.log("new package published")

    // there are actually is a way of modifying through using the package but it takes extra work and so for now whatever
  } else {
    console.log(
      "skipping as git commit doesn't include (fix, feat or BREAKING)"
    );
  }
};

const substitueENVVariables = () => {
  let openapiData = fs.readFileSync(
    `${process.env.OPENAPI_DIR}/test.openapi.json`,
    {
      encoding: "UTF8",
    }
  );

  openapiData = replaceAll("{{VMANAGEIP}}", process.env.VMANAGEIP, openapiData);
  fs.writeFileSync(`${process.env.OPENAPI_DIR}/test.openapi.json`, openapiData);
};

const main = async () => {
  process.env.VMANAGEIP = `https://${process.env.VMANAGEIP}`;

  console.log("VMANAGE IP IS: ", process.env.VMANAGEIP);
  substitueENVVariables();

  try {
    // await publishGoPackage();
    await publishPythonPackage();
  } catch (e) {
    console.log(e);
    throw new Error(e);
  }
};
if (require.main === module) {
  main();
}

module.exports = { main, publishToGitHubRepo };
