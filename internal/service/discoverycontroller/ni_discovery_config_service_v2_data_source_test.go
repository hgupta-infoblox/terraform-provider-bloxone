
package discoverycontroller_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNIDiscoveryConfigServiceV2DataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_ni_service_configs_v2.test"
	resourceName := "bloxone_discovery_ni_service_config_v2.test"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIDiscoveryConfigServiceV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIDiscoveryConfigServiceV2DataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
						}, testAccCheckNIDiscoveryConfigServiceV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2DataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_ni_service_configs_v2.test"
	resourceName := "bloxone_discovery_ni_service_config_v2.test"
	var v discoverycontroller.NIDiscoveryServiceV2
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIDiscoveryConfigServiceV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIDiscoveryConfigServiceV2DataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
						}, testAccCheckNIDiscoveryConfigServiceV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNIDiscoveryConfigServiceV2ResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
        resource.TestCheckResourceAttrPair(resourceName, "host_id", dataSourceName, "results.0.host_id"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "ni_discovery_config", dataSourceName, "results.0.ni_discovery_config"),
        resource.TestCheckResourceAttrPair(resourceName, "ni_job_config_id", dataSourceName, "results.0.ni_job_config_id"),
        resource.TestCheckResourceAttrPair(resourceName, "ni_job_config_name", dataSourceName, "results.0.ni_job_config_name"),
        resource.TestCheckResourceAttrPair(resourceName, "ophid", dataSourceName, "results.0.ophid"),
        resource.TestCheckResourceAttrPair(resourceName, "pool_id", dataSourceName, "results.0.pool_id"),
        resource.TestCheckResourceAttrPair(resourceName, "service_id", dataSourceName, "results.0.service_id"),
        resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
    }
}

func testAccNIDiscoveryConfigServiceV2DataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test" {
}

data "bloxone_discovery_ni_service_configs_v2" "test" {
  filters = {
	 = bloxone_discovery_ni_service_config_v2.test.
  }
}
`)
}

func testAccNIDiscoveryConfigServiceV2DataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_discovery_ni_service_configs_v2" "test" {
  tag_filters = {
	tag1 = bloxone_discovery_ni_service_config_v2.test.tags.tag1
  }
}
`,tagValue)
}

