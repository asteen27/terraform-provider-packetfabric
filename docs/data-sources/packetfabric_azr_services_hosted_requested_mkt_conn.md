---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_azr_services_hosted_requested_mkt_conn Data Source - terraform-provider-packetfabric"
subcategory: ""
description: |-

---

# packetfabric_azr_services_hosted_requested_mkt_conn (Data Source)

## Example Usage

```terraform
blank
```

## Schema

### Read-Only

- `hosted_service_requests` (List of Object) (see [below for nested schema](#nestedatt--hosted_service_requests))
- `id` (String) The ID of this resource.

<a id="nestedatt--hosted_service_requests"></a>
### Nested Schema for `hosted_service_requests`

Read-Only:

- `allow_untagged_z` (Boolean)
- `bandwidth` (Set of Object) (see [below for nested schema](#nestedobjatt--hosted_service_requests--bandwidth))
- `from_customer` (Set of Object) (see [below for nested schema](#nestedobjatt--hosted_service_requests--from_customer))
- `request_type` (String)
- `status` (String)
- `text` (String)
- `time_created` (String)
- `time_updated` (String)
- `to_customer` (Set of Object) (see [below for nested schema](#nestedobjatt--hosted_service_requests--to_customer))
- `vc_mode` (String)
- `vc_request_uuid` (String)

<a id="nestedobjatt--hosted_service_requests--bandwidth"></a>
### Nested Schema for `hosted_service_requests.bandwidth`

Read-Only:

- `account_uuid` (String)
- `longhaul_type` (String)
- `speed` (String)
- `subscription_term` (Number)


<a id="nestedobjatt--hosted_service_requests--from_customer"></a>
### Nested Schema for `hosted_service_requests.from_customer`

Read-Only:

- `contact_email` (String)
- `contact_first_name` (String)
- `contact_last_name` (String)
- `contact_phone` (String)
- `customer_uuid` (String)
- `market` (String)
- `market_description` (String)
- `name` (String)


<a id="nestedobjatt--hosted_service_requests--to_customer"></a>
### Nested Schema for `hosted_service_requests.to_customer`

Read-Only:

- `customer_uuid` (String)
- `market` (String)
- `market_description` (String)
- `name` (String)