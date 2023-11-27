//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessReviewInstanceDecisionsServer is a fake server for instances of the armauthorization.AccessReviewInstanceDecisionsClient type.
type AccessReviewInstanceDecisionsServer struct {
	// NewListPager is the fake for method AccessReviewInstanceDecisionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceDecisionsClientListOptions) (resp azfake.PagerResponder[armauthorization.AccessReviewInstanceDecisionsClientListResponse])
}

// NewAccessReviewInstanceDecisionsServerTransport creates a new instance of AccessReviewInstanceDecisionsServerTransport with the provided implementation.
// The returned AccessReviewInstanceDecisionsServerTransport instance is connected to an instance of armauthorization.AccessReviewInstanceDecisionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessReviewInstanceDecisionsServerTransport(srv *AccessReviewInstanceDecisionsServer) *AccessReviewInstanceDecisionsServerTransport {
	return &AccessReviewInstanceDecisionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.AccessReviewInstanceDecisionsClientListResponse]](),
	}
}

// AccessReviewInstanceDecisionsServerTransport connects instances of armauthorization.AccessReviewInstanceDecisionsClient to instances of AccessReviewInstanceDecisionsServer.
// Don't use this type directly, use NewAccessReviewInstanceDecisionsServerTransport instead.
type AccessReviewInstanceDecisionsServerTransport struct {
	srv          *AccessReviewInstanceDecisionsServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.AccessReviewInstanceDecisionsClientListResponse]]
}

// Do implements the policy.Transporter interface for AccessReviewInstanceDecisionsServerTransport.
func (a *AccessReviewInstanceDecisionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccessReviewInstanceDecisionsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccessReviewInstanceDecisionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/decisions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
		if err != nil {
			return nil, err
		}
		idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.AccessReviewInstanceDecisionsClientListOptions
		if filterParam != nil {
			options = &armauthorization.AccessReviewInstanceDecisionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := a.srv.NewListPager(scheduleDefinitionIDParam, idParam, options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.AccessReviewInstanceDecisionsClientListResponse, createLink func() string) {
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
