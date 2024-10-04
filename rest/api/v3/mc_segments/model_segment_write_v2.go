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

// SegmentWriteV2 struct for SegmentWriteV2
type SegmentWriteV2 struct {
	// Name of the segment.
	Name string `json:"name"`
	// The array of list ids to filter contacts on when building this segment. It allows only one such list id for now. We will support more in future
	ParentListIds *[]string `json:"parent_list_ids,omitempty"`
	// SQL query which will filter contacts based on the conditions provided
	QueryDsl string `json:"query_dsl"`
}
