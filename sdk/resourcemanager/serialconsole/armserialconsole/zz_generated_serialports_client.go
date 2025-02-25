//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SerialPortsClient contains the methods for the SerialPorts group.
// Don't use this type directly, use NewSerialPortsClient() instead.
type SerialPortsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSerialPortsClient creates a new instance of SerialPortsClient with the specified values.
// subscriptionID - Subscription ID which uniquely identifies the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call requiring it.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSerialPortsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SerialPortsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SerialPortsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Connect - Connect to serial port of the target resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// resourceGroupName - The name of the resource group.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourceType - The resource type of the parent resource. For example: 'virtualMachines' or 'virtualMachineScaleSets'
// parentResource - The resource name, or subordinate path, for the parent of the serial port. For example: the name of the
// virtual machine.
// serialPort - The name of the serial port to connect to.
// options - SerialPortsClientConnectOptions contains the optional parameters for the SerialPortsClient.Connect method.
func (client *SerialPortsClient) Connect(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientConnectOptions) (SerialPortsClientConnectResponse, error) {
	req, err := client.connectCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourceType, parentResource, serialPort, options)
	if err != nil {
		return SerialPortsClientConnectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientConnectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SerialPortsClientConnectResponse{}, runtime.NewResponseError(resp)
	}
	return client.connectHandleResponse(resp)
}

