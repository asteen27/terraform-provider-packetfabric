package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGoogleProvision() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleProvisionCreate,
		ReadContext:   resourceGoogleProvisionRead,
		UpdateContext: resourceGoogleProvisionUpdate,
		DeleteContext: resourceGoogleProvisionDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: resourceProvision(),
	}
}

func resourceGoogleProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceProvisionCreate(ctx, d, m, c.CreateMktProvisionReq, googleProvider)
}

func resourceGoogleProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedRead(ctx, d, m)
}

func resourceGoogleProvisionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesHostedUpdate(ctx, d, m, c.UpdateServiceHostedConn)
}

func resourceGoogleProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "Google Service Delete")
}
