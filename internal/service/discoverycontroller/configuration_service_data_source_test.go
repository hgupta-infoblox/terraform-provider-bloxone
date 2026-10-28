
package discoverycontroller_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccConfigurationServiceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_configurations.test"
	resourceName := "bloxone_discovery_configuration.test"
	var v discoverycontroller.Configuration

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckConfigurationServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigurationServiceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
						}, testAccCheckConfigurationServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccConfigurationServiceDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_discovery_configurations.test"
	resourceName := "bloxone_discovery_configuration.test"
	var v discoverycontroller.Configuration
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckConfigurationServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigurationServiceDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
						}, testAccCheckConfigurationServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckConfigurationServiceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "assets_unmanaged", dataSourceName, "results.0.assets_unmanaged"),
        resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
        resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
        resource.TestCheckResourceAttrPair(resourceName, "destination", dataSourceName, "results.0.destination"),
        resource.TestCheckResourceAttrPair(resourceName, "destination_id", dataSourceName, "results.0.destination_id"),
        resource.TestCheckResourceAttrPair(resourceName, "enable", dataSourceName, "results.0.enable"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "source_id", dataSourceName, "results.0.source_id"),
        resource.TestCheckResourceAttrPair(resourceName, "source_name", dataSourceName, "results.0.source_name"),
        resource.TestCheckResourceAttrPair(resourceName, "source_type", dataSourceName, "results.0.source_type"),
        resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
    }
}

func testAccConfigurationServiceDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test" {
}

data "bloxone_discovery_configurations" "test" {
  filters = {
	 = bloxone_discovery_configuration.test.
  }
}
`)
}

func testAccConfigurationServiceDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_discovery_configurations" "test" {
  tag_filters = {
	tag1 = bloxone_discovery_configuration.test.tags.tag1
  }
}
`,tagValue)
}

