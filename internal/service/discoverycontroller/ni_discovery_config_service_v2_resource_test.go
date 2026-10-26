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

func TestAccNIDiscoveryConfigServiceV2Resource_basic(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2BasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "ni_job_config_name"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_disappears(t *testing.T) {
	resourceName := "bloxone_discovery_ni_service_config_v2.test"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNIDiscoveryConfigServiceV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNIDiscoveryConfigServiceV2BasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					testAccCheckNIDiscoveryConfigServiceV2Disappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_HostId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_host_id"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2HostId("HOST_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_id", "HOST_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2HostId("HOST_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_id", "HOST_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_NiDiscoveryConfig(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_ni_discovery_config"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2NiDiscoveryConfig("NI_DISCOVERY_CONFIG_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ni_discovery_config", "NI_DISCOVERY_CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2NiDiscoveryConfig("NI_DISCOVERY_CONFIG_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ni_discovery_config", "NI_DISCOVERY_CONFIG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_NiJobConfigId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_ni_job_config_id"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2NiJobConfigId("NI_JOB_CONFIG_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ni_job_config_id", "NI_JOB_CONFIG_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2NiJobConfigId("NI_JOB_CONFIG_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ni_job_config_id", "NI_JOB_CONFIG_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_Ophid(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_ophid"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2Ophid("OPHID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ophid", "OPHID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2Ophid("OPHID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ophid", "OPHID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_PoolId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_pool_id"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2PoolId("POOL_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pool_id", "POOL_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2PoolId("POOL_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pool_id", "POOL_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNIDiscoveryConfigServiceV2Resource_ServiceId(t *testing.T) {
	var resourceName = "bloxone_discovery_ni_service_config_v2.test_service_id"
	var v discoverycontroller.NIDiscoveryServiceV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2ServiceId("SERVICE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_id", "SERVICE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNIDiscoveryConfigServiceV2ServiceId("SERVICE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNIDiscoveryConfigServiceV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_id", "SERVICE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNIDiscoveryConfigServiceV2Exists(ctx context.Context, resourceName string, v *discoverycontroller.NIDiscoveryServiceV2) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			NIDiscoveryConfigServiceV2API.
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

func testAccCheckNIDiscoveryConfigServiceV2Destroy(ctx context.Context, v *discoverycontroller.NIDiscoveryServiceV2) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			NIDiscoveryConfigServiceV2API.
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

func testAccCheckNIDiscoveryConfigServiceV2Disappears(ctx context.Context, v *discoverycontroller.NIDiscoveryServiceV2) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DiscoveryControllerAPI.
			NIDiscoveryConfigServiceV2API.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNIDiscoveryConfigServiceV2BasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test" {
}
`)
}

func testAccNIDiscoveryConfigServiceV2HostId(hostId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_host_id" {
    host_id = %q
}
`, hostId)
}

func testAccNIDiscoveryConfigServiceV2NiDiscoveryConfig(niDiscoveryConfig string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_ni_discovery_config" {
    ni_discovery_config = %q
}
`, niDiscoveryConfig)
}

func testAccNIDiscoveryConfigServiceV2NiJobConfigId(niJobConfigId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_ni_job_config_id" {
    ni_job_config_id = %q
}
`, niJobConfigId)
}

func testAccNIDiscoveryConfigServiceV2Ophid(ophid string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_ophid" {
    ophid = %q
}
`, ophid)
}

func testAccNIDiscoveryConfigServiceV2PoolId(poolId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_pool_id" {
    pool_id = %q
}
`, poolId)
}

func testAccNIDiscoveryConfigServiceV2ServiceId(serviceId string) string {
	return fmt.Sprintf(`
resource "bloxone_discovery_ni_service_config_v2" "test_service_id" {
    service_id = %q
}
`, serviceId)
}
