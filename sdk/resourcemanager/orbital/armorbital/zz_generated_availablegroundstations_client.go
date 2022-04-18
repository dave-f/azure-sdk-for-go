//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

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

// AvailableGroundStationsClient contains the methods for the AvailableGroundStations group.
// Don't use this type directly, use NewAvailableGroundStationsClient() instead.
type AvailableGroundStationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableGroundStationsClient creates a new instance of AvailableGroundStationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableGroundStationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableGroundStationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableGroundStationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the specified available ground station
// If the operation fails it returns an *azcore.ResponseError type.
// groundStationName - Ground Station name
// options - AvailableGroundStationsClientGetOptions contains the optional parameters for the AvailableGroundStationsClient.Get
// method.
func (client *AvailableGroundStationsClient) Get(ctx context.Context, groundStationName string, options *AvailableGroundStationsClientGetOptions) (AvailableGroundStationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groundStationName, options)
	if err != nil {
		return AvailableGroundStationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailableGroundStationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailableGroundStationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AvailableGroundStationsClient) getCreateRequest(ctx context.Context, groundStationName string, options *AvailableGroundStationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Orbital/availableGroundStations/{groundStationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groundStationName == "" {
		return nil, errors.New("parameter groundStationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groundStationName}", url.PathEscape(groundStationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AvailableGroundStationsClient) getHandleResponse(resp *http.Response) (AvailableGroundStationsClientGetResponse, error) {
	result := AvailableGroundStationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableGroundStation); err != nil {
		return AvailableGroundStationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCapabilityPager - Returns list of available ground stations
// If the operation fails it returns an *azcore.ResponseError type.
// capability - Ground Station Capability
// options - AvailableGroundStationsClientListByCapabilityOptions contains the optional parameters for the AvailableGroundStationsClient.ListByCapability
// method.
func (client *AvailableGroundStationsClient) NewListByCapabilityPager(capability CapabilityType, options *AvailableGroundStationsClientListByCapabilityOptions) *runtime.Pager[AvailableGroundStationsClientListByCapabilityResponse] {
	return runtime.NewPager(runtime.PageProcessor[AvailableGroundStationsClientListByCapabilityResponse]{
		More: func(page AvailableGroundStationsClientListByCapabilityResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableGroundStationsClientListByCapabilityResponse) (AvailableGroundStationsClientListByCapabilityResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCapabilityCreateRequest(ctx, capability, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailableGroundStationsClientListByCapabilityResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AvailableGroundStationsClientListByCapabilityResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailableGroundStationsClientListByCapabilityResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCapabilityHandleResponse(resp)
		},
	})
}

// listByCapabilityCreateRequest creates the ListByCapability request.
func (client *AvailableGroundStationsClient) listByCapabilityCreateRequest(ctx context.Context, capability CapabilityType, options *AvailableGroundStationsClientListByCapabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Orbital/availableGroundStations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	reqQP.Set("capability", string(capability))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByCapabilityHandleResponse handles the ListByCapability response.
func (client *AvailableGroundStationsClient) listByCapabilityHandleResponse(resp *http.Response) (AvailableGroundStationsClientListByCapabilityResponse, error) {
	result := AvailableGroundStationsClientListByCapabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableGroundStationListResult); err != nil {
		return AvailableGroundStationsClientListByCapabilityResponse{}, err
	}
	return result, nil
}
