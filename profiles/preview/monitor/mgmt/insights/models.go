// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package insights

import original "github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2018-03-01/insights"

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Count   AggregationType = original.Count
	Maximum AggregationType = original.Maximum
	Minimum AggregationType = original.Minimum
	None    AggregationType = original.None
	Total   AggregationType = original.Total
)

func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}

type CategoryType = original.CategoryType

const (
	Logs    CategoryType = original.Logs
	Metrics CategoryType = original.Metrics
)

func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}

type ComparisonOperationType = original.ComparisonOperationType

const (
	Equals             ComparisonOperationType = original.Equals
	GreaterThan        ComparisonOperationType = original.GreaterThan
	GreaterThanOrEqual ComparisonOperationType = original.GreaterThanOrEqual
	LessThan           ComparisonOperationType = original.LessThan
	LessThanOrEqual    ComparisonOperationType = original.LessThanOrEqual
	NotEquals          ComparisonOperationType = original.NotEquals
)

func PossibleComparisonOperationTypeValues() []ComparisonOperationType {
	return original.PossibleComparisonOperationTypeValues()
}

type ConditionOperator = original.ConditionOperator

const (
	ConditionOperatorGreaterThan        ConditionOperator = original.ConditionOperatorGreaterThan
	ConditionOperatorGreaterThanOrEqual ConditionOperator = original.ConditionOperatorGreaterThanOrEqual
	ConditionOperatorLessThan           ConditionOperator = original.ConditionOperatorLessThan
	ConditionOperatorLessThanOrEqual    ConditionOperator = original.ConditionOperatorLessThanOrEqual
)

func PossibleConditionOperatorValues() []ConditionOperator {
	return original.PossibleConditionOperatorValues()
}

type EventLevel = original.EventLevel

const (
	Critical      EventLevel = original.Critical
	Error         EventLevel = original.Error
	Informational EventLevel = original.Informational
	Verbose       EventLevel = original.Verbose
	Warning       EventLevel = original.Warning
)

func PossibleEventLevelValues() []EventLevel {
	return original.PossibleEventLevelValues()
}

type MetricStatisticType = original.MetricStatisticType

const (
	MetricStatisticTypeAverage MetricStatisticType = original.MetricStatisticTypeAverage
	MetricStatisticTypeMax     MetricStatisticType = original.MetricStatisticTypeMax
	MetricStatisticTypeMin     MetricStatisticType = original.MetricStatisticTypeMin
	MetricStatisticTypeSum     MetricStatisticType = original.MetricStatisticTypeSum
)

func PossibleMetricStatisticTypeValues() []MetricStatisticType {
	return original.PossibleMetricStatisticTypeValues()
}

type OdataType = original.OdataType

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource          OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource
	OdataTypeRuleDataSource                                                      OdataType = original.OdataTypeRuleDataSource
)

func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}

type OdataTypeBasicRuleAction = original.OdataTypeBasicRuleAction

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction   OdataTypeBasicRuleAction = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction OdataTypeBasicRuleAction = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction
	OdataTypeRuleAction                                              OdataTypeBasicRuleAction = original.OdataTypeRuleAction
)

func PossibleOdataTypeBasicRuleActionValues() []OdataTypeBasicRuleAction {
	return original.PossibleOdataTypeBasicRuleActionValues()
}

type OdataTypeBasicRuleCondition = original.OdataTypeBasicRuleCondition

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition   OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition         OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition
	OdataTypeRuleCondition                                                        OdataTypeBasicRuleCondition = original.OdataTypeRuleCondition
)

func PossibleOdataTypeBasicRuleConditionValues() []OdataTypeBasicRuleCondition {
	return original.PossibleOdataTypeBasicRuleConditionValues()
}

type ReceiverStatus = original.ReceiverStatus

const (
	Disabled     ReceiverStatus = original.Disabled
	Enabled      ReceiverStatus = original.Enabled
	NotSpecified ReceiverStatus = original.NotSpecified
)

