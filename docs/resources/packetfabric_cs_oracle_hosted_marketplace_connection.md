---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cs_oracle_hosted_marketplace_connection Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_oracle_hosted_marketplace_connection (Resource)

Connect a marketplace provider to your Oracle cloud environment. For more information, see [Cloud Connections in the PacketFabric documentation](https://docs.packetfabric.com/cloud/) and [Marketplace to Cloud Connections](https://docs.packetfabric.com/eco/marketplace_cloud/).

Once your request has been approved, [import the resource](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/guides/importing) as [`packetfabric_cs_oracle_hosted_connection`](https://registry.terraform.io/providers/PacketFabric/packetfabric/latest/docs/resources/packetfabric_cs_oracle_hosted_connection) into Terraform.

## Example Usage

```terraform
resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
  provider    = packetfabric
  description = var.pf_description
  vc_ocid     = var.oracle_vc_ocid
  region      = var.oracle_region
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  pop         = var.pf_cs_pop
}

output "packetfabric_cs_oracle_hosted_marketplace_connection" {
  value     = packetfabric_cs_oracle_hosted_marketplace_connection.cs_conn1_marketplace_oracle
  sensitive = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed. Can also be set with the PF_ACCOUNT_ID environment variable.
- `market` (String) The market code (e.g. "ATL" or "DAL") in which you would like the marketplace provider to provision their side of the connection.

	If the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in "DAL5", enter "DAL" for the market).
- `pop` (String) The POP in which the connection should be provisioned (the cloud on-ramp).
- `region` (String) The region in which you created the FastConnect virtual circuit.
- `routing_id` (String) The routing ID of the marketplace provider that will be receiving this request.

	Example: TR-1RI-OQ85
- `vc_ocid` (String) OCID of the FastConnect virtual circuit that you created from the Oracle side.

### Optional

- `description` (String) A brief description of this connection.
- `service_uuid` (String) UUID of the marketplace service being requested.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)


