//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

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

// OperationClient contains the methods for the SaaSOperation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationClient, error) {
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
	client := &OperationClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginGet - Gets information about the specified operation progress.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01-beta
// operationID - the operation Id parameter.
// options - OperationClientBeginGetOptions contains the optional parameters for the OperationClient.BeginGet method.
func (client *OperationClient) BeginGet(ctx context.Context, operationID string, options *OperationClientBeginGetOptions) (*runtime.Poller[OperationClientGetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.get(ctx, operationID, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[OperationClientGetResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OperationClientGetResponse](options.ResumeToken, client.pl, nil)
	}
}

// Get - Gets information about the specified operation progress.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01-beta
func (client *OperationClient) get(ctx context.Context, operationID string, options *OperationClientBeginGetOptions) (*http.Response, error) {
	req, err := client.getCreateRequest(ctx, operationID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// getCreateRequest creates the Get request.
func (client *OperationClient) getCreateRequest(ctx context.Context, operationID string, options *OperationClientBeginGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/operationResults/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
