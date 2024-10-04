/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Segments API
* This API was deprecated on December 31, 2022. Following deprecation, all segments created in the Marketing Campaigns user interface began using the [Segmentation v2 API](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2).  To enable manual migration and data retrieval, this API's GET and DELETE operations will remain available. The POST (create) and PATCH (update) endpoints were removed on January 31, 2023 because it is no longer possible to create new v1 segments or modify existing ones. See our [Segmentation v1 to v2 upgrade instructions](https://docs.sendgrid.com/for-developers/sending-email/getting-started-the-marketing-campaigns-v2-segmentation-api#upgrade-a-v1-segment-to-v2) to manually migrate your segments to the v2 API.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ContactResponse struct for ContactResponse
type ContactResponse struct {
	// ID assigned to a contact when added to the system.
	Id string `json:"id"`
	// Email of the contact. This is a reserved field.
	Email *string `json:"email,omitempty"`
	// The contact's Phone Number ID. This must be a valid phone number.
	PhoneNumberId *string `json:"phone_number_id,omitempty"`
	// The contact's External ID.
	ExternalId *string `json:"external_id,omitempty"`
	// The contact's Anonymous ID.
	AnonymousId *string `json:"anonymous_id,omitempty"`
	// Alternate emails of the contact. This is a reserved field.
	AlternateEmails []string `json:"alternate_emails"`
	// First name of the contact. This is a reserved field.
	FirstName string `json:"first_name"`
	// Last name of the contact. This is a reserved field.
	LastName string `json:"last_name"`
	// First line of address of the contact. This is a reserved field.
	AddressLine1 string `json:"address_line_1"`
	// Second line of address of the contact. This is a reserved field.
	AddressLine2 string `json:"address_line_2"`
	// City associated with the contact. This is a reserved field.
	City string `json:"city"`
	// State associated with the contact. This is a reserved field.
	StateProvinceRegion string `json:"state_province_region"`
	// Zipcode associated with the address of the contact. This is a reserved field.
	PostalCode int32 `json:"postal_code"`
	// Country associated with the address of the contact. This is a reserved field.
	Country string `json:"country"`
	// IDs of all lists the contact is part of
	ListIds      *[]string                   `json:"list_ids,omitempty"`
	CustomFields ContactResponseCustomFields `json:"custom_fields"`
	// IDs of all segments the contact is part of
	SegmentIds *[]string `json:"segment_ids,omitempty"`
}
