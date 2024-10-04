/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Address API
* The Twilio SendGrid IP Address API allows you to add and manage dedicated IP addresses and IP Pools for your SendGrid account and Subusers. If you are sending any significant amount of email, SendGrid typically suggests sending from dedicated IPs. It's also best to send marketing and transactional emails from separate IP addresses. In order to do this, you'll need to set up IP Pools, which are groups of dedicated IP addresses you define to send particular types of messages. See the [**Dedicated IP Addresses**](https://docs.sendgrid.com/ui/account-and-settings/dedicated-ip-addresses) for more information about obtaining and allocating IPs.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// AddIpToIpPool201Response struct for AddIpToIpPool201Response
type AddIpToIpPool201Response struct {
	// The IP address.
	Ip string `json:"ip"`
	// The IP pools that this IP address has been added to.
	Pools []string `json:"pools"`
	// A Unix timestamp indicating when the warmup process began for the added IP address.
	StartDate int32 `json:"start_date"`
	// Indicates if the IP address is in warmup.
	Warmup bool `json:"warmup"`
}
