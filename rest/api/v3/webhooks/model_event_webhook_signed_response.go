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
	"time"
)

// EventWebhookSignedResponse struct for EventWebhookSignedResponse
type EventWebhookSignedResponse struct {
	// Indicates if the Event Webhook is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The URL where SendGrid will send event data.
	Url *string `json:"url,omitempty"`
	// Indicates if the webhook is configured to send account status change events related to compliance action taken by SendGrid.
	AccountStatusChange *bool `json:"account_status_change,omitempty"`
	// Indicates if the webhook is configured to send group resubscribe events. Group resubscribes occur when recipients resubscribe to a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	GroupResubscribe *bool `json:"group_resubscribe,omitempty"`
	// Indicates if the webhook is configured to send delivered events. Delivered events occur when a message has been successfully delivered to the receiving server.
	Delivered *bool `json:"delivered,omitempty"`
	// Indicates if the webhook is configured to send group unsubscribe events. Group unsubscribes occur when recipients unsubscribe from a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) either by direct link or by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	GroupUnsubscribe *bool `json:"group_unsubscribe,omitempty"`
	// Indicates if the webhook is configured to send spam report events. Spam reports occur when recipients mark a message as spam.
	SpamReport *bool `json:"spam_report,omitempty"`
	// Indicates if the webhook is configured to send bounce events. A bounce occurs when a receiving server could not or would not accept a message.
	Bounce *bool `json:"bounce,omitempty"`
	// Indicates if the webhook is configured to send deferred events. Deferred events occur when a recipient's email server temporarily rejects a message.
	Deferred *bool `json:"deferred,omitempty"`
	// Indicates if the webhook is configured to send unsubscribe events. Unsubscribes occur when recipients click on a message's subscription management link. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	Unsubscribe *bool `json:"unsubscribe,omitempty"`
	// Indicates if the webhook is configured to send processed events. Processed events occur when a message has been received by Twilio SendGrid and is ready to be delivered.
	Processed *bool `json:"processed,omitempty"`
	// Indicates if the webhook is configured to send open events. Open events occur when a recipient has opened the HTML message. You must [enable Open Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#open-tracking) to receive this type of event.
	Open *bool `json:"open,omitempty"`
	// Indicates if the webhook is configured to send click events. Click events occur when a recipient clicks on a link within the message. You must [enable Click Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#click-tracking) to receive this type of event.
	Click *bool `json:"click,omitempty"`
	// Indicates if the webhook is configured to send dropped events. Dropped events occur when your message is not delivered by Twilio SendGrid. Dropped events are accomponied by a `reason` property, which indicates why the message was dropped. Reasons for a dropped message include: Invalid SMTPAPI header, Spam Content (if spam checker app enabled), Unsubscribed Address, Bounced Address, Spam Reporting Address, Invalid, Recipient List over Package Quota.
	Dropped *bool `json:"dropped,omitempty"`
	// An optional friendly name assigned to the Event Webhook to help you differentiate it. The friendly name is for convenience only. You should use the webhook `id` property for any programmatic tasks.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A unique string used to identify the webhook. A webhook's ID is generated programmatically and cannot be changed after creation. You can assign a natural language identifier to your webhook using the `friendly_name` property.
	Id *string `json:"id,omitempty"`
	// An ISO 8601 timestamp in UTC timezone when the Event Webhook was created. If a Webhook's `created_date` is `null`, it is a [legacy Event Webook](https://www.twilio.com/en-us/changelog/event-webhooks), which means it is your oldest Webhook.
	CreatedDate *time.Time `json:"created_date,omitempty"`
	// An ISO 8601 timestamp in UTC timezone when the Event Webhook was last modified.
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	// The OAuth client ID SendGrid sends to your OAuth server or service provider to generate an OAuth access token.
	OauthClientId *string `json:"oauth_client_id,omitempty"`
	// The URL where SendGrid sends the OAuth client ID and client secret to generate an access token. This should be your OAuth server or service provider.
	OauthTokenUrl *string `json:"oauth_token_url,omitempty"`
	// The public key you can use to verify the SendGrid signature.
	PublicKey *string `json:"public_key,omitempty"`
}
