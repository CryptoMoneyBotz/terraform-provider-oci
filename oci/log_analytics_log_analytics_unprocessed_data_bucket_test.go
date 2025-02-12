// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsUnprocessedDataBucketSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsUnprocessedDataBucketDependencies = "" +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)

	LogAnalyticsUnprocessedDataBucketResourceConfig = ""
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsUnprocessedDataBucketResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsUnprocessedDataBucketResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_unprocessed_data_bucket_management.test_log_analytics_unprocessed_data_bucket_management"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_unprocessed_data_bucket.test_log_analytics_unprocessed_data_bucket"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// set unprocessed data bucket
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", Required, Update, logAnalyticsUnprocessedDataBucketManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "udb_tf"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
			),
		},

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketDependencies +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket", "test_log_analytics_unprocessed_data_bucket", Required, Create, logAnalyticsUnprocessedDataBucketSingularDataSourceRepresentation) +
				LogAnalyticsUnprocessedDataBucketResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "udb_tf"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
