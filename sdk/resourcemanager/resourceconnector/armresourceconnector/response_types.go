//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourceconnector

// AppliancesClientCreateOrUpdateResponse contains the response from method AppliancesClient.CreateOrUpdate.
type AppliancesClientCreateOrUpdateResponse struct {
	Appliance
}

// AppliancesClientDeleteResponse contains the response from method AppliancesClient.Delete.
type AppliancesClientDeleteResponse struct {
	// placeholder for future response values
}

// AppliancesClientGetResponse contains the response from method AppliancesClient.Get.
type AppliancesClientGetResponse struct {
	Appliance
}

// AppliancesClientGetUpgradeGraphResponse contains the response from method AppliancesClient.GetUpgradeGraph.
type AppliancesClientGetUpgradeGraphResponse struct {
	UpgradeGraph
}

// AppliancesClientListByResourceGroupResponse contains the response from method AppliancesClient.ListByResourceGroup.
type AppliancesClientListByResourceGroupResponse struct {
	ApplianceListResult
}

// AppliancesClientListBySubscriptionResponse contains the response from method AppliancesClient.ListBySubscription.
type AppliancesClientListBySubscriptionResponse struct {
	ApplianceListResult
}

// AppliancesClientListClusterCustomerUserCredentialResponse contains the response from method AppliancesClient.ListClusterCustomerUserCredential.
type AppliancesClientListClusterCustomerUserCredentialResponse struct {
	ApplianceListClusterCustomerUserCredentialResults
}

// AppliancesClientListClusterUserCredentialResponse contains the response from method AppliancesClient.ListClusterUserCredential.
type AppliancesClientListClusterUserCredentialResponse struct {
	ApplianceListCredentialResults
}

// AppliancesClientListOperationsResponse contains the response from method AppliancesClient.ListOperations.
type AppliancesClientListOperationsResponse struct {
	ApplianceOperationsList
}

// AppliancesClientUpdateResponse contains the response from method AppliancesClient.Update.
type AppliancesClientUpdateResponse struct {
	Appliance
}
