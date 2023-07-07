/*
Provide a resource to manage a domain authentication validation.
Example Usage
```hcl

	resource "sendgrid_domain_authentication_validation" "foo" {
		domain_authentication_id = sendgrid_domain_authentication.foo.id
	}

```
*/
package sendgrid

import (
	"context"

	sendgrid "github.com/anna-money/terraform-provider-sendgrid/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSendgridDomainAuthenticationValidation() *schema.Resource { //nolint:funlen
	return &schema.Resource{
		CreateContext: resourceSendgridDomainAuthenticationValidationCreate,
		ReadContext:   resourceSendgridDomainAuthenticationValidationRead,
		UpdateContext: resourceSendgridDomainAuthenticationValidationUpdate,
		DeleteContext: resourceSendgridDomainAuthenticationValidationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"domain_authentication_id": {
				Type:        schema.TypeInt,
				Description: "Id of the domain authentication to validate.",
				Required:    true,
			},

			"valid": {
				Type:        schema.TypeBool,
				Description: "Indicates if this is a valid authenticated domain or not.",
				Computed:    true,
			},
		},
	}
}

func resourceSendgridDomainAuthenticationValidationCreate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	return validateDomain(ctx, d, m)
}

func validateDomain(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*sendgrid.Client)

	if err := c.ValidateDomainAuthentication(ctx, d.Get("domain_authentication_id").(string)); err.Err != nil || err.StatusCode != 200 {
		if err.Err != nil {
			return diag.FromErr(err.Err)
		}
		return diag.Errorf("unable to validate domain DNS configuration")
	}

	return resourceSendgridDomainAuthenticationValidationRead(ctx, d, m)
}

func resourceSendgridDomainAuthenticationValidationRead( //nolint:funlen,cyclop
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	c := m.(*sendgrid.Client)

	auth, err := c.ReadDomainAuthentication(ctx, d.Get("domain_authentication_id").(string))
	if err.Err != nil {
		return diag.FromErr(err.Err)
	}

	//nolint:errcheck
	d.Set("valid", auth.Valid)
	return nil
}

func resourceSendgridDomainAuthenticationValidationUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	return validateDomain(ctx, d, m)
}

func resourceSendgridDomainAuthenticationValidationDelete(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	return nil
}
