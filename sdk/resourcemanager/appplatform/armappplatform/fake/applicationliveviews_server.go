//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationLiveViewsServer is a fake server for instances of the armappplatform.ApplicationLiveViewsClient type.
type ApplicationLiveViewsServer struct {
	// BeginCreateOrUpdate is the fake for method ApplicationLiveViewsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource armappplatform.ApplicationLiveViewResource, options *armappplatform.ApplicationLiveViewsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationLiveViewsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *armappplatform.ApplicationLiveViewsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationLiveViewsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *armappplatform.ApplicationLiveViewsClientGetOptions) (resp azfake.Responder[armappplatform.ApplicationLiveViewsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationLiveViewsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, options *armappplatform.ApplicationLiveViewsClientListOptions) (resp azfake.PagerResponder[armappplatform.ApplicationLiveViewsClientListResponse])
}

// NewApplicationLiveViewsServerTransport creates a new instance of ApplicationLiveViewsServerTransport with the provided implementation.
// The returned ApplicationLiveViewsServerTransport instance is connected to an instance of armappplatform.ApplicationLiveViewsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationLiveViewsServerTransport(srv *ApplicationLiveViewsServer) *ApplicationLiveViewsServerTransport {
	return &ApplicationLiveViewsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappplatform.ApplicationLiveViewsClientListResponse]](),
	}
}

// ApplicationLiveViewsServerTransport connects instances of armappplatform.ApplicationLiveViewsClient to instances of ApplicationLiveViewsServer.
// Don't use this type directly, use NewApplicationLiveViewsServerTransport instead.
type ApplicationLiveViewsServerTransport struct {
	srv                 *ApplicationLiveViewsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappplatform.ApplicationLiveViewsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappplatform.ApplicationLiveViewsClientListResponse]]
}

// Do implements the policy.Transporter interface for ApplicationLiveViewsServerTransport.
func (a *ApplicationLiveViewsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationLiveViewsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "ApplicationLiveViewsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "ApplicationLiveViewsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationLiveViewsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationLiveViewsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationLiveViews/(?P<applicationLiveViewName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.ApplicationLiveViewResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationLiveViewNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationLiveViewName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, applicationLiveViewNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *ApplicationLiveViewsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationLiveViews/(?P<applicationLiveViewName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationLiveViewNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationLiveViewName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, applicationLiveViewNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *ApplicationLiveViewsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationLiveViews/(?P<applicationLiveViewName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	applicationLiveViewNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationLiveViewName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, applicationLiveViewNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationLiveViewResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationLiveViewsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationLiveViews`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, serviceNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.ApplicationLiveViewsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
