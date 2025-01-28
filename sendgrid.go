/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */
// Package sendgrid provides bindings for Sendgrid's REST APIs.
package sendgrid

import (
	"errors"
	"net/url"
	"os"
	"time"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/v4/client"
	AccountProvisioningV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/account_provisioning"
	AlertsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/alerts"
	ApiKeysV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/api_keys"
	DomainAuthenticationV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/domain_authentication"
	EmailActivityV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/email_activity"
	EmailValidationV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/email_validation"
	EnforcedTlsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/enforced_tls"
	IntegrationsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/integrations"
	IpAccessManagementV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/ip_access_management"
	IpAddressManagementV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/ip_address_management"
	IpWarmupV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/ip_warmup"
	IpsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/ips"
	LinkBrandingV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/link_branding"
	LmcCampaignsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/lmc_campaigns"
	LmcContactdbV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/lmc_contactdb"
	LmcSendersV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/lmc_senders"
	MailV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mail"
	MailSettingsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mail_settings"
	McContactsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_contacts"
	McCustomFieldsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_custom_fields"
	McDesignsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_designs"
	McListsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_lists"
	McSegmentsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_segments"
	McSegments2V3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_segments_2"
	McSendersV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_senders"
	McSinglesendsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_singlesends"
	McStatsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_stats"
	McTestV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/mc_test"
	PartnerV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/partner"
	RecipientsDataErasureV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/recipients_data_erasure"
	ReverseDnsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/reverse_dns"
	ScheduledSendsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/scheduled_sends"
	ScopesV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/scopes"
	SeqV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/seq"
	SsoV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/sso"
	StatsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/stats"
	SubusersV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/subusers"
	SuppressionsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/suppressions"
	TemplatesV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/templates"
	TrackingSettingsV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/tracking_settings"
	UserV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/user"
	VerifiedSendersV3 "github.com/sendgrid/sendgrid-go/v4/rest/api/v3/verified_senders"
)