func PossibleReceiverStatusValues() []ReceiverStatus {
	return original.PossibleReceiverStatusValues()
}

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = original.RecurrenceFrequencyDay
	RecurrenceFrequencyHour   RecurrenceFrequency = original.RecurrenceFrequencyHour
	RecurrenceFrequencyMinute RecurrenceFrequency = original.RecurrenceFrequencyMinute
	RecurrenceFrequencyMonth  RecurrenceFrequency = original.RecurrenceFrequencyMonth
	RecurrenceFrequencyNone   RecurrenceFrequency = original.RecurrenceFrequencyNone
	RecurrenceFrequencySecond RecurrenceFrequency = original.RecurrenceFrequencySecond
	RecurrenceFrequencyWeek   RecurrenceFrequency = original.RecurrenceFrequencyWeek
	RecurrenceFrequencyYear   RecurrenceFrequency = original.RecurrenceFrequencyYear
)

func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return original.PossibleRecurrenceFrequencyValues()
}

type ResultType = original.ResultType

const (
	Data     ResultType = original.Data
	Metadata ResultType = original.Metadata
)

func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}

type ScaleDirection = original.ScaleDirection

const (
	ScaleDirectionDecrease ScaleDirection = original.ScaleDirectionDecrease
	ScaleDirectionIncrease ScaleDirection = original.ScaleDirectionIncrease
	ScaleDirectionNone     ScaleDirection = original.ScaleDirectionNone
)

func PossibleScaleDirectionValues() []ScaleDirection {
	return original.PossibleScaleDirectionValues()
}

type ScaleType = original.ScaleType

const (
	ChangeCount        ScaleType = original.ChangeCount
	ExactCount         ScaleType = original.ExactCount
	PercentChangeCount ScaleType = original.PercentChangeCount
)

func PossibleScaleTypeValues() []ScaleType {
	return original.PossibleScaleTypeValues()
}

type Sensitivity = original.Sensitivity

const (
	High   Sensitivity = original.High
	Low    Sensitivity = original.Low
	Medium Sensitivity = original.Medium
)

func PossibleSensitivityValues() []Sensitivity {
	return original.PossibleSensitivityValues()
}

type TimeAggregationOperator = original.TimeAggregationOperator

const (
	TimeAggregationOperatorAverage TimeAggregationOperator = original.TimeAggregationOperatorAverage
	TimeAggregationOperatorLast    TimeAggregationOperator = original.TimeAggregationOperatorLast
	TimeAggregationOperatorMaximum TimeAggregationOperator = original.TimeAggregationOperatorMaximum
	TimeAggregationOperatorMinimum TimeAggregationOperator = original.TimeAggregationOperatorMinimum
	TimeAggregationOperatorTotal   TimeAggregationOperator = original.TimeAggregationOperatorTotal
)

func PossibleTimeAggregationOperatorValues() []TimeAggregationOperator {
	return original.PossibleTimeAggregationOperatorValues()
}

type TimeAggregationType = original.TimeAggregationType

const (
	TimeAggregationTypeAverage TimeAggregationType = original.TimeAggregationTypeAverage
	TimeAggregationTypeCount   TimeAggregationType = original.TimeAggregationTypeCount
	TimeAggregationTypeMaximum TimeAggregationType = original.TimeAggregationTypeMaximum
	TimeAggregationTypeMinimum TimeAggregationType = original.TimeAggregationTypeMinimum
	TimeAggregationTypeTotal   TimeAggregationType = original.TimeAggregationTypeTotal
)

func PossibleTimeAggregationTypeValues() []TimeAggregationType {
	return original.PossibleTimeAggregationTypeValues()
}

type Unit = original.Unit

const (
	UnitBytes          Unit = original.UnitBytes
	UnitByteSeconds    Unit = original.UnitByteSeconds
	UnitBytesPerSecond Unit = original.UnitBytesPerSecond
	UnitCount          Unit = original.UnitCount
	UnitCountPerSecond Unit = original.UnitCountPerSecond
	UnitMilliSeconds   Unit = original.UnitMilliSeconds
	UnitPercent        Unit = original.UnitPercent
	UnitSeconds        Unit = original.UnitSeconds
	UnitUnspecified    Unit = original.UnitUnspecified
)

func PossibleUnitValues() []Unit {
	return original.PossibleUnitValues()
}

