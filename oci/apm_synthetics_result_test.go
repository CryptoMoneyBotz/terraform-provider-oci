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
	resultSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":       Representation{RepType: Required, Create: `${var.apm_domain_id}`},
		"execution_time":      Representation{RepType: Required, Create: `${var.execution_time}`},
		"monitor_id":          Representation{RepType: Required, Create: `${var.monitor_id}`},
		"result_content_type": Representation{RepType: Required, Create: `raw`},
		"result_type":         Representation{RepType: Required, Create: `log`},
		"vantage_point":       Representation{RepType: Required, Create: `OraclePublic-us-ashburn-1`},
	}
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsResultResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//This is a manual test. It requires apm_domain_id, monitor_id
	//and execution_time as environment variables.
	apmDomainId := getEnvSettingWithBlankDefault("apm_domain_id")
	monitorId := getEnvSettingWithBlankDefault("monitor_id")
	executionTime := getEnvSettingWithBlankDefault("execution_time")

	if apmDomainId == "" || monitorId == "" || executionTime == "" {
		t.Skip("Set apm_domain_id, monitor_id and execution_time to run this test")
	}
	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	monitorIdVariableStr := fmt.Sprintf("variable \"monitor_id\" { default = \"%s\" }\n", monitorId)
	executionTimeVariableStr := fmt.Sprintf("variable \"execution_time\" { default = \"%s\" }\n", executionTime)

	singularDatasourceName := "data.oci_apm_synthetics_result.test_result"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + monitorIdVariableStr + executionTimeVariableStr +
				GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_result", "test_result", Required, Create, resultSingularDataSourceRepresentation) +
				compartmentIdVariableStr, //+ ResultResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result_content_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result_content_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point"),
			),
		},
	})
}
