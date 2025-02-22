/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Suppressions API
* The Twilio SendGrid Suppressions API allows you to manage your Suppressions or Unsubscribes and Suppression or Unsubscribe groups. With SendGrid, an unsubscribe is the action an email recipient takes when they opt-out of receiving your messages. A suppression is the action you take as a sender to filter or suppress an unsubscribed address from your email send. From the perspective of the recipient, your suppression is the result of their unsubscribe.  You can have global suppressions, which represent addresses that have been unsubscribed from all of your emails. You can also have suppression groups, also known as ASM groups, which represent categories or groups of emails that your recipients can unsubscribe from, rather than unsubscribing from all of your messages.  SendGrid automatically suppresses emails sent to users for a variety of reasons, including blocks, bounces, invalid email addresses, spam reports, and unsubscribes. SendGrid suppresses these messages to help you maintain the best possible sender reputation by attempting to prevent unwanted mail. You may also add addresses to your suppressions.  See [**Suppressions**](https://docs.sendgrid.com/for-developers/sending-email/suppressions) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type DeleteSpamReportsParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	DeleteSpamReportsRequest *DeleteSpamReportsRequest `json:"DeleteSpamReportsRequest,omitempty"`
}

func (params *DeleteSpamReportsParam) SetOnbehalfof(Onbehalfof string) *DeleteSpamReportsParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *DeleteSpamReportsParam) SetDeleteSpamReportsRequest(DeleteSpamReportsRequest DeleteSpamReportsRequest) *DeleteSpamReportsParam {
	params.DeleteSpamReportsRequest = &DeleteSpamReportsRequest
	return params
}

// **This endpoint allows you to delete your spam reports.**  Deleting a spam report will remove the suppression, meaning email will once again be sent to the previously suppressed address. This should be avoided unless a recipient indicates they wish to receive email from you again. You can use our [bypass filters](https://sendgrid.com/docs/ui/sending-email/index-suppressions/#bypass-suppressions) to deliver messages to otherwise suppressed addresses when exceptions are required.  There are two options for deleting spam reports:   1. You can delete all spam reports by setting the `delete_all` field to `true` in the request body. 2. You can delete a list of select spam reports by specifying the email addresses in the `emails` array of the request body.
func (c *ApiService) DeleteSpamReports(params *DeleteSpamReportsParam) (interface{}, error) {
	path := "/v3/suppression/spam_reports"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.DeleteSpamReportsRequest != nil {
		b, err := json.Marshal(*params.DeleteSpamReportsRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
