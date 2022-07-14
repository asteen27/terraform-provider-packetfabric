---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cloud_services_aws_create_backbone_dedicated_cr Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-

---

# packetfabric_cloud_services_aws_create_backbone_dedicated_cr (Resource)



## Example Usage

```terraform
blank
```

## Schema

### Required

- `bandwidth` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--bandwidth))
- `description` (String) AWS Backbone Dedicated CR Description
- `epl` (Boolean) f true, created circuit will be an EPL otherwise EVPL
		EPL provides Point-to-Point connection between a pair of interfaces
		EVPL supports multiple Ethernet Virtual Connections per interface
- `interface_a` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--interface_a))
- `interface_z` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--interface_z))
- `rate_limit_in` (Number) The upper bound, in Mbps, to limit incoming data by.
- `rate_limit_out` (Number) The upper bound, in Mbps, to limit outgoing data by.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--bandwidth"></a>
### Nested Schema for `bandwidth`

Required:

- `account_uuid` (String) PacketFabric account UUID. The contact that will be billed.
- `longhaul_type` (String) Dedicated (no limits or additional charges), usage-based (per transfered GB) pricing model or hourly billing
		Enum ["dedicated" "usage" "hourly"]
- `subscription_term` (String) The billing term, in months, for this connection.
		Enum: ["1", "12", "24", "36"]


<a id="nestedblock--interface_a"></a>
### Nested Schema for `interface_a`

Required:

- `port_circuit_id` (String) The circuit ID of the customer's port.
- `untagged` (Boolean) Whether or not the interface should be untagged.
- `vlan` (Number) Valid VLAN range is from 4-4094, inclusive.


<a id="nestedblock--interface_z"></a>
### Nested Schema for `interface_z`

Required:

- `port_circuit_id` (String) The circuit ID of the customer's port.
- `untagged` (Boolean) Whether or not the interface should be untagged.
- `vlan` (Number) Valid VLAN range is from 4-4094, inclusive.