package jwt

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a JWT token converted from a private key thanks to Zitadel tools.",
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Description: "The private key used to generate the JWT token.",
				Sensitive:   true,
				Required:    true,
			},
			"audience": {
				Type:        schema.TypeString,
				Description: "The audience of the JWT token.",
				Required:    true,
			},
			"jwt": {
				Type: schema.TypeString,
				// Sensitive: true,
				Description: "The JWT token generated from the key and the audience.",
				Computed:    true,
			},
		},
		CreateContext: create,
		ReadContext:   get,
		UpdateContext: update,
		DeleteContext: delete,
	}
}