type ActionGroup = original.ActionGroup
type ActionGroupList = original.ActionGroupList
type ActionGroupPatch = original.ActionGroupPatch
type ActionGroupPatchBody = original.ActionGroupPatchBody
type ActionGroupResource = original.ActionGroupResource
type ActivityLogAlert = original.ActivityLogAlert
type ActivityLogAlertActionGroup = original.ActivityLogAlertActionGroup
type ActivityLogAlertActionList = original.ActivityLogAlertActionList
type ActivityLogAlertAllOfCondition = original.ActivityLogAlertAllOfCondition
type ActivityLogAlertLeafCondition = original.ActivityLogAlertLeafCondition
type ActivityLogAlertList = original.ActivityLogAlertList
type ActivityLogAlertPatch = original.ActivityLogAlertPatch
type ActivityLogAlertPatchBody = original.ActivityLogAlertPatchBody
type ActivityLogAlertResource = original.ActivityLogAlertResource
type AlertRule = original.AlertRule
type AlertRuleResource = original.AlertRuleResource
type AlertRuleResourceCollection = original.AlertRuleResourceCollection
type AlertRuleResourcePatch = original.AlertRuleResourcePatch
type AutomationRunbookReceiver = original.AutomationRunbookReceiver
type AutoscaleNotification = original.AutoscaleNotification
type AutoscaleProfile = original.AutoscaleProfile
type AutoscaleSetting = original.AutoscaleSetting
type AutoscaleSettingResource = original.AutoscaleSettingResource
type AutoscaleSettingResourceCollection = original.AutoscaleSettingResourceCollection
type AutoscaleSettingResourceCollectionIterator = original.AutoscaleSettingResourceCollectionIterator
type AutoscaleSettingResourceCollectionPage = original.AutoscaleSettingResourceCollectionPage
type AutoscaleSettingResourcePatch = original.AutoscaleSettingResourcePatch
type AzureAppPushReceiver = original.AzureAppPushReceiver
type Baseline = original.Baseline
type BaselineMetadataValue = original.BaselineMetadataValue
type BaselineProperties = original.BaselineProperties
type BaselineResponse = original.BaselineResponse
type CalculateBaselineResponse = original.CalculateBaselineResponse
type DiagnosticSettings = original.DiagnosticSettings
type DiagnosticSettingsCategory = original.DiagnosticSettingsCategory
type DiagnosticSettingsCategoryResource = original.DiagnosticSettingsCategoryResource
type DiagnosticSettingsCategoryResourceCollection = original.DiagnosticSettingsCategoryResourceCollection
type DiagnosticSettingsResource = original.DiagnosticSettingsResource
type DiagnosticSettingsResourceCollection = original.DiagnosticSettingsResourceCollection
type EmailNotification = original.EmailNotification
type EmailReceiver = original.EmailReceiver
type EnableRequest = original.EnableRequest
type ErrorResponse = original.ErrorResponse
type EventCategoryCollection = original.EventCategoryCollection
type EventData = original.EventData
type EventDataCollection = original.EventDataCollection
type EventDataCollectionIterator = original.EventDataCollectionIterator
type EventDataCollectionPage = original.EventDataCollectionPage
type HTTPRequestInfo = original.HTTPRequestInfo
type Incident = original.Incident
type IncidentListResult = original.IncidentListResult
type ItsmReceiver = original.ItsmReceiver
type LocalizableString = original.LocalizableString
type LocationThresholdRuleCondition = original.LocationThresholdRuleCondition
type LogicAppReceiver = original.LogicAppReceiver
type LogProfileCollection = original.LogProfileCollection
type LogProfileProperties = original.LogProfileProperties
type LogProfileResource = original.LogProfileResource
type LogProfileResourcePatch = original.LogProfileResourcePatch
type LogSettings = original.LogSettings
type ManagementEventAggregationCondition = original.ManagementEventAggregationCondition
type ManagementEventRuleCondition = original.ManagementEventRuleCondition
type MetadataValue = original.MetadataValue
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionCollection = original.MetricDefinitionCollection
type MetricSettings = original.MetricSettings
type MetricTrigger = original.MetricTrigger
type MetricValue = original.MetricValue
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type ProxyOnlyResource = original.ProxyOnlyResource
type Recurrence = original.Recurrence
type RecurrentSchedule = original.RecurrentSchedule
type Resource = original.Resource
type Response = original.Response
type RetentionPolicy = original.RetentionPolicy
type BasicRuleAction = original.BasicRuleAction
type RuleAction = original.RuleAction
type BasicRuleCondition = original.BasicRuleCondition
type RuleCondition = original.RuleCondition
type BasicRuleDataSource = original.BasicRuleDataSource
type RuleDataSource = original.RuleDataSource
type RuleEmailAction = original.RuleEmailAction
type RuleManagementEventClaimsDataSource = original.RuleManagementEventClaimsDataSource
type RuleManagementEventDataSource = original.RuleManagementEventDataSource
type RuleMetricDataSource = original.RuleMetricDataSource
type RuleWebhookAction = original.RuleWebhookAction
type ScaleAction = original.ScaleAction
type ScaleCapacity = original.ScaleCapacity
type ScaleRule = original.ScaleRule
type SenderAuthorization = original.SenderAuthorization
type SmsReceiver = original.SmsReceiver
type ThresholdRuleCondition = original.ThresholdRuleCondition
type TimeSeriesElement = original.TimeSeriesElement
type TimeSeriesInformation = original.TimeSeriesInformation
type TimeWindow = original.TimeWindow
type VoiceReceiver = original.VoiceReceiver
type WebhookNotification = original.WebhookNotification
type WebhookReceiver = original.WebhookReceiver

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type AlertRulesClient = original.AlertRulesClient

