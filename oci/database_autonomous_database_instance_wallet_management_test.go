// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseInstanceWalletManagementResourceConfig = AutonomousDatabaseInstanceWalletManagementResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Optional, Update, autonomousDatabaseInstanceWalletManagementRepresentation)

	autonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseInstanceWalletManagementRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"should_rotate":          Representation{repType: Optional, create: `false`, update: `true`},
	}

	AutonomousDatabaseInstanceWalletManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation)
)

func TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	singularDatasourceName := "data.oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Required, Create, autonomousDatabaseInstanceWalletManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Optional, Update, autonomousDatabaseInstanceWalletManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_rotated"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Required, Create, autonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
				),
			},
		},
	})
}
