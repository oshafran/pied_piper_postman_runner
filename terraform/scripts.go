
func CustomControl(body map[string]interface{}, plan vpnSiteListResourceModel) map[string]interface{} {
	body["defaultAction"] = map[string]interface{}{
		"type": "reject",
	}
	body["sequences"].([]map[string]interface{})[0]["actions"] = []string{}
	entries := []map[string]interface{}{
		{
			"field": "vpnList",
			"ref":   plan.Vpn.ListId.Value,
		},
	}
	body["sequences"].([]map[string]interface{})[0]["match"] = map[string]interface{}{
		"entries": entries,
	}
	return body
}
