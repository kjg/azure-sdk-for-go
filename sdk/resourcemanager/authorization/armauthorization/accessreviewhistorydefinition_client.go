//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessReviewHistoryDefinitionClient contains the methods for the AccessReviewHistoryDefinition group.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionClient() instead.
type AccessReviewHistoryDefinitionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewHistoryDefinitionClient creates a new instance of AccessReviewHistoryDefinitionClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewHistoryDefinitionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewHistoryDefinitionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewHistoryDefinitionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a scheduled or one-time Access Review History Definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - historyDefinitionID - The id of the access review history definition.
//   - properties - Access review history definition properties.
//   - options - AccessReviewHistoryDefinitionClientCreateOptions contains the optional parameters for the AccessReviewHistoryDefinitionClient.Create
//     method.
func (client *AccessReviewHistoryDefinitionClient) Create(ctx context.Context, historyDefinitionID string, properties AccessReviewHistoryDefinitionProperties, options *AccessReviewHistoryDefinitionClientCreateOptions) (AccessReviewHistoryDefinitionClientCreateResponse, error) {
	var err error
	const operationName = "AccessReviewHistoryDefinitionClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, historyDefinitionID, properties, options)
	if err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *AccessReviewHistoryDefinitionClient) createCreateRequest(ctx context.Context, historyDefinitionID string, properties AccessReviewHistoryDefinitionProperties, options *AccessReviewHistoryDefinitionClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *AccessReviewHistoryDefinitionClient) createHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionClientCreateResponse, error) {
	result := AccessReviewHistoryDefinitionClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryDefinition); err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	return result, nil
}

// DeleteByID - Delete an access review history definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - historyDefinitionID - The id of the access review history definition.
//   - options - AccessReviewHistoryDefinitionClientDeleteByIDOptions contains the optional parameters for the AccessReviewHistoryDefinitionClient.DeleteByID
//     method.
func (client *AccessReviewHistoryDefinitionClient) DeleteByID(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionClientDeleteByIDOptions) (AccessReviewHistoryDefinitionClientDeleteByIDResponse, error) {
	var err error
	const operationName = "AccessReviewHistoryDefinitionClient.DeleteByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByIDCreateRequest(ctx, historyDefinitionID, options)
	if err != nil {
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, err
	}
	return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *AccessReviewHistoryDefinitionClient) deleteByIDCreateRequest(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionClientDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
