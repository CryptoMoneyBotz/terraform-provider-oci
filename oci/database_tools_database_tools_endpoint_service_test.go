// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	databaseToolsEndpointServiceSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_endpoint_service_id": Representation{RepType: Required, Create: `${data.oci_database_tools_database_tools_endpoint_services.test_database_tools_endpoint_services.database_tools_endpoint_service_collection.0.items.0.id}`},
	}

	databaseToolsEndpointServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`},
		"name":           Representation{RepType: Optional, Create: `name`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
	}

	DatabaseToolsEndpointServiceResourceConfig = ""
)

func TestDatabaseToolsDatabaseToolsEndpointServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsEndpointServiceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_tools_database_tools_endpoint_services.test_database_tools_endpoint_services"
	singularDatasourceName := "data.oci_database_tools_database_tools_endpoint_service.test_database_tools_endpoint_service"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// 0. verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsEndpointServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_endpoint_service_collection.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_endpoint_service_collection.0.items.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_endpoint_service_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_endpoint_service_collection.0.items.0.time_created"),
				),
			},
			// 1. verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", Required, Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsEndpointServiceResourceConfig,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_endpoint_service_id"),

					//resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"), // endpoint service is not specific to a compartment, so this is expected.
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
		},
	})
}
