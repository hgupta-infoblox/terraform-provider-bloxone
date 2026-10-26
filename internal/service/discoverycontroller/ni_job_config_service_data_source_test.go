
package discoverycontroller_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNIJobConfigServiceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_ni_job_configs.test"
	resourceName := "bloxone_discovery_ni_job_config.test"
	var v discoverycontroller.NIJobConfig

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIJobConfigServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIJobConfigServiceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
						}, testAccCheckNIJobConfigServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNIJobConfigServiceDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_ni_job_configs.test"
	resourceName := "bloxone_discovery_ni_job_config.test"
	var v discoverycontroller.NIJobConfig
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIJobConfigServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIJobConfigServiceDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
						}, testAccCheckNIJobConfigServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNIJobConfigServiceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "config_id", dataSourceName, "results.0.config_id"),
        resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
        resource.TestCheckResourceAttrPair(resourceName, "force_service_list_override", dataSourceName, "results.0.force_service_list_override"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "ni_job_config_name", dataSourceName, "results.0.ni_job_config_name"),
        resource.TestCheckResourceAttrPair(resourceName, "ophid", dataSourceName, "results.0.ophid"),
        resource.TestCheckResourceAttrPair(resourceName, "pool_id", dataSourceName, "results.0.pool_id"),
        resource.TestCheckResourceAttrPair(resourceName, "properties", dataSourceName, "results.0.properties"),
        resource.TestCheckResourceAttrPair(resourceName, "service_id", dataSourceName, "results.0.service_id"),
        resource.TestCheckResourceAttrPair(resourceName, "service_name", dataSourceName, "results.0.service_name"),
        resource.TestCheckResourceAttrPair(resourceName, "services", dataSourceName, "results.0.services"),
        resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
    }
}

func testAccNIJobConfigServiceDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test" {
}

data "bloxone_discovery_ni_job_configs" "test" {
  filters = {
	 = bloxone_discovery_ni_job_config.test.
  }
}
`)
}

func testAccNIJobConfigServiceDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_discovery_ni_job_configs" "test" {
  tag_filters = {
	tag1 = bloxone_discovery_ni_job_config.test.tags.tag1
  }
}
`,tagValue)
}

