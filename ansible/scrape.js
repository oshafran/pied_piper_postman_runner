const axios = require("axios");
const parser = require("node-html-parser");
const parseMD = require("parse-md").default;
const { snakeCase, pascalCase } = require("change-case");
const {
  createMarkdownArrayTable,
  createMarkdownObjectTable,
} = require("parse-markdown-table");

const fs = require("fs");
const fse = require("fs-extra");
const path = require("path");
const { execSync } = require("child_process");
const endpoints = [
  {
    method: "POST",
    url: "/template/policy/list/site",
    name: "create_site",
  },
];

const main = async () => {
  fse.ensureDirSync(path.resolve(__dirname, "library"));
  fse.emptyDirSync(path.resolve(__dirname, "library"));

  // fse.ensureDirSync(path.resolve(__dirname, "sdk"));
  // fse.emptyDirSync(path.resolve(__dirname, "sdk"));
  // execSync(`git clone https://github.com/oshafran/pied-piper-openapi-client-python ${path.resolve(__dirname, "sdk")}`)

  //github.com/oshafran/pied-piper-openapi-client-python
  const res = await axios.get(
    "https://github.com/oshafran/pied-piper-openapi-client-python/tree/main/docs"
  );

  const data = [];

  const allowed_packages = parser
    .parse(res.data)
    .querySelectorAll(".css-truncate.css-truncate-target.d-block.width-fit")
    .map((el) => {
      if (el.toString().includes(".md")) {
        const url = el.toString().split('href="')[1].split('.md">')[0] + ".md";
        console.log(el.toString().split('.md">')[1].split(".md</a>")[0]);
        return {
          package_name: el.toString().split('.md">')[1].split(".md</a>")[0],
        };
        // if (
        //   url ==
        //   "/oshafran/pied-piper-openapi-client-python/blob/main/docs/ConfigurationPolicySiteListBuilderApi.md"
        // ) {
        //   return {
        //     package_name: el.toString().split('.md">')[1].split(".md</a>")[0],
        //     url: el.toString().split('href="')[1].split('.md">')[0] + ".md",
        //   };
        // }
      }
    })
    .filter((el) => el);
  for (const allowed_package of allowed_packages) {
    const res = fs
      .readFileSync(
        path.resolve(
          __dirname,
          "sdk",
          "docs",
          allowed_package.package_name + ".md"
        ),
        "utf-8"
      )
      .toString();
    try {
      const table_md = res
        .split("All URIs are relative to *https://1.1.1.1*\n")[1]
        .split("\n#")[0];
      const table = await createMarkdownArrayTable(table_md);
      for await (const row of table.rows) {
        const url = row[1].split("** ")[1];
        const method = row[1].split("**")[1].split("**")[0];
        console.log(url, method);
        const python_function = row[0].split("[**")[1].split("**]")[0];
        let name = `${method}_${python_function}`.toLowerCase();

        let include = false;
        endpoints.map((el) => {
          if (el.url === url && method == el.method) {
            include = true;
            if (el.name) {
              name = el.name;
            }
            // stop map
            return;
          }
        });
        if (include) {
          data.push({
            method,
            url,
            python_function,
            name,
          });
        }
      }
    } catch (e) {}

    // const res = await axios.get(`https://github.com${allowed_package.url}`);
    // const data = parser
    //   .parse(res.data)
    //   .querySelector("tbody")
    //   .toString()
    //   .split(/<tr>/)
    //   .map((el) => {
    //     if (!el.includes("<tbody>")) {
    //       const url = el.split("</strong> ")[1].split("</td>")[0];
    //       const method = el.split("<td><strong>")[1].split("</strong>")[0];
    //       const python_function = el.split("<strong>")[1].split("</strong>")[0];
    //       let include = false;
    //       let name = `${method}_${python_function}`.toLowerCase();
    //       endpoints.map((el) => {
    //         if (el.url === url && method == el.method) {
    //           include = true;
    //           if (el.name) {
    //             name = el.name;
    //           }
    //           // stop map
    //           return;
    //         }
    //       });
    //       if (include) {
    //         return {
    //           method,
    //           url,
    //           python_function,
    //           name,
    //         };
    //       }
    //     }
    //   })
    //   .filter((el) => el);

    for (const item of data) {
      let base = `
#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import absolute_import, division, print_function
import openapi_client
from openapi_client.api import ${snakeCase(allowed_package.package_name)} 
__metaclass__ = type
DOCUMENTATION = r"""
---
module: ${item.name} 
"""

EXAMPLES = r"""
---
"""

RETURN = r"""
---
"""

from ansible.module_utils.basic import AnsibleModule


def run_module():

    module_args = dict(
        body=dict(type="dict", required=True),
        url=dict(type="str", required=True),
        jsessionid=dict(type="str", required=False, default=False),
        x_xsrf_token=dict(type="str", required=False, default=False),
    )
    result = dict(changed=False)
    module = AnsibleModule(argument_spec=module_args, supports_check_mode=True)
    configuration = openapi_client.Configuration(
        host=f"{module.params['url']}/dataservice"
    )

    configuration.verify_ssl = False 

    configuration.discard_unknown_keys=True
    with openapi_client.ApiClient(configuration) as api_client:
        api_client.default_headers = {
            "Cookie": module.params["jsessionid"],
            "X-XSRF-TOKEN": module.params["x_xsrf_token"],
        }
        try:
            api_instance = ${snakeCase(
              allowed_package.package_name
            )}.${pascalCase(allowed_package.package_name)}(
                api_client
            )
            # Get  Prefix Template List
            body = module.params["body"] 
            result["results"] = api_instance.${item.python_function}(body=body)
            result["changed"] = True
        except openapi_client.ApiException as e:
            result["error"] = e 
    module.exit_json(**result)


def main():
    run_module()


if __name__ == "__main__":
    main()
`;
      fs.writeFileSync(
        path.resolve(__dirname, "library", item.name + ".py"),
        base
      );
    }
  }
  fse.copyFileSync(
    path.resolve(__dirname, "login.py"),
    path.resolve(__dirname, "library", "login.py")
  );
};

main();

module.exports = main;
