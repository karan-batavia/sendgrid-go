/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Scheduled Sends API
* The Twilio SendGrid Scheduled Sends API allows you to cancel or pause the send of one or more emails using a batch ID.  A `batch_id` groups multiple scheduled mail send requests together with the same ID. You can cancel or pause all of the requests associated with a batch ID up to 10 minutes before the scheduled send time.  See [**Canceling a Scheduled Send**](https://docs.sendgrid.com/for-developers/sending-email/stopping-a-scheduled-send) for a guide on creating a `batch_id`, assigning it to a scheduled send, and modifying the send.  See the Mail API's batching operations to validate a `batch_id` and retrieve all scheduled sends as an array.  When a batch is canceled, all messages associated with that batch will stay in your sending queue. When their `send_at` value is reached, they will be discarded.  When a batch is paused, all messages associated with that batch will stay in your sending queue, even after their `send_at` value has passed. This means you can remove a pause status, and your scheduled send will be delivered once the pause is removed. Any messages left with a pause status that are more than 72 hours old will be discarded as Expired.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

type CreateScheduledSendParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	CancelOrPauseAScheduledSendRequest *CancelOrPauseAScheduledSendRequest `json:"CancelOrPauseAScheduledSendRequest,omitempty"`
}

func (params *CreateScheduledSendParam) SetOnbehalfof(Onbehalfof string) *CreateScheduledSendParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *CreateScheduledSendParam) SetCancelOrPauseAScheduledSendRequest(CancelOrPauseAScheduledSendRequest CancelOrPauseAScheduledSendRequest) *CreateScheduledSendParam {
	params.CancelOrPauseAScheduledSendRequest = &CancelOrPauseAScheduledSendRequest
	return params
}

// **This endpoint allows you to cancel or pause a scheduled send associated with a `batch_id`.**  Passing this endpoint a `batch_id` and status will cancel or pause the scheduled send.  Once a scheduled send is set to `pause` or `cancel` you must use the \"Update a scheduled send\" endpoint to change its status or the \"Delete a cancellation or pause from a scheduled send\" endpoint to remove the status. Passing a status change to a scheduled send that has already been paused or cancelled will result in a `400` level status code.  If the maximum number of cancellations/pauses are added to a send, a `400` level status code will be returned.
func (c *ApiService) CreateScheduledSend(params *CreateScheduledSendParam) (interface{}, error) {
	path := "/v3/user/scheduled_sends"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.CancelOrPauseAScheduledSendRequest != nil {
		b, err := json.Marshal(*params.CancelOrPauseAScheduledSendRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &ScheduledSendStatus{}
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
	if resp.StatusCode == 401 {
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
