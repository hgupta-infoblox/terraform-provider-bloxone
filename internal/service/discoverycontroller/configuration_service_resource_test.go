package discoverycontroller_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"
)

func TestAccConfigurationServiceResource_basic(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_disappears(t *testing.T) {
	resourceName := "bloxone_discovery_configuration.test"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckConfigurationServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigurationServiceBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					testAccCheckConfigurationServiceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccConfigurationServiceResource_AssetsUnmanaged(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_assets_unmanaged"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceAssetsUnmanaged("ASSETS_UNMANAGED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "assets_unmanaged", "ASSETS_UNMANAGED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceAssetsUnmanaged("ASSETS_UNMANAGED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "assets_unmanaged", "ASSETS_UNMANAGED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_Description(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_description"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceDescription("DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceDescription("DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_Destination(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_destination"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceDestination("DESTINATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination", "DESTINATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceDestination("DESTINATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination", "DESTINATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_DestinationId(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_destination_id"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceDestinationId("DESTINATION_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination_id", "DESTINATION_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceDestinationId("DESTINATION_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination_id", "DESTINATION_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_Enable(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_enable"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceEnable("ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable", "ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceEnable("ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable", "ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_Name(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_name"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_SourceId(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_source_id"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceSourceId("SOURCE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_id", "SOURCE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceSourceId("SOURCE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_id", "SOURCE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_SourceName(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_source_name"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceSourceName("SOURCE_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_name", "SOURCE_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceSourceName("SOURCE_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_name", "SOURCE_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccConfigurationServiceResource_SourceType(t *testing.T) {
	var resourceName = "bloxone_discovery_configuration.test_source_type"
	var v discoverycontroller.Configuration

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccConfigurationServiceSourceType("SOURCE_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_type", "SOURCE_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccConfigurationServiceSourceType("SOURCE_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConfigurationServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_type", "SOURCE_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckConfigurationServiceExists(ctx context.Context, resourceName string, v *discoverycontroller.Configuration) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			ConfigurationServiceAPI.
			Read(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckConfigurationServiceDestroy(ctx context.Context, v *discoverycontroller.Configuration) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			ConfigurationServiceAPI.
			Read(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckConfigurationServiceDisappears(ctx context.Context, v *discoverycontroller.Configuration) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			ConfigurationServiceAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccConfigurationServiceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test" {
}
`)
}

func testAccConfigurationServiceAssetsUnmanaged(assetsUnmanaged string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_assets_unmanaged" {
    assets_unmanaged = %q
}
`, assetsUnmanaged)
}

func testAccConfigurationServiceDescription(description string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_description" {
    description = %q
}
`, description)
}

func testAccConfigurationServiceDestination(destination string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_destination" {
    destination = %q
}
`, destination)
}

func testAccConfigurationServiceDestinationId(destinationId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_destination_id" {
    destination_id = %q
}
`, destinationId)
}

func testAccConfigurationServiceEnable(enable string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_enable" {
    enable = %q
}
`, enable)
}

func testAccConfigurationServiceName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_name" {
    name = %q
}
`, name)
}

func testAccConfigurationServiceSourceId(sourceId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_source_id" {
    source_id = %q
}
`, sourceId)
}

func testAccConfigurationServiceSourceName(sourceName string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_source_name" {
    source_name = %q
}
`, sourceName)
}

func testAccConfigurationServiceSourceType(sourceType string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_configuration" "test_source_type" {
    source_type = %q
}
`, sourceType)
}
