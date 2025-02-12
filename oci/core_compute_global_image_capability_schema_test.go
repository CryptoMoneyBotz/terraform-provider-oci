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
	computeGlobalImageCapabilitySchemaSingularDataSourceRepresentation = map[string]interface{}{
		"compute_global_image_capability_schema_id": Representation{RepType: Required, Create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
	}

	computeGlobalImageCapabilitySchemaDataSourceRepresentation = map[string]interface{}{}

	ComputeGlobalImageCapabilitySchemaResourceConfig = ""
)

// issue-routing-tag: core/computeImaging
func TestCoreComputeGlobalImageCapabilitySchemaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGlobalImageCapabilitySchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas"
	singularDatasourceName := "data.oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", Required, Create, computeGlobalImageCapabilitySchemaDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeGlobalImageCapabilitySchemaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.current_version_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schema", "test_compute_global_image_capability_schema", Required, Create, computeGlobalImageCapabilitySchemaSingularDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", Required, Create, computeGlobalImageCapabilitySchemaDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeGlobalImageCapabilitySchemaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_version_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
