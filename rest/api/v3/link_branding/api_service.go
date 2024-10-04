/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Link Branding API
* The Twilio SendGrid Link Branding API allows you to configure your domain settings so that all of the click-tracked links, opens, and images in your emails are served from your domain rather than `sendgrid.net`. Spam filters and recipient servers look at the links within emails to determine whether the email appear trustworthy. They use the reputation of the root domain to determine whether the links can be trusted.  You can also manage Link Branding in the **Sender Authentication** section of the Twilio SendGrid application user interface.   See [**How to Set Up Link Branding**](https: //sendgrid.com/docs/ui/account-and-settings/how-to-set-up-link-branding/) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	sendgrid "github.com/sendgrid/sendgrid-go/v4/client"
)

type ApiService struct {
	baseURL        string
	requestHandler *sendgrid.RequestHandler
}

func NewApiService(requestHandler *sendgrid.RequestHandler) *ApiService {
	return &ApiService{
		requestHandler: requestHandler,
		baseURL:        "https://api.sendgrid.com",
	}
}

func NewApiServiceWithClient(client sendgrid.BaseClient) *ApiService {
	return NewApiService(sendgrid.NewRequestHandler(client))
}
