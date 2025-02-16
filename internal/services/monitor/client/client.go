package client

import (
	"github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2017-04-01/aad"                                           // nolint: staticcheck
	"github.com/Azure/azure-sdk-for-go/services/preview/alertsmanagement/mgmt/2019-06-01-preview/alertsmanagement" // nolint: staticcheck
	classic "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"          // nolint: staticcheck
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2021-08-08/alertprocessingrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalerts"
	scheduledqueryrules2018 "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-04-16/scheduledqueryrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privatelinkscopedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privatelinkscopesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2020-10-01/activitylogalertsapis"
	diagnosticSettingClient "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettings"
	diagnosticCategoryClient "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettingscategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-08-01/scheduledqueryrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionruleassociations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionrules"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	// AAD
	AADDiagnosticSettingsClient *aad.DiagnosticSettingsClient

	// Autoscale Settings
	AutoscaleSettingsClient *autoscalesettings.AutoScaleSettingsClient

	// alerts management
	ActionRulesClient             *alertsmanagement.ActionRulesClient
	AlertProcessingRulesClient    *alertprocessingrules.AlertProcessingRulesClient
	SmartDetectorAlertRulesClient *alertsmanagement.SmartDetectorAlertRulesClient

	// Monitor
	ActionGroupsClient                   *actiongroupsapis.ActionGroupsAPIsClient
	ActivityLogsClient                   *classic.ActivityLogsClient
	ActivityLogAlertsClient              *activitylogalertsapis.ActivityLogAlertsAPIsClient
	AlertRulesClient                     *classic.AlertRulesClient
	DataCollectionEndpointsClient        *datacollectionendpoints.DataCollectionEndpointsClient
	DataCollectionRuleAssociationsClient *datacollectionruleassociations.DataCollectionRuleAssociationsClient
	DataCollectionRulesClient            *datacollectionrules.DataCollectionRulesClient
	DiagnosticSettingsClient             *diagnosticSettingClient.DiagnosticSettingsClient
	DiagnosticSettingsCategoryClient     *diagnosticCategoryClient.DiagnosticSettingsCategoriesClient
	LogProfilesClient                    *logprofiles.LogProfilesClient
	MetricAlertsClient                   *metricalerts.MetricAlertsClient
	PrivateLinkScopesClient              *privatelinkscopesapis.PrivateLinkScopesAPIsClient
	PrivateLinkScopedResourcesClient     *privatelinkscopedresources.PrivateLinkScopedResourcesClient
	ScheduledQueryRulesClient            *scheduledqueryrules2018.ScheduledQueryRulesClient
	ScheduledQueryRulesV2Client          *scheduledqueryrules.ScheduledQueryRulesClient
}

func NewClient(o *common.ClientOptions) *Client {
	AADDiagnosticSettingsClient := aad.NewDiagnosticSettingsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AADDiagnosticSettingsClient.Client, o.ResourceManagerAuthorizer)

	AutoscaleSettingsClient := autoscalesettings.NewAutoScaleSettingsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AutoscaleSettingsClient.Client, o.ResourceManagerAuthorizer)

	ActionRulesClient := alertsmanagement.NewActionRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ActionRulesClient.Client, o.ResourceManagerAuthorizer)

	AlertProcessingRulesClient := alertprocessingrules.NewAlertProcessingRulesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AlertProcessingRulesClient.Client, o.ResourceManagerAuthorizer)

	SmartDetectorAlertRulesClient := alertsmanagement.NewSmartDetectorAlertRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&SmartDetectorAlertRulesClient.Client, o.ResourceManagerAuthorizer)

	ActionGroupsClient := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ActionGroupsClient.Client, o.ResourceManagerAuthorizer)

	activityLogsClient := classic.NewActivityLogsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&activityLogsClient.Client, o.ResourceManagerAuthorizer)

	ActivityLogAlertsClient := activitylogalertsapis.NewActivityLogAlertsAPIsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ActivityLogAlertsClient.Client, o.ResourceManagerAuthorizer)

	AlertRulesClient := classic.NewAlertRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&AlertRulesClient.Client, o.ResourceManagerAuthorizer)

	DataCollectionEndpointsClient := datacollectionendpoints.NewDataCollectionEndpointsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DataCollectionEndpointsClient.Client, o.ResourceManagerAuthorizer)

	DataCollectionRuleAssociationsClient := datacollectionruleassociations.NewDataCollectionRuleAssociationsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DataCollectionRuleAssociationsClient.Client, o.ResourceManagerAuthorizer)

	DataCollectionRulesClient := datacollectionrules.NewDataCollectionRulesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DataCollectionRulesClient.Client, o.ResourceManagerAuthorizer)

	DiagnosticSettingsClient := diagnosticSettingClient.NewDiagnosticSettingsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DiagnosticSettingsClient.Client, o.ResourceManagerAuthorizer)

	DiagnosticSettingsCategoryClient := diagnosticCategoryClient.NewDiagnosticSettingsCategoriesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DiagnosticSettingsCategoryClient.Client, o.ResourceManagerAuthorizer)

	LogProfilesClient := logprofiles.NewLogProfilesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&LogProfilesClient.Client, o.ResourceManagerAuthorizer)

	MetricAlertsClient := metricalerts.NewMetricAlertsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&MetricAlertsClient.Client, o.ResourceManagerAuthorizer)

	PrivateLinkScopesClient := privatelinkscopesapis.NewPrivateLinkScopesAPIsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&PrivateLinkScopesClient.Client, o.ResourceManagerAuthorizer)

	PrivateLinkScopedResourcesClient := privatelinkscopedresources.NewPrivateLinkScopedResourcesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&PrivateLinkScopedResourcesClient.Client, o.ResourceManagerAuthorizer)

	ScheduledQueryRulesClient := scheduledqueryrules2018.NewScheduledQueryRulesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ScheduledQueryRulesClient.Client, o.ResourceManagerAuthorizer)

	ScheduledQueryRulesV2Client := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ScheduledQueryRulesV2Client.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AADDiagnosticSettingsClient:          &AADDiagnosticSettingsClient,
		AutoscaleSettingsClient:              &AutoscaleSettingsClient,
		ActionRulesClient:                    &ActionRulesClient,
		SmartDetectorAlertRulesClient:        &SmartDetectorAlertRulesClient,
		ActionGroupsClient:                   &ActionGroupsClient,
		ActivityLogsClient:                   &activityLogsClient,
		ActivityLogAlertsClient:              &ActivityLogAlertsClient,
		AlertRulesClient:                     &AlertRulesClient,
		AlertProcessingRulesClient:           &AlertProcessingRulesClient,
		DataCollectionEndpointsClient:        &DataCollectionEndpointsClient,
		DataCollectionRuleAssociationsClient: &DataCollectionRuleAssociationsClient,
		DataCollectionRulesClient:            &DataCollectionRulesClient,
		DiagnosticSettingsClient:             &DiagnosticSettingsClient,
		DiagnosticSettingsCategoryClient:     &DiagnosticSettingsCategoryClient,
		LogProfilesClient:                    &LogProfilesClient,
		MetricAlertsClient:                   &MetricAlertsClient,
		PrivateLinkScopesClient:              &PrivateLinkScopesClient,
		PrivateLinkScopedResourcesClient:     &PrivateLinkScopedResourcesClient,
		ScheduledQueryRulesClient:            &ScheduledQueryRulesClient,
		ScheduledQueryRulesV2Client:          &ScheduledQueryRulesV2Client,
	}
}
