/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Webhook Configuration API
* The Twilio SendGrid Webhooks API allows you to configure the Event and Parse Webhooks.  Email is a data-rich channel, and implementing the Event Webhook will allow you to consume those data in nearly real time. This means you can actively monitor and participate in the health of your email program throughout the send cycle.  The Inbound Parse Webhook processes all incoming email for a domain or subdomain, parses the contents and attachments and then POSTs `multipart/form-data` to a URL that you choose.
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

type UpdateSignedEventWebhookParam struct {
	// The ID of the Event Webhook you want to retrieve.
	Id *string `json:"id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateSignedEventWebhookRequest *UpdateSignedEventWebhookRequest `json:"UpdateSignedEventWebhookRequest,omitempty"`
}

func (params *UpdateSignedEventWebhookParam) SetId(Id string) *UpdateSignedEventWebhookParam {
	params.Id = &Id
	return params
}
func (params *UpdateSignedEventWebhookParam) SetOnbehalfof(Onbehalfof string) *UpdateSignedEventWebhookParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateSignedEventWebhookParam) SetUpdateSignedEventWebhookRequest(UpdateSignedEventWebhookRequest UpdateSignedEventWebhookRequest) *UpdateSignedEventWebhookParam {
	params.UpdateSignedEventWebhookRequest = &UpdateSignedEventWebhookRequest
	return params
}

// **This endpoint allows you to enable or disable signature verification for a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will enable signature verification for your oldest webhook by `created_date`. This means the default webhook operated on by this endpoint when no ID is provided will be the first one you created. This functionality allows customers who do not have multiple webhooks to enable or disable signature verifiction for their only webhook, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  This endpoint accepts a single boolean request property, `enabled`, that can be set `true` or `false` to enable or disable signature verification. This endpoint will return the public key required to verify Twilio SendGrid signatures if it is enabled or an empty string if signing is disabled. You can also retrieve your public key using the [**Get an Event Webhook's Public Key**](https://docs.sendgrid.com/api-reference/webhooks/get-signed-event-webhooks-public-key) endpoint.  For more information about cryptographically signing the Event Webhook, see [**Getting Started with the Event Webhook Security Features**](https://sendgrid.com/docs/for-developers/tracking-events/getting-started-event-webhook-security-features).
func (c *ApiService) UpdateSignedEventWebhook(params *UpdateSignedEventWebhookParam) (interface{}, error) {
	path := "/v3/user/webhooks/event/settings/signed/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateSignedEventWebhookRequest != nil {
		b, err := json.Marshal(*params.UpdateSignedEventWebhookRequest)
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
		ps := &GetSignedEventWebhook200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &GetSignedEventWebhook404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
