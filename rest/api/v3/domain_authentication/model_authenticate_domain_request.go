/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Domain Authentication API
* The Twilio SendGrid Domain Authentication API allows you to manage your authenticated domains and their settings.  Domain Authentication is a required step when setting up your Twilio SendGrid account because it's essential to ensuring the deliverability of your email. Domain Authentication signals trustworthiness to email inbox providers and your recipients by approving SendGrid to send email on behalf of your domain. For more information, see [**How to Set Up Domain Authentication**](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-domain-authentication/).  Each user may have a maximum of 3,000 authenticated domains and 3,000 link brandings. This limit is at the user level, meaning each Subuser belonging to a parent account may have its own 3,000 authenticated domains and 3,000 link brandings.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// AuthenticateDomainRequest struct for AuthenticateDomainRequest
type AuthenticateDomainRequest struct {
	// Domain being authenticated.
	Domain string `json:"domain"`
	// The subdomain to use for this authenticated domain.
	Subdomain *string `json:"subdomain,omitempty"`
	// The username associated with this domain.
	Username *string `json:"username,omitempty"`
	// The IP addresses that will be included in the custom SPF record for this authenticated domain.
	Ips *[]string `json:"ips,omitempty"`
	// Specify whether to use a custom SPF or allow SendGrid to manage your SPF. This option is only available to authenticated domains set up for manual security.
	CustomSpf *bool `json:"custom_spf,omitempty"`
	// Whether to use this authenticated domain as the fallback if no authenticated domains match the sender's domain.
	Default *bool `json:"default,omitempty"`
	// Whether to allow SendGrid to manage your SPF records, DKIM keys, and DKIM key rotation.
	AutomaticSecurity *bool `json:"automatic_security,omitempty"`
	// Add a custom DKIM selector. Accepts three letters or numbers.
	CustomDkimSelector *string `json:"custom_dkim_selector,omitempty"`
	// The region of the domain. Allowed values are `global` and `eu`. The default value is `global`.
	Region *string `json:"region,omitempty"`
}
