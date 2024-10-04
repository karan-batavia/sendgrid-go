/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Engagement Quality API
* The SendGrid Engagement Quality (SEQ) API allows you to retrieve metrics that define the quality of your email program.  An SEQ score is correlated with: - The quality of the data in a sender’s program. - How “wanted” the sender's email is by their recipients.  Because “wanted” email and deliverability are closely related, a higher SEQ metric is correlated with greater deliverability. This means the higher your SEQ score, the more likely you are to land in your recipients' inboxes. See the SEQ Overview page to understand SEQ, how it's calculated, and how you can address your score. The SEQ endpoints allow you to retrieve your scores and scores for your Subusers.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"fmt"
	"net/url"
)

type ListSubuserEngagementQualityScoreParam struct {
	// The date in YYYY-MM-DD format (UTC) for which you want to retrieve a SendGrid Engagement Quality score.
	Date *string `json:"date"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	// Specifies the number of results to be returned by the API. This parameter can be used to limit the results returned or in combination with the `after_key` parameter to iterate through paginated results.
	Limit *int32 `json:"limit,omitempty"`
	// Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
	AfterKey *int32 `json:"after_key,omitempty"`
}

func (params *ListSubuserEngagementQualityScoreParam) SetDate(Date string) *ListSubuserEngagementQualityScoreParam {
	params.Date = &Date
	return params
}
func (params *ListSubuserEngagementQualityScoreParam) SetOnbehalfof(Onbehalfof string) *ListSubuserEngagementQualityScoreParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *ListSubuserEngagementQualityScoreParam) SetLimit(Limit int32) *ListSubuserEngagementQualityScoreParam {
	params.Limit = &Limit
	return params
}
func (params *ListSubuserEngagementQualityScoreParam) SetAfterKey(AfterKey int32) *ListSubuserEngagementQualityScoreParam {
	params.AfterKey = &AfterKey
	return params
}

// **This operation allows you to retrieve SendGrid Engagement Quality (SEQ) scores for your Subusers or customer accounts for a specific date.** A successful request with this API operation will return either a `200` or `202` response. ### 202 This operation returns a `202` response when SendGrid does not yet have scores available for the specified date range. Scores are calculated asynchronously from requests to this endpoint. This means a score may be available for the specified date at a later time, but a score is not available at the time of your API request. ### 200 A `200` response will include scores for all Subusers or customer accounts belonging to the requesting parent or reseller account. The `score` and `metrics` properties will be omitted from the response if a Subuser or customer account is not eligible to receive a score for the specified date. The `score` property represents a Subuser or customer account's overall engagement quality. The `metrics` property provides additional scores for the input categories that contribute to that overall score. All scores range from `1` to `5` with a higher number representing better engagement quality. See [**SendGrid Engagement Quality Overview**](https://docs.sendgrid.com/api-reference/sendgrid-engagement-quality-api/overview) for more information
func (c *ApiService) ListSubuserEngagementQualityScore(params *ListSubuserEngagementQualityScoreParam) (interface{}, error) {
	path := "/v3/engagementquality/subusers/scores"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Date != nil {
		data.Set("date", fmt.Sprint(*params.Date))
	}
	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.AfterKey != nil {
		data.Set("after_key", fmt.Sprint(*params.AfterKey))
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return nil, nil
}
