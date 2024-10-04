/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Designs
* The Twilio SendGrid Designs API offers the ability to manage assets stored in the Twilio SendGrid [Design Library](https://mc.sendgrid.com/design-library/my-designs).  The Design Library is a feature-rich email layout tool and media repository. You can [build designs for all your marketing email needs](https://sendgrid.com/docs/ui/sending-email/working-with-marketing-campaigns-email-designs/), including Single Sends and Automations.  You can also duplicate and then modify one of the pre-built designs provided by Twilio SendGrid to get you started.  The Designs API provides a REST-like interface for creating new designs, retrieving a list of existing designs, duplicating or updating a design, and deleting a design.
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

type DuplicatePreBuiltDesignParam struct {
	// The ID of the pre-built Design you want to duplicate.
	Id *string `json:"id"`
	//
	DesignDuplicateInput *DesignDuplicateInput `json:"DesignDuplicateInput,omitempty"`
}

func (params *DuplicatePreBuiltDesignParam) SetId(Id string) *DuplicatePreBuiltDesignParam {
	params.Id = &Id
	return params
}
func (params *DuplicatePreBuiltDesignParam) SetDesignDuplicateInput(DesignDuplicateInput DesignDuplicateInput) *DuplicatePreBuiltDesignParam {
	params.DesignDuplicateInput = &DesignDuplicateInput
	return params
}

// **This endpoint allows you to duplicate one of the pre-built Twilio SendGrid designs**.  Like duplicating one of your existing designs, you are not required to pass any data in the body of a request to this endpoint. If you choose to leave the `name` field blank, your duplicate will be assigned the name of the design it was copied from with the text \"Duplicate: \" prepended to it. This name change is only a convenience, as the duplicate design will be assigned a unique ID that differentiates it from your other designs. You can retrieve the IDs for Twilio SendGrid pre-built designs using the \"List SendGrid Pre-built Designs\" endpoint.  You can modify your duplicate’s name at the time of creation by passing an updated value to the `name` field when making the initial request. More on retrieving design IDs can be found above.
func (c *ApiService) DuplicatePreBuiltDesign(params *DuplicatePreBuiltDesignParam) (interface{}, error) {
	path := "/v3/designs/pre-builts/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.DesignDuplicateInput != nil {
		b, err := json.Marshal(*params.DesignDuplicateInput)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &DesignOutput{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ApiErrors{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ApiErrors{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
