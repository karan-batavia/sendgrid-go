/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Address Management API
* The Twilio SendGrid IP Address Management API combines functionality that was previously split between the Twilio SendGrid [IP Address API](https://docs.sendgrid.com/api-reference/ip-address) and [IP Pools API](https://docs.sendgrid.com/api-reference/ip-pools). This functionality includes adding IP addresses to your account, assigning IP addresses to IP Pools and Subusers, among other tasks.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ListSubUserAssignedToIp200Response struct for ListSubUserAssignedToIp200Response
type ListSubUserAssignedToIp200Response struct {
	// An array of Subuser IDs that have been assigned the specified IP address.
	Result   *[]string                                   `json:"result,omitempty"`
	Metadata *ListSubUserAssignedToIp200ResponseMetadata `json:"_metadata,omitempty"`
}
