
#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import absolute_import, division, print_function
import openapi_client
from openapi_client.api import workflow_management_api 
__metaclass__ = type
DOCUMENTATION = r"""
---
module: create_site 
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
            api_instance = workflow_management_api.WorkflowManagementApi(
                api_client
            )
            # Get  Prefix Template List
            body = module.params["body"] 
            result["results"] = api_instance.create_policy_list30(body=body)
            result["changed"] = True
        except openapi_client.ApiException as e:
            result["error"] = e 
    module.exit_json(**result)


def main():
    run_module()


if __name__ == "__main__":
    main()
