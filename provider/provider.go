package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/socrate-tech/terraform-provider-zitadeltools/provider/jwt"
)

func ZitadelToolsProvider() *schema.Provider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"zitadeltools_jwt": jwt.GetResource(),
		},
	}

	return provider

}
