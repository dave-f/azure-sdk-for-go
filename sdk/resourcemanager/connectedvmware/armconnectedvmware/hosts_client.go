//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

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
	"strconv"
	"strings"
)

// HostsClient contains the methods for the Hosts group.
// Don't use this type directly, use NewHostsClient() instead.
type HostsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHostsClient creates a new instance of HostsClient with the specified values.
// subscriptionID - The Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HostsClient, error) {
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
	client := &HostsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create Or Update host.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// hostName - Name of the host.
// body - Request payload.
// options - HostsClientBeginCreateOptions contains the optional parameters for the HostsClient.BeginCreate method.
func (client *HostsClient) BeginCreate(ctx context.Context, resourceGroupName string, hostName string, body Host, options *HostsClientBeginCreateOptions) (*runtime.Poller[HostsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, hostName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[HostsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[HostsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create Or Update host.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
func (client *HostsClient) create(ctx context.Context, resourceGroupName string, hostName string, body Host, options *HostsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, hostName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *HostsClient) createCreateRequest(ctx context.Context, resourceGroupName string, hostName string, body Host, options *HostsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/hosts/{hostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements host DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// hostName - Name of the host.
// options - HostsClientBeginDeleteOptions contains the optional parameters for the HostsClient.BeginDelete method.
func (client *HostsClient) BeginDelete(ctx context.Context, resourceGroupName string, hostName string, options *HostsClientBeginDeleteOptions) (*runtime.Poller[HostsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hostName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[HostsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[HostsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Implements host DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
func (client *HostsClient) deleteOperation(ctx context.Context, resourceGroupName string, hostName string, options *HostsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostName string, options *HostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/hosts/{hostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements host GET method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// hostName - Name of the host.
// options - HostsClientGetOptions contains the optional parameters for the HostsClient.Get method.
func (client *HostsClient) Get(ctx context.Context, resourceGroupName string, hostName string, options *HostsClientGetOptions) (HostsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostName, options)
	if err != nil {
		return HostsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HostsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostName string, options *HostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/hosts/{hostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HostsClient) getHandleResponse(resp *http.Response) (HostsClientGetResponse, error) {
	result := HostsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Host); err != nil {
		return HostsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List of hosts in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// options - HostsClientListOptions contains the optional parameters for the HostsClient.List method.
func (client *HostsClient) NewListPager(options *HostsClientListOptions) *runtime.Pager[HostsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HostsClientListResponse]{
		More: func(page HostsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HostsClientListResponse) (HostsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HostsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HostsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HostsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HostsClient) listCreateRequest(ctx context.Context, options *HostsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HostsClient) listHandleResponse(resp *http.Response) (HostsClientListResponse, error) {
	result := HostsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostsList); err != nil {
		return HostsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of hosts in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// options - HostsClientListByResourceGroupOptions contains the optional parameters for the HostsClient.ListByResourceGroup
// method.
func (client *HostsClient) NewListByResourceGroupPager(resourceGroupName string, options *HostsClientListByResourceGroupOptions) *runtime.Pager[HostsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[HostsClientListByResourceGroupResponse]{
		More: func(page HostsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HostsClientListByResourceGroupResponse) (HostsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HostsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HostsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HostsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *HostsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *HostsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/hosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *HostsClient) listByResourceGroupHandleResponse(resp *http.Response) (HostsClientListByResourceGroupResponse, error) {
	result := HostsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostsList); err != nil {
		return HostsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the host resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// resourceGroupName - The Resource Group Name.
// hostName - Name of the host.
// body - Resource properties to update.
// options - HostsClientUpdateOptions contains the optional parameters for the HostsClient.Update method.
func (client *HostsClient) Update(ctx context.Context, resourceGroupName string, hostName string, body ResourcePatch, options *HostsClientUpdateOptions) (HostsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostName, body, options)
	if err != nil {
		return HostsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HostsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HostsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostName string, body ResourcePatch, options *HostsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/hosts/{hostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *HostsClient) updateHandleResponse(resp *http.Response) (HostsClientUpdateResponse, error) {
	result := HostsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Host); err != nil {
		return HostsClientUpdateResponse{}, err
	}
	return result, nil
}
