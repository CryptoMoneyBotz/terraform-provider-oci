// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"io"
	"net/http"
)

// GetPrivateApplicationPackageActionDownloadConfigRequest wrapper for the GetPrivateApplicationPackageActionDownloadConfig operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplicationPackageActionDownloadConfig.go.html to see an example of how to use GetPrivateApplicationPackageActionDownloadConfigRequest.
type GetPrivateApplicationPackageActionDownloadConfigRequest struct {

	// The unique identifier for the private application package.
	PrivateApplicationPackageId *string `mandatory:"true" contributesTo:"path" name:"privateApplicationPackageId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPrivateApplicationPackageActionDownloadConfigRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPrivateApplicationPackageActionDownloadConfigRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPrivateApplicationPackageActionDownloadConfigRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPrivateApplicationPackageActionDownloadConfigRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPrivateApplicationPackageActionDownloadConfigResponse wrapper for the GetPrivateApplicationPackageActionDownloadConfig operation
type GetPrivateApplicationPackageActionDownloadConfigResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPrivateApplicationPackageActionDownloadConfigResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPrivateApplicationPackageActionDownloadConfigResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
