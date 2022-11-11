#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import absolute_import, division, print_function
import requests

__metaclass__ = type
DOCUMENTATION = r"""
---
module: my_test
"""

EXAMPLES = r"""
---
"""

RETURN = r"""
---
"""

from ansible.module_utils.basic import AnsibleModule

def login(j_username, j_password, v_manage):
    s = requests.Session()

    port = "443" # str | 
    content_type = "application/x-www-form-urlencoded" # str |  (optional)
    data = {"j_username": j_username, "j_password": j_password}
    headers = {'Content-Type': content_type}
    r = s.post(f'{v_manage}:{port}/j_security_check', data=data, headers=headers, verify=False)

    

    return r.headers['set-cookie'], s.get(f'{v_manage}:{port}/dataservice/client/token').text 

def run_module():
    # define available arguments/parameters a user can pass to the module
    module_args = dict(
        username=dict(type="str", required=True),
        password=dict(type="str", required=True),
        url=dict(type="str", required=True),
    )

    # seed the result dict in the object
    # we primarily care about changed and state
    # changed is if this module effectively modified the target
    # state will include any data that you want your module to pass back
    # for consumption, for example, in a subsequent task
    result = dict(jsessionid="", x_xsrf_token="")

    # the AnsibleModule object will be our abstraction working with Ansible
    # this includes instantiation, a couple of common attr would be the
    # args/params passed to the execution, as well as if the module
    # supports check mode
    module = AnsibleModule(argument_spec=module_args, supports_check_mode=True)



    jsessionid, x_xsrf_token = login(module.params["username"], module.params["password"], module.params["url"])
    
    result = dict(jsessionid=jsessionid.split(";")[0], x_xsrf_token=x_xsrf_token)


    # in the event of a successful module execution, you will want to
    # simple AnsibleModule.exit_json(), passing the key/value results
    module.exit_json(**result)


def main():
    run_module()


if __name__ == "__main__":
    main()
