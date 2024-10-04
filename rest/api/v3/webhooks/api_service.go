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
	sendgrid "github.com/sendgrid/sendgrid-go/client"
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
