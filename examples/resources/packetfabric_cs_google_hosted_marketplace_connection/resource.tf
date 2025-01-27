resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_conn1_marketplace_google" {
  provider                    = packetfabric
  description                 = var.pf_description
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  routing_id                  = var.pf_routing_id
  market                      = var.pf_market
  speed                       = var.pf_cs_speed
  pop                         = var.pf_cs_pop

}

output "packetfabric_cs_google_hosted_marketplace_connection" {
  value     = packetfabric_cs_google_hosted_marketplace_connection.cs_conn1_marketplace_google
  sensitive = true
}