func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}

type MetricDefinitionsClient = original.MetricDefinitionsClient

func NewMetricDefinitionsClient(subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClient(subscriptionID)
}
func NewMetricDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}

type MetricsClient = original.MetricsClient

func NewMetricsClient(subscriptionID string) MetricsClient {
	return original.NewMetricsClient(subscriptionID)
}
func NewMetricsClientWithBaseURI(baseURI string, subscriptionID string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI, subscriptionID)
}

type ActionGroupsClient = original.ActionGroupsClient

func NewActionGroupsClient(subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClient(subscriptionID)
}
func NewActionGroupsClientWithBaseURI(baseURI string, subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClientWithBaseURI(baseURI, subscriptionID)
}

type LogProfilesClient = original.LogProfilesClient

func NewLogProfilesClient(subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClient(subscriptionID)
}
func NewLogProfilesClientWithBaseURI(baseURI string, subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type EventCategoriesClient = original.EventCategoriesClient

func NewEventCategoriesClient(subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClient(subscriptionID)
}
func NewEventCategoriesClientWithBaseURI(baseURI string, subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClientWithBaseURI(baseURI, subscriptionID)
}

type AlertRuleIncidentsClient = original.AlertRuleIncidentsClient

func NewAlertRuleIncidentsClient(subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClient(subscriptionID)
}
func NewAlertRuleIncidentsClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClientWithBaseURI(baseURI, subscriptionID)
}

type MetricBaselineClient = original.MetricBaselineClient

func NewMetricBaselineClient(subscriptionID string) MetricBaselineClient {
	return original.NewMetricBaselineClient(subscriptionID)
}
func NewMetricBaselineClientWithBaseURI(baseURI string, subscriptionID string) MetricBaselineClient {
	return original.NewMetricBaselineClientWithBaseURI(baseURI, subscriptionID)
}

type DiagnosticSettingsCategoryClient = original.DiagnosticSettingsCategoryClient

func NewDiagnosticSettingsCategoryClient(subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClient(subscriptionID)
}
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI, subscriptionID)
}

type TenantActivityLogsClient = original.TenantActivityLogsClient

func NewTenantActivityLogsClient(subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClient(subscriptionID)
}
func NewTenantActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}

type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type ActivityLogAlertsClient = original.ActivityLogAlertsClient

func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClient(subscriptionID)
}
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClientWithBaseURI(baseURI, subscriptionID)
}

type ActivityLogsClient = original.ActivityLogsClient

func NewActivityLogsClient(subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClient(subscriptionID)
}
func NewActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}

type DiagnosticSettingsClient = original.DiagnosticSettingsClient

func NewDiagnosticSettingsClient(subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClient(subscriptionID)
}
func NewDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClientWithBaseURI(baseURI, subscriptionID)
}

type AutoscaleSettingsClient = original.AutoscaleSettingsClient

func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClient(subscriptionID)
}
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClientWithBaseURI(baseURI, subscriptionID)
}
