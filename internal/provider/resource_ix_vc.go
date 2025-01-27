package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIxVC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIxVCCreate,
		ReadContext:   resourceThirdPartyVirtualCircuitRead,
		DeleteContext: resourceIXVCDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The routing ID of the IX provider that will be receiving this request.\n\n\tExample: TR-1RI-OQ85",
			},
			"market": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The market code (e.g. \"ATL\" or \"DAL\") in which you would like the IX provider to provision their side of the connection.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A brief description of this connection.",
			},
			"asn": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Your ASN.",
			},
			"rate_limit_in": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The upper bound, in Mbps, to limit incoming data by.",
			},
			"rate_limit_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The upper bound, in Mbps, to limit outgoing data by.",
			},
			"bandwidth": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_uuid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
							ValidateFunc: validation.IsUUID,
							Description: "The UUID for the billing account that should be billed. " +
								"Can also be set with the PF_ACCOUNT_ID environment variable.",
						},
						"speed": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice(ixVcSpeedOptions(), true),
							Description:  "The desired speed of the new connection. Only applicable if `longhaul_type` is \"dedicated\" or \"hourly\".\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\" \"20Gbps\" \"30Gbps\" \"40Gbps\" \"50Gbps\" \"60Gbps\" \"80Gbps\" \"100Gbps\"]",
						},
						"subscription_term": {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
							Description:  "The billing term, in months, for this connection. Only applicable if `longhaul_type` is \"dedicated.\"\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
						},
						"longhaul_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"dedicated", "usage", "hourly"}, true),
							Description:  "Dedicated (no limits or additional charges), usage-based (per transferred GB) or hourly billing.\n\n\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
						},
					},
				},
			},
			"interface": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_circuit_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The circuit ID for the port. This starts with \"PF-AP-\"",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "Valid VLAN range is from 4-4094, inclusive.",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the interface should be untagged.",
						},
					},
				},
			},
			"flex_bandwidth_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the flex bandwidth container from which to subtract this VC's speed.",
			},
		},
	}
}

func resourceIxVCCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ixVC := extractIXVC(d)
	resp, err := c.CreateIXVirtualCircuit(ixVC)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.VcRequestUUID)
	return diags
}

func resourceIXVCDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.DeleteVCRequest(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func extractIXVC(d *schema.ResourceData) packetfabric.IxVirtualCircuit {
	ixVC := packetfabric.IxVirtualCircuit{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		ixVC.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		ixVC.Market = market.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		ixVC.Description = description.(string)
	}
	if asn, ok := d.GetOk("asn"); ok {
		ixVC.Asn = asn.(int)
	}
	if rateLimitIn, ok := d.GetOk("rate_limit_in"); ok {
		ixVC.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk("rate_limit_out"); ok {
		ixVC.RateLimitOut = rateLimitOut.(int)
	}
	for _, bw := range d.Get("bandwidth").(*schema.Set).List() {
		ixVC.Bandwidth = extractBandwidth(bw.(map[string]interface{}))
	}
	for _, interf := range d.Get("interface").(*schema.Set).List() {
		ixVC.Interface = extractIXVcInterface(interf.(map[string]interface{}))
	}
	if flexBandID, ok := d.GetOk("flex_bandwidth_id"); ok {
		ixVC.FlexBandwidthID = flexBandID.(string)
	}
	return ixVC
}

func extractIXVcInterface(interf map[string]interface{}) packetfabric.Interfaces {
	vxInterf := packetfabric.Interfaces{}
	if portCID := interf["port_circuit_id"]; portCID != nil {
		vxInterf.PortCircuitID = portCID.(string)
	}
	if vlan := interf["vlan"]; vlan != nil {
		vxInterf.Vlan = vlan.(int)
	}
	if untagged := interf["untagged"]; untagged != nil {
		vxInterf.Untagged = untagged.(bool)
	}
	if svlan := interf["svlan"]; svlan != nil {
		vxInterf.Svlan = svlan.(int)
	}
	return vxInterf
}

func ixVcSpeedOptions() []string {
	return []string{
		"50Mbps", "100Mbps", "200Mbps", "300Mbps",
		"400Mbps", "500Mbps", "1Gbps", "2Gbps",
		"5Gbps", "10Gbps", "20Gbps", "30Gbps",
		"40Gbps", "50Gbps", "60Gbps", "80Gbps",
		"100Gbps"}
}
