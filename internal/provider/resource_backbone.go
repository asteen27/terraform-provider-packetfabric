package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackbone() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Backbone Dedicated CR Description",
		},
		"bandwidth": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"account_uuid": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "PacketFabric account UUID. The contact that will be billed.",
					},
					"speed": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The desired speed of the new connection.\n\t\tEnum: []\"1gps\", \"10gbps\"]",
					},
					"subscription_term": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The billing term, in months, for this connection.\n\t\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
					},
					"longhaul_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Dedicated (no limits or additional charges), usage-based (per transfered GB) pricing model or hourly billing\n\t\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
					},
				},
			},
		},
		"interface_a": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port_circuit_id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The circuit ID of the customer's port.",
					},
					"vlan": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "Valid VLAN range is from 4-4094, inclusive.",
					},
					"untagged": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether or not the interface should be untagged.",
					},
				},
			},
		},
		"interface_z": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"port_circuit_id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The circuit ID of the customer's port.",
					},
					"vlan": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "Valid VLAN range is from 4-4094, inclusive.",
					},
					"untagged": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether or not the interface should be untagged.",
					},
				},
			},
		},
		"rate_limit_in": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The upper bound, in Mbps, to limit incoming data by.",
		},
		"rate_limit_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The upper bound, in Mbps, to limit outgoing data by.",
		},
		"epl": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "f true, created circuit will be an EPL otherwise EVPL\n\t\tEPL provides Point-to-Point connection between a pair of interfaces\n\t\tEVPL supports multiple Ethernet Virtual Connections per interface",
		},
	}
}

func resourceBackboneCreate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(packetfabric.Backbone) (*packetfabric.BackboneResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	awsBack := extractBack(d)
	resp, err := fn(awsBack)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.VcCircuitID)
	return diags
}

func resourceServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}, fn func() ([]packetfabric.AwsDedicatedConnResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConns, err := fn()
	if err != nil {
		return diag.FromErr(err)
	}
	for _, conn := range dedicatedConns {
		if conn.CloudCircuitID == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Services Read",
				Detail:   cloudCidNotFoundDetailsMsg,
			})
		} else {
			_ = d.Set("cloud_circuit_id", conn.CloudCircuitID)
		}
	}
	return diags
}

func resourceServicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(string, string) (*packetfabric.HostedConnectionResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cloudCID, ok := d.GetOk("id")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Cloud Service Create",
			Detail:   cloudCidNotFoundDetailsMsg,
		})
		return diags
	}
	desc, ok := d.GetOk("description")
	if !ok {
		return diag.Errorf("please provide a valid description for Cloud Service")
	}
	resp, err := fn(desc.(string), cloudCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("description", resp.Description)
	return diags
}

func resourceBackboneDelete(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(string) (*packetfabric.BackboneDeleteResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCircuitID, ok := d.GetOk("id"); ok {
		resp, err := c.DeleteBackbone(vcCircuitID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Backbone Delete result",
			Detail:   resp.Message,
		})
		return diags
	}
	return diag.Errorf("please provide a valid VC Circuit ID for deletion")
}

func extractBack(d *schema.ResourceData) packetfabric.Backbone {
	awsBack := packetfabric.Backbone{
		Description: d.Get("description").(string),
		Epl:         d.Get("epl").(bool),
	}
	for _, interfA := range d.Get("interface_a").(*schema.Set).List() {
		awsBack.Interfaces = append(awsBack.Interfaces, extractBackboneInterface(interfA.(map[string]interface{})))
	}
	for _, interfZ := range d.Get("interface_z").(*schema.Set).List() {
		awsBack.Interfaces = append(awsBack.Interfaces, extractBackboneInterface(interfZ.(map[string]interface{})))
	}
	for _, bw := range d.Get("bandwidth").(*schema.Set).List() {
		awsBack.Bandwidth = extractBandwidth(bw.(map[string]interface{}))
	}
	return awsBack
}

func extractBandwidth(bw map[string]interface{}) packetfabric.BackboneBandwidth {
	bandwidth := packetfabric.BackboneBandwidth{}
	bandwidth.AccountUUID = bw["account_uuid"].(string)
	bandwidth.SubscriptionTerm = bw["subscription_term"].(int)
	bandwidth.Speed = bw["speed"].(string)
	return bandwidth
}

func extractBackboneInterface(interf map[string]interface{}) packetfabric.BackBoneInterface {
	backboneInter := packetfabric.BackBoneInterface{}
	backboneInter.PortCircuitID = interf["port_circuit_id"].(string)
	backboneInter.Vlan = interf["vlan"].(int)
	backboneInter.Untagged = interf["untagged"].(bool)
	return backboneInter
}

func speedOptions() []string {
	return []string{
		"50Mbps", "100Mbps", "200Mbps", "300Mbps",
		"400Mbps", "500Mbps", "1Gbps", "2Gbps",
		"5Gbps", "10Gbps"}
}