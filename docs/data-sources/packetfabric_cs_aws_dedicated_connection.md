---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cs_aws_dedicated_connection_conn Data Source - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_aws_dedicated_connection_conn (Data Source)



## Data Example

```terraform
{
  "dedicated_connections" = tolist([
    {
      "account_uuid" = "bbbfb3fe-cdd1-48a9-90ea-9fc59ea41234"
      "cloud_circuit_id" = "PF-AE-PDX1-1739482"
      "cloud_provider" = toset([
        {
          "pop" = "PDX1"
          "site" = "Pittock Building"
        },
      ])
      "customer_uuid" = "58c80946-5fbc-400e-8060-95b5dfbf1234"
      "deleted" = false
      "description" = "AWS Dedicated connection for Foo update"
      "is_cloud_router_connection" = false
      "pop" = "PDX1"
      "port_type" = "dedicated"
      "service_class" = "metro"
      "service_provider" = "aws"
      "settings" = toset([
        {
          "autoneg" = false
          "aws_region" = "us-west-2"
          "zone_dest" = "B"
        },
      ])
      "settings_aws_region" = ""
      "site" = "Pittock Building"
      "speed" = "10Gbps"
      "state" = "active"
      "time_created" = "2022-06-16T23:13:21.126145-0700"
      "time_updated" = "2022-06-16T23:15:09.089127-0700"
      "user_uuid" = "4e3bb859-9f64-4d12-ae9c-be3a0231234"
      "uuid" = "3adadf96-3c27-4598-baf8-f4d993401234"
    },
  ])
  "id" = "f9a76a87-e7d0-44b9-a35e-9a89b2241234"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `dedicated_connections` (List of Object) (see [below for nested schema](#nestedatt--dedicated_connections))
- `id` (String) The ID of this resource.

<a id="nestedatt--dedicated_connections"></a>
### Nested Schema for `dedicated_connections`

Read-Only:

- `account_uuid` (String)
- `cloud_circuit_id` (String)
- `cloud_provider` (Set of Object) (see [below for nested schema](#nestedobjatt--dedicated_connections--cloud_provider))
- `customer_uuid` (String)
- `deleted` (Boolean)
- `description` (String)
- `is_cloud_router_connection` (Boolean)
- `pop` (String)
- `port_type` (String)
- `service_class` (String)
- `service_provider` (String)
- `settings` (Set of Object) (see [below for nested schema](#nestedobjatt--dedicated_connections--settings))
- `settings_aws_region` (String)
- `site` (String)
- `speed` (String)
- `state` (String)
- `time_created` (String)
- `time_updated` (String)
- `user_uuid` (String)
- `uuid` (String)

<a id="nestedobjatt--dedicated_connections--cloud_provider"></a>
### Nested Schema for `dedicated_connections.cloud_provider`

Read-Only:

- `pop` (String)
- `site` (String)


<a id="nestedobjatt--dedicated_connections--settings"></a>
### Nested Schema for `dedicated_connections.settings`

Read-Only:

- `autoneg` (Boolean)
- `aws_region` (String)
- `zone_dest` (String)

