# Use Case: PacketFabric Cloud Router with AWS and Google

This use case builds a PacketFabric Cloud Router between AWS and Google Cloud Platform.
Terraform providers used: PacketFabric, AWS and Google. This example uses AWS Transit VIF & Gateway.

![Deployment Diagram](./images/diagram_cloud_router_aws_google.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [HashiCorp Google Terraform Provider](https://registry.terraform.io/providers/hashicorp/google)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

## Terraform resources deployed

- "aws_dx_gateway"
- "aws_dx_transit_virtual_interface"
- "aws_dx_gateway_association"
- "aws_security_group"
- "aws_network_interface"
- "aws_key_pair"
- "aws_instance"
- "aws_eip"
- "aws_ec2_transit_gateway"
- "aws_ec2_transit_gateway_vpc_attachment"
- "aws_vpc"
- "aws_subnet"
- "aws_internet_gateway"
- "aws_route_table_association"
- "packetfabric_cloud_router"
- "packetfabric_cloud_router_connection_aws"
- "time_sleep"
- "aws_dx_connection_confirmation"
- "google_compute_router"
- "google_compute_interconnect_attachment"
- "packetfabric_cloud_router_connection_google"
- "packetfabric_cloud_router_bgp_session"
- "google_compute_firewall"
- "google_compute_instance"
- "google_compute_network"
- "google_compute_subnetwork"
- "random_pet"

**Estimated time:** ~10 min for Google, AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Transit Gateway

**Note**: Because the BGP session is created automatically, we use gcloud terraform module to retreive the BGP addresses and set the PacketFabric Cloud Router ASN in the BGP settings in the Google Cloud Router. Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/11458), [vote](https://github.com/hashicorp/terraform-provider-google/issues/12624) and [vote](https://github.com/hashicorp/terraform-provider-google/issues/12630) for these issues on GitHub.

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
- Don't have an Google Account? [Get Started](https://cloud.google.com/free)

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- [gcloud](https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest)
- [jq](https://stedolan.github.io/jq/download/)

Make sure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
- [Google Service Account](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)

## Quick Start

1. Create the file ``secret.tfvars`` and update each variables as needed (edit ``variables.tf``).

```sh
cp secret.tfvars.sample secret.tfvars
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan:

```sh
terraform apply -var-file="secret.tfvars"
```

3. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy -var-file="secret.tfvars"
```

**Note:** Default login/password for Locust is ``demo:packetfabric`` edit ``user-data-ubuntu.sh`` script to change it.

## Troubleshooting

1. In case you get the following error:

```
╷
│ Error: Error when reading or editing InterconnectAttachment: googleapi: Error 400: The resource 'projects/prefab-setting-123456/regions/us-west1/interconnectAttachments/demo-pf-gcp-vpn-master-cricket' is not ready, resourceNotReady
│ 
│ 
```

This seems to be a problem with Google Terraform Provider, run again the terraform destroy command and the destroy will complete correctly the 2nd try.
Please [vote](https://github.com/hashicorp/terraform-provider-google/issues/12631) for this issue on GitHub.

2. In case the ``gcloud_bgp_address`` module fails, check the error, fix it and manually remove the state before re-running the terraform config.

```sh
terraform state rm module.gcloud_bgp_addresses
terraform state rm module.gcloud_bgp_peer_update
```

3. In case you get the following error:

```
╷
│ Error: error waiting for Direct Connection Connection (dxcon-fgohxwui) confirm: timeout while waiting for state to become 'available' (last state: 'pending', timeout: 10m0s)
│ 
│   with aws_dx_connection_confirmation.confirmation_1,
│   on cloud_router_connection_aws.tf line 46, in resource "aws_dx_connection_confirmation" "confirmation_1":
│   46: resource "aws_dx_connection_confirmation" "confirmation_1" {
│ 
```

You are hitting a timeout issue in AWS [aws_dx_connection_confirmation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/dx_connection_confirmation) resource. Please [vote](https://github.com/hashicorp/terraform-provider-aws/issues/26335) for this issue on GitHub.

As a workaround, edit the `cloud_router_connection_aws.tf` and comment out the following resource:

```
# resource "aws_dx_connection_confirmation" "confirmation_1" {
#   provider      = aws
#   connection_id = data.aws_dx_connection.current_1.id
# }
```

Edit the `aws_dx_transit_vif.tf` and comment out the dependency with `confirmation_1` in `packetfabric_cloud_router_connection_aws` data source: 

```
data "packetfabric_cloud_router_connections" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    aws_dx_connection_confirmation.confirmation_1,
    # aws_dx_connection_confirmation.confirmation_2,
  ]
}
```

Then remove the `confirmation_2` state, check the Direct Connect connection is **available** and re-apply the terraform plan:
```
terraform state rm aws_dx_connection_confirmation.confirmation_1
terraform apply -var-file="secret.tfvars"
```