/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Address API
* The Twilio SendGrid IP Address API allows you to add and manage dedicated IP addresses and IP Pools for your SendGrid account and Subusers. If you are sending any significant amount of email, SendGrid typically suggests sending from dedicated IPs. It's also best to send marketing and transactional emails from separate IP addresses. In order to do this, you'll need to set up IP Pools, which are groups of dedicated IP addresses you define to send particular types of messages. See the [**Dedicated IP Addresses**](https://docs.sendgrid.com/ui/account-and-settings/dedicated-ip-addresses) for more information about obtaining and allocating IPs.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"

	"strings"
)

type GetIpPoolParam struct {
	// The name of the IP pool that you want to retrieve the IP addresses for.
	PoolName *string `json:"pool_name"`
}

func (params *GetIpPoolParam) SetPoolName(PoolName string) *GetIpPoolParam {
	params.PoolName = &PoolName
	return params
}

// **This endpoint allows you to get all of the IP addresses that are in a specific IP pool.**
func (c *ApiService) GetIpPool(params *GetIpPoolParam) (interface{}, error) {
	path := "/v3/ips/pools/{PoolName}"
	if params != nil && params.PoolName != nil {
		path = strings.Replace(path, "{"+"PoolName"+"}", *params.PoolName, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &GetIpPool200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &GetIpPool404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
