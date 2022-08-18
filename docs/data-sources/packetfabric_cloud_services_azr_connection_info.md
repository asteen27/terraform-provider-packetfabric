---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cloud_services_azr_connection_info Data Source - terraform-provider-packetfabric"
subcategory: ""
description: |-

---

# packetfabric_cloud_services_azr_connection_info (Data Source)

## Example Usage

```terraform
blank
```

## Schema

### Required

- `cloud_circuit_id` (String) The unique PF circuit ID for this connection
		Example: PF-AP-LAX1-1002

### Optional

- `cloud_provider_pop` (String) Point of Presence for the cloud provider location.
		Example: DAL1
- `cloud_provider_region` (String) Region short name.
		Example: us-west-1
- `customer_uuid` (String) The UUID for the customer this connection belongs to.
- `description` (String) The description of this connection.
		Example: AWS connection for Foo Corp.
- `pop` (String) Point of Presence for the connection.
		Example: LAS1
- `port_type` (String) The port type for the given port.
		Enum: [ "hosted", "dedicated" ]
- `service_class` (String) The service class for the given port, either long haul or metro.
		Enum: [ "longhaul", "metro" ]
- `service_provider` (String) The service provider of the connection
		Enum: [ "aws", "azure", "packet", "google", "ibm", "salesforce", "webex" ]
- `site` (String) Site name
		Example: SwitchNAP Las Vegas 7
- `speed` (String) The desired speed of the connection.
		Enum: [ "50Mbps", "100Mbps", "200Mbps", "300Mbps", "400Mbps", "500Mbps", "1Gbps", "2Gbps", "5Gbps", "10Gbps" ]
- `state` (String) The state of the connection.
		Enum: [ "active", "deleting", "inactive", "pending", "requested" ]
- `time_created` (String) Date and time of connection creation
- `time_updated` (String) Date and time connection was last updated
- `user_uuid` (String) The UUID for the user this connection belongs to.

### Read-Only

- `id` (String) The ID of this resource.