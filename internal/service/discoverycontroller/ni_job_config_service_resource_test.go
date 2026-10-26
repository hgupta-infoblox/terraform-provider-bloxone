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

func TestAccNIJobConfigServiceResource_basic(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "ni_job_config_name"),
					resource.TestCheckResourceAttrSet(resourceName, "service_name"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_disappears(t *testing.T) {
	resourceName := "bloxone_discovery_ni_job_config.test"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIJobConfigServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIJobConfigServiceBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					testAccCheckNIJobConfigServiceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNIJobConfigServiceResource_ConfigId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_config_id"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceConfigId("CONFIG_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_id", "CONFIG_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceConfigId("CONFIG_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_id", "CONFIG_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_ForceServiceListOverride(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_force_service_list_override"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceForceServiceListOverride("FORCE_SERVICE_LIST_OVERRIDE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "force_service_list_override", "FORCE_SERVICE_LIST_OVERRIDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceForceServiceListOverride("FORCE_SERVICE_LIST_OVERRIDE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "force_service_list_override", "FORCE_SERVICE_LIST_OVERRIDE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_Ophid(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_ophid"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceOphid("OPHID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ophid", "OPHID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceOphid("OPHID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ophid", "OPHID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_PoolId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_pool_id"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServicePoolId("POOL_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pool_id", "POOL_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServicePoolId("POOL_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pool_id", "POOL_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_Properties(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_properties"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceProperties("PROPERTIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "properties", "PROPERTIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceProperties("PROPERTIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "properties", "PROPERTIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_ServiceId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_service_id"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceServiceId("SERVICE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_id", "SERVICE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceServiceId("SERVICE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_id", "SERVICE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIJobConfigServiceResource_Services(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_job_config.test_services"
	var v discoverycontroller.NIJobConfig

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIJobConfigServiceServices("SERVICES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SERVICES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIJobConfigServiceServices("SERVICES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIJobConfigServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SERVICES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNIJobConfigServiceExists(ctx context.Context, resourceName string, v *discoverycontroller.NIJobConfig) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			NIJobConfigServiceAPI.
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

func testAccCheckNIJobConfigServiceDestroy(ctx context.Context, v *discoverycontroller.NIJobConfig) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			NIJobConfigServiceAPI.
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

func testAccCheckNIJobConfigServiceDisappears(ctx context.Context, v *discoverycontroller.NIJobConfig) resource.TestCheckFunc {
	// NIJobConfigService has no Delete endpoint. Deleting the parent discovery_configuration
	// cascades and removes this ni_job_config along with it, simulating an out-of-band delete.
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			ConfigurationServiceAPI.
			Delete(ctx, *v.ConfigId).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNIJobConfigServiceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test" {
}
`)
}

func testAccNIJobConfigServiceConfigId(configId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_config_id" {
    config_id = %q
}
`, configId)
}

func testAccNIJobConfigServiceForceServiceListOverride(forceServiceListOverride string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_force_service_list_override" {
    force_service_list_override = %q
}
`, forceServiceListOverride)
}

func testAccNIJobConfigServiceOphid(ophid string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_ophid" {
    ophid = %q
}
`, ophid)
}

func testAccNIJobConfigServicePoolId(poolId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_pool_id" {
    pool_id = %q
}
`, poolId)
}

func testAccNIJobConfigServiceProperties(properties string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_properties" {
    properties = %q
}
`, properties)
}

func testAccNIJobConfigServiceServiceId(serviceId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_service_id" {
    service_id = %q
}
`, serviceId)
}

func testAccNIJobConfigServiceServices(services string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_job_config" "test_services" {
    services = %q
}
`, services)
}