// RestClient provides access to Sendgrid services.
type RestClient struct {
	*client.RequestHandler
	AccountProvisioningV3   *AccountProvisioningV3.ApiService
	AlertsV3                *AlertsV3.ApiService
	ApiKeysV3               *ApiKeysV3.ApiService
	DomainAuthenticationV3  *DomainAuthenticationV3.ApiService
	EmailActivityV3         *EmailActivityV3.ApiService
	EmailValidationV3       *EmailValidationV3.ApiService
	EnforcedTlsV3           *EnforcedTlsV3.ApiService
	IntegrationsV3          *IntegrationsV3.ApiService
	IpAccessManagementV3    *IpAccessManagementV3.ApiService
	IpAddressManagementV3   *IpAddressManagementV3.ApiService
	IpWarmupV3              *IpWarmupV3.ApiService
	IpsV3                   *IpsV3.ApiService
	LinkBrandingV3          *LinkBrandingV3.ApiService
	LmcCampaignsV3          *LmcCampaignsV3.ApiService
	LmcContactdbV3          *LmcContactdbV3.ApiService
	LmcSendersV3            *LmcSendersV3.ApiService
	MailV3                  *MailV3.ApiService
	MailSettingsV3          *MailSettingsV3.ApiService
	McContactsV3            *McContactsV3.ApiService
	McCustomFieldsV3        *McCustomFieldsV3.ApiService
	McDesignsV3             *McDesignsV3.ApiService
	McListsV3               *McListsV3.ApiService
	McSegmentsV3            *McSegmentsV3.ApiService
	McSegments2V3           *McSegments2V3.ApiService
	McSendersV3             *McSendersV3.ApiService
	McSinglesendsV3         *McSinglesendsV3.ApiService
	McStatsV3               *McStatsV3.ApiService
	McTestV3                *McTestV3.ApiService
	PartnerV3               *PartnerV3.ApiService
	RecipientsDataErasureV3 *RecipientsDataErasureV3.ApiService
	ReverseDnsV3            *ReverseDnsV3.ApiService
	ScheduledSendsV3        *ScheduledSendsV3.ApiService
	ScopesV3                *ScopesV3.ApiService
	SeqV3                   *SeqV3.ApiService
	SsoV3                   *SsoV3.ApiService
	StatsV3                 *StatsV3.ApiService
	SubusersV3              *SubusersV3.ApiService
	SuppressionsV3          *SuppressionsV3.ApiService
	TemplatesV3             *TemplatesV3.ApiService
	TrackingSettingsV3      *TrackingSettingsV3.ApiService
	UserV3                  *UserV3.ApiService
	VerifiedSendersV3       *VerifiedSendersV3.ApiService
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

// sendGridOptions for CreateRequest
type sendGridOptions struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

// sendgrid host map for different regions
var allowedRegionsHostMap = map[string]string{
	"eu":     "https://api.eu.sendgrid.com",
	"global": "https://api.sendgrid.com",
}

// map for allowed regions, global and eu currently
var allowedRegions = map[string]bool{
	"eu":     true,
	"global": true,
}

// GetRequest
// @return [Request] a default request object
func GetRequest(key, endpoint, host string) rest.Request {
	return createSendGridRequest(sendGridOptions{key, endpoint, host, ""})
}

// GetRequestSubuser like GetRequest but with On-Behalf of Subuser
// @return [Request] a default request object
func GetRequestSubuser(key, endpoint, host, subuser string) rest.Request {
	return createSendGridRequest(sendGridOptions{key, endpoint, host, subuser})
}

// createSendGridRequest create Request
// @return [Request] a default request object
func createSendGridRequest(sgOptions sendGridOptions) rest.Request {
	options := options{
		"Bearer " + sgOptions.Key,
		sgOptions.Endpoint,
		sgOptions.Host,
		sgOptions.Subuser,
	}

	if options.Host == "" {
		options.Host = "https://api.sendgrid.com"
	}

	return requestNew(options)
}

type ClientParams struct {
	ApiKey string
	Client client.BaseClient
}

// NewRestClientWithParams provides an initialized Sendgrid RestClient with params.
func NewRestClientWithParams(params ClientParams) *RestClient {
	requestHandler := client.NewRequestHandler(params.Client)

	if params.Client == nil {
		apiKey := params.ApiKey
		if apiKey == "" {
			apiKey = os.Getenv("SENDGRID_API_KEY")
		}

		defaultClient := &client.Client{
			Credentials: client.NewCredentials(apiKey),
		}

		requestHandler = client.NewRequestHandler(defaultClient)
	}

	c := &RestClient{
		RequestHandler: requestHandler,
	}

	c.AccountProvisioningV3 = AccountProvisioningV3.NewApiService(c.RequestHandler)
	c.AlertsV3 = AlertsV3.NewApiService(c.RequestHandler)
	c.ApiKeysV3 = ApiKeysV3.NewApiService(c.RequestHandler)
	c.DomainAuthenticationV3 = DomainAuthenticationV3.NewApiService(c.RequestHandler)
	c.EmailActivityV3 = EmailActivityV3.NewApiService(c.RequestHandler)
	c.EmailValidationV3 = EmailValidationV3.NewApiService(c.RequestHandler)
	c.EnforcedTlsV3 = EnforcedTlsV3.NewApiService(c.RequestHandler)
	c.IntegrationsV3 = IntegrationsV3.NewApiService(c.RequestHandler)
	c.IpAccessManagementV3 = IpAccessManagementV3.NewApiService(c.RequestHandler)
	c.IpAddressManagementV3 = IpAddressManagementV3.NewApiService(c.RequestHandler)
	c.IpWarmupV3 = IpWarmupV3.NewApiService(c.RequestHandler)
	c.IpsV3 = IpsV3.NewApiService(c.RequestHandler)
	c.LinkBrandingV3 = LinkBrandingV3.NewApiService(c.RequestHandler)
	c.LmcCampaignsV3 = LmcCampaignsV3.NewApiService(c.RequestHandler)
	c.LmcContactdbV3 = LmcContactdbV3.NewApiService(c.RequestHandler)
	c.LmcSendersV3 = LmcSendersV3.NewApiService(c.RequestHandler)
	c.MailV3 = MailV3.NewApiService(c.RequestHandler)
	c.MailSettingsV3 = MailSettingsV3.NewApiService(c.RequestHandler)
	c.McContactsV3 = McContactsV3.NewApiService(c.RequestHandler)
	c.McCustomFieldsV3 = McCustomFieldsV3.NewApiService(c.RequestHandler)
	c.McDesignsV3 = McDesignsV3.NewApiService(c.RequestHandler)
	c.McListsV3 = McListsV3.NewApiService(c.RequestHandler)
	c.McSegmentsV3 = McSegmentsV3.NewApiService(c.RequestHandler)
	c.McSegments2V3 = McSegments2V3.NewApiService(c.RequestHandler)
	c.McSendersV3 = McSendersV3.NewApiService(c.RequestHandler)
	c.McSinglesendsV3 = McSinglesendsV3.NewApiService(c.RequestHandler)
	c.McStatsV3 = McStatsV3.NewApiService(c.RequestHandler)
	c.McTestV3 = McTestV3.NewApiService(c.RequestHandler)
	c.PartnerV3 = PartnerV3.NewApiService(c.RequestHandler)
	c.RecipientsDataErasureV3 = RecipientsDataErasureV3.NewApiService(c.RequestHandler)
	c.ReverseDnsV3 = ReverseDnsV3.NewApiService(c.RequestHandler)
	c.ScheduledSendsV3 = ScheduledSendsV3.NewApiService(c.RequestHandler)
	c.ScopesV3 = ScopesV3.NewApiService(c.RequestHandler)
	c.SeqV3 = SeqV3.NewApiService(c.RequestHandler)
	c.SsoV3 = SsoV3.NewApiService(c.RequestHandler)
	c.StatsV3 = StatsV3.NewApiService(c.RequestHandler)
	c.SubusersV3 = SubusersV3.NewApiService(c.RequestHandler)
	c.SuppressionsV3 = SuppressionsV3.NewApiService(c.RequestHandler)
	c.TemplatesV3 = TemplatesV3.NewApiService(c.RequestHandler)
	c.TrackingSettingsV3 = TrackingSettingsV3.NewApiService(c.RequestHandler)
	c.UserV3 = UserV3.NewApiService(c.RequestHandler)
	c.VerifiedSendersV3 = VerifiedSendersV3.NewApiService(c.RequestHandler)

	return c
}

// NewRestClient provides an initialized Sendgrid RestClient.
func NewRestClient() *RestClient {
	return NewRestClientWithParams(ClientParams{})
}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(key string) *Client {
	request := GetRequest(key, "/v3/mail/send", "")
	request.Method = "POST"
	return &Client{request}
}

// extractEndpoint extracts the endpoint from a baseURL
func extractEndpoint(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	return parsedURL.Path, nil
}

// SetTimeout sets the Timeout for Sendgrid HTTP requests.
func (c *RestClient) SetTimeout(timeout time.Duration) {
	c.RequestHandler.Client.SetTimeout(timeout)
}

// SetEdge sets the Edge for the Sendgrid request.
// Not supported in sendgrid currently
func (c *RestClient) SetEdge(edge string) {
	c.RequestHandler.Edge = edge
}

// SetRegion sets the Region for the Sendgrid request. Defaults to "us1" if an edge is provided.
func (c *RestClient) SetRegion(region string) {
	c.RequestHandler.Region = region
}

// SetDataResidency modifies the host as per the region
/*
* This allows support for global and eu regions only. This set will likely expand in the future.
* Global should be the default
* Global region means the message should be sent through:
* HTTP: api.sendgrid.com
* EU region means the message should be sent through:
* HTTP: api.eu.sendgrid.com
 */
// @return [Request] the modified request object
func SetDataResidency(request rest.Request, region string) (rest.Request, error) {
	regionalHost, present := allowedRegionsHostMap[region]
	if !present {
		return request, errors.New("error: region can only be \"eu\" or \"global\"")
	}
	endpoint, err := extractEndpoint(request.BaseURL)
	if err != nil {
		return request, err
	}
	request.BaseURL = regionalHost + endpoint
	return request, nil
}
