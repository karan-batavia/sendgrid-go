/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Legacy Marketing Campaigns Campaigns API
* The Twilio SendGrid Legacy Marketing Campaigns Campaigns API allows you to manage your marketing email messages programmatically. This API is operational, but we recommend using the current version of Marketing Campaigns to manage your marketing messages with SendGrid [Single Sends](https://docs.sendgrid.com/api-reference/single-sends/) and [Automations](https://docs.sendgrid.com/ui/sending-email/getting-started-with-automation).  See [**Migrating from Legacy Marketing Campaigns**](https://docs.sendgrid.com/ui/sending-email/migrating-from-legacy-marketing-campaigns) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

type UpdateScheduledCampaignParam struct {
	//
	CampaignId *int32 `json:"campaign_id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateAScheduledCampaignRequest *UpdateAScheduledCampaignRequest `json:"UpdateAScheduledCampaignRequest,omitempty"`
}

func (params *UpdateScheduledCampaignParam) SetCampaignId(CampaignId int32) *UpdateScheduledCampaignParam {
	params.CampaignId = &CampaignId
	return params
}
func (params *UpdateScheduledCampaignParam) SetOnbehalfof(Onbehalfof string) *UpdateScheduledCampaignParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateScheduledCampaignParam) SetUpdateAScheduledCampaignRequest(UpdateAScheduledCampaignRequest UpdateAScheduledCampaignRequest) *UpdateScheduledCampaignParam {
	params.UpdateAScheduledCampaignRequest = &UpdateAScheduledCampaignRequest
	return params
}

// **This endpoint allows to you change the scheduled time and date for a campaign to be sent.**
func (c *ApiService) UpdateScheduledCampaign(params *UpdateScheduledCampaignParam) (interface{}, error) {
	path := "/v3/campaigns/{CampaignId}/schedules"
	if params != nil && params.CampaignId != nil {
		path = strings.Replace(path, "{"+"CampaignId"+"}", fmt.Sprint(*params.CampaignId), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateAScheduledCampaignRequest != nil {
		b, err := json.Marshal(*params.UpdateAScheduledCampaignRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &UpdateAScheduledCampaignResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 403 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
