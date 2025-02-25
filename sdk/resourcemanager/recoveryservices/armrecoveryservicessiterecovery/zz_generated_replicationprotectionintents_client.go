//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationProtectionIntentsClient contains the methods for the ReplicationProtectionIntents group.
// Don't use this type directly, use NewReplicationProtectionIntentsClient() instead.
type ReplicationProtectionIntentsClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationProtectionIntentsClient creates a new instance of ReplicationProtectionIntentsClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationProtectionIntentsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationProtectionIntentsClient, error) {
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
	client := &ReplicationProtectionIntentsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Create - The operation to create an ASR replication protection intent item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// intentObjectName - A name for the replication protection item.
// input - Create Protection Intent Input.
// options - ReplicationProtectionIntentsClientCreateOptions contains the optional parameters for the ReplicationProtectionIntentsClient.Create
// method.
func (client *ReplicationProtectionIntentsClient) Create(ctx context.Context, intentObjectName string, input CreateProtectionIntentInput, options *ReplicationProtectionIntentsClientCreateOptions) (ReplicationProtectionIntentsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, intentObjectName, input, options)
	if err != nil {
		return ReplicationProtectionIntentsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationProtectionIntentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationProtectionIntentsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ReplicationProtectionIntentsClient) createCreateRequest(ctx context.Context, intentObjectName string, input CreateProtectionIntentInput, options *ReplicationProtectionIntentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionIntents/{intentObjectName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if intentObjectName == "" {
		return nil, errors.New("parameter intentObjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intentObjectName}", url.PathEscape(intentObjectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, input)
}

// createHandleResponse handles the Create response.
func (client *ReplicationProtectionIntentsClient) createHandleResponse(resp *http.Response) (ReplicationProtectionIntentsClientCreateResponse, error) {
	result := ReplicationProtectionIntentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationProtectionIntent); err != nil {
		return ReplicationProtectionIntentsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of an ASR replication protection intent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// intentObjectName - Replication protection intent name.
// options - ReplicationProtectionIntentsClientGetOptions contains the optional parameters for the ReplicationProtectionIntentsClient.Get
// method.
func (client *ReplicationProtectionIntentsClient) Get(ctx context.Context, intentObjectName string, options *ReplicationProtectionIntentsClientGetOptions) (ReplicationProtectionIntentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, intentObjectName, options)
	if err != nil {
		return ReplicationProtectionIntentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationProtectionIntentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationProtectionIntentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectionIntentsClient) getCreateRequest(ctx context.Context, intentObjectName string, options *ReplicationProtectionIntentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionIntents/{intentObjectName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if intentObjectName == "" {
		return nil, errors.New("parameter intentObjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intentObjectName}", url.PathEscape(intentObjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationProtectionIntentsClient) getHandleResponse(resp *http.Response) (ReplicationProtectionIntentsClientGetResponse, error) {
	result := ReplicationProtectionIntentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationProtectionIntent); err != nil {
		return ReplicationProtectionIntentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of ASR replication protection intent objects in the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// options - ReplicationProtectionIntentsClientListOptions contains the optional parameters for the ReplicationProtectionIntentsClient.List
// method.
func (client *ReplicationProtectionIntentsClient) NewListPager(options *ReplicationProtectionIntentsClientListOptions) *runtime.Pager[ReplicationProtectionIntentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionIntentsClientListResponse]{
		More: func(page ReplicationProtectionIntentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionIntentsClientListResponse) (ReplicationProtectionIntentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationProtectionIntentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationProtectionIntentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationProtectionIntentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationProtectionIntentsClient) listCreateRequest(ctx context.Context, options *ReplicationProtectionIntentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionIntents"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("skipToken", *options.SkipToken)
	}
	if options != nil && options.TakeToken != nil {
		reqQP.Set("takeToken", *options.TakeToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationProtectionIntentsClient) listHandleResponse(resp *http.Response) (ReplicationProtectionIntentsClientListResponse, error) {
	result := ReplicationProtectionIntentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationProtectionIntentCollection); err != nil {
		return ReplicationProtectionIntentsClientListResponse{}, err
	}
	return result, nil
}
