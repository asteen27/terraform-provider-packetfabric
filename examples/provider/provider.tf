terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.4"
    }
  }
}

provider "packetfabric" {}
