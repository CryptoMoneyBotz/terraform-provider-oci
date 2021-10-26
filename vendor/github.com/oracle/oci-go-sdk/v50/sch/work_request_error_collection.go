// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// WorkRequestErrorCollection Collection of work request errors.
type WorkRequestErrorCollection struct {

	// The list of items.
	Items []WorkRequestError `mandatory:"true" json:"items"`
}

func (m WorkRequestErrorCollection) String() string {
	return common.PointerString(m)
}
