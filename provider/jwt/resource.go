package jwt

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a JWT token converted from a private key thanks to Zitadel tools.",
		Schema: map[string]*schema.Schema{
			"key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
			},
			"audience": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jwt": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
		},
		CreateContext: create,
		ReadContext:   get,
		UpdateContext: update,
		DeleteContext: delete,
	}
}
