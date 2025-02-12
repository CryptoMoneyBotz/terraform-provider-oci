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
	caNameForCaVersionTests = "test-ca-version-ca-" + RandomString(10, charsetWithoutDigits)

	certificateAuthorityVersionSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id":             Representation{RepType: Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
		"certificate_authority_version_number": Representation{RepType: Required, Create: `1`},
	}

	certificateAuthorityVersionDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id": Representation{RepType: Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
		"version_number":           Representation{RepType: Optional, Create: `1`},
	}

	CertificateAuthorityVersionResourceConfig = GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", Required, Create,
		RepresentationCopyWithNewProperties(certificateAuthorityRepresentation, map[string]interface{}{
			"name": Representation{RepType: Required, Create: caNameForCaVersionTests},
		}))
)

func TestCertificatesManagementCertificateAuthorityVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_certificate_authority_versions.test_certificate_authority_versions"
	singularDatasourceName := "data.oci_certificates_management_certificate_authority_version.test_certificate_authority_version"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_versions", "test_certificate_authority_versions", Optional, Create, certificateAuthorityVersionDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateAuthorityVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_authority_id"),
					resource.TestCheckResourceAttr(datasourceName, "version_number", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_authority_version_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_version_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_version_collection.0.items.0.version_number", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_version", Required, Create, certificateAuthorityVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateAuthorityVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_authority_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "version_number", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),

					resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "validity.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
