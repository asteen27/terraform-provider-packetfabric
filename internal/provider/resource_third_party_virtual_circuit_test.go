package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclBackboneVirtualCircuitMarketplace(pop, zone, description, routingID, market, vlan, accountUUID, speed string) (hcl string, resourceName string) {
	portHcl, portResourceName := hclPort(
		testutil.GenerateUniqueName(testPrefix),
		accountUUID,
		speed,
		"LX",
		"1",
		pop,
		zone,
		true,
		false,
	)
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_backbone_virtual_circuit_marketplace." + hclName
	backboneVcHcl := fmt.Sprintf(`
	resource "packetfabric_backbone_virtual_circuit_marketplace" "%s" {
		description = "%s"
		routing_id  = "%s"
		market      = "%s"
		interface {
			port_circuit_id = %s.id
			untagged        = false
			vlan            = "%s"
		}
		bandwidth {
			account_uuid  = "%s"
			longhaul_type = "usage"
		}
	}`, hclName, description, routingID, market, portResourceName, vlan, accountUUID)

	hcl = fmt.Sprintf("%s\n%s", portHcl, backboneVcHcl)
	return
}

func TestAccBackboneVirtualCircuitMarketplace(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := testutil.GenerateUniqueName(testPrefix)
	pop, zone, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps", "LX")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclBackboneVirtualCircuitMarketplace(
		pop,
		zone,
		description,
		os.Getenv("PF_ACC_TEST_ROUTING_ID"), // TODO(medzin): Refactor to use the data from the customer endpoint.
		os.Getenv("PF_ACC_TEST_MARKET"),     // TODO(medzin): Refactor to use the data from the customer endpoint.
		"5",
		testutil.GetAccountUUID(),
		"1Gbps",
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				"PF_ACC_TEST_ROUTING_ID",
				"PF_ACC_TEST_MARKET",
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "interface.0.vlan", "5"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.account_uuid", testutil.GetAccountUUID()),
				),
			},
		},
	})
}