// connectCreateRequest creates the Connect request.
func (client *SerialPortsClient) connectCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientConnectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourceType}/{parentResource}/providers/Microsoft.SerialConsole/serialPorts/{serialPort}/connect"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", parentResourceType)
	if parentResource == "" {
		return nil, errors.New("parameter parentResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResource}", url.PathEscape(parentResource))
	if serialPort == "" {
		return nil, errors.New("parameter serialPort cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialPort}", url.PathEscape(serialPort))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// connectHandleResponse handles the Connect response.
func (client *SerialPortsClient) connectHandleResponse(resp *http.Response) (SerialPortsClientConnectResponse, error) {
	result := SerialPortsClientConnectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SerialPortConnectResult); err != nil {
		return SerialPortsClientConnectResponse{}, err
	}
	return result, nil
}

// Create - Creates or updates a serial port
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// resourceGroupName - The name of the resource group.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourceType - The resource type of the parent resource. For example: 'virtualMachines' or 'virtualMachineScaleSets'
// parentResource - The resource name, or subordinate path, for the parent of the serial port. For example: the name of the
// virtual machine.
// serialPort - The name of the serial port to create.
// parameters - Parameters supplied to create the serial port.
// options - SerialPortsClientCreateOptions contains the optional parameters for the SerialPortsClient.Create method.
func (client *SerialPortsClient) Create(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, parameters SerialPort, options *SerialPortsClientCreateOptions) (SerialPortsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourceType, parentResource, serialPort, parameters, options)
	if err != nil {
		return SerialPortsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return SerialPortsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SerialPortsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, parameters SerialPort, options *SerialPortsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourceType}/{parentResource}/providers/Microsoft.SerialConsole/serialPorts/{serialPort}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", parentResourceType)
	if parentResource == "" {
		return nil, errors.New("parameter parentResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResource}", url.PathEscape(parentResource))
	if serialPort == "" {
		return nil, errors.New("parameter serialPort cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialPort}", url.PathEscape(serialPort))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *SerialPortsClient) createHandleResponse(resp *http.Response) (SerialPortsClientCreateResponse, error) {
	result := SerialPortsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SerialPort); err != nil {
		return SerialPortsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a serial port
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// resourceGroupName - The name of the resource group.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourceType - The resource type of the parent resource. For example: 'virtualMachines' or 'virtualMachineScaleSets'
// parentResource - The resource name, or subordinate path, for the parent of the serial port. For example: the name of the
// virtual machine.
// serialPort - The name of the serial port to delete.
// options - SerialPortsClientDeleteOptions contains the optional parameters for the SerialPortsClient.Delete method.
func (client *SerialPortsClient) Delete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientDeleteOptions) (SerialPortsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourceType, parentResource, serialPort, options)
	if err != nil {
		return SerialPortsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SerialPortsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SerialPortsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SerialPortsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourceType}/{parentResource}/providers/Microsoft.SerialConsole/serialPorts/{serialPort}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", parentResourceType)
	if parentResource == "" {
		return nil, errors.New("parameter parentResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResource}", url.PathEscape(parentResource))
	if serialPort == "" {
		return nil, errors.New("parameter serialPort cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialPort}", url.PathEscape(serialPort))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the configured settings for a serial port
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// resourceGroupName - The name of the resource group.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourceType - The resource type of the parent resource. For example: 'virtualMachines' or 'virtualMachineScaleSets'
// parentResource - The resource name, or subordinate path, for the parent of the serial port. For example: the name of the
// virtual machine.
// serialPort - The name of the serial port to connect to.
// options - SerialPortsClientGetOptions contains the optional parameters for the SerialPortsClient.Get method.
func (client *SerialPortsClient) Get(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientGetOptions) (SerialPortsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourceType, parentResource, serialPort, options)
	if err != nil {
		return SerialPortsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SerialPortsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SerialPortsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, serialPort string, options *SerialPortsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourceType}/{parentResource}/providers/Microsoft.SerialConsole/serialPorts/{serialPort}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", parentResourceType)
	if parentResource == "" {
		return nil, errors.New("parameter parentResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResource}", url.PathEscape(parentResource))
	if serialPort == "" {
		return nil, errors.New("parameter serialPort cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialPort}", url.PathEscape(serialPort))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SerialPortsClient) getHandleResponse(resp *http.Response) (SerialPortsClientGetResponse, error) {
	result := SerialPortsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SerialPort); err != nil {
		return SerialPortsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all of the configured serial ports for a parent resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// resourceGroupName - The name of the resource group.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourceType - The resource type of the parent resource. For example: 'virtualMachines' or 'virtualMachineScaleSets'
// parentResource - The resource name, or subordinate path, for the parent of the serial port. For example: the name of the
// virtual machine.
// options - SerialPortsClientListOptions contains the optional parameters for the SerialPortsClient.List method.
func (client *SerialPortsClient) List(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, options *SerialPortsClientListOptions) (SerialPortsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourceType, parentResource, options)
	if err != nil {
		return SerialPortsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SerialPortsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SerialPortsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourceType string, parentResource string, options *SerialPortsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourceType}/{parentResource}/providers/Microsoft.SerialConsole/serialPorts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", parentResourceType)
	if parentResource == "" {
		return nil, errors.New("parameter parentResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResource}", url.PathEscape(parentResource))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SerialPortsClient) listHandleResponse(resp *http.Response) (SerialPortsClientListResponse, error) {
	result := SerialPortsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SerialPortListResult); err != nil {
		return SerialPortsClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscriptions - Handles requests to list all SerialPort resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-05-01
// options - SerialPortsClientListBySubscriptionsOptions contains the optional parameters for the SerialPortsClient.ListBySubscriptions
// method.
func (client *SerialPortsClient) ListBySubscriptions(ctx context.Context, options *SerialPortsClientListBySubscriptionsOptions) (SerialPortsClientListBySubscriptionsResponse, error) {
	req, err := client.listBySubscriptionsCreateRequest(ctx, options)
	if err != nil {
		return SerialPortsClientListBySubscriptionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SerialPortsClientListBySubscriptionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SerialPortsClientListBySubscriptionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBySubscriptionsHandleResponse(resp)
}

// listBySubscriptionsCreateRequest creates the ListBySubscriptions request.
func (client *SerialPortsClient) listBySubscriptionsCreateRequest(ctx context.Context, options *SerialPortsClientListBySubscriptionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/serialPorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionsHandleResponse handles the ListBySubscriptions response.
func (client *SerialPortsClient) listBySubscriptionsHandleResponse(resp *http.Response) (SerialPortsClientListBySubscriptionsResponse, error) {
	result := SerialPortsClientListBySubscriptionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SerialPortListResult); err != nil {
		return SerialPortsClientListBySubscriptionsResponse{}, err
	}
	return result, nil
}
