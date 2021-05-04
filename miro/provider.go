package miro

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *scema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Description: "Access key for Miro API",
				Required:    true,
				Swnsitive:   true,
			},
		},
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(ctx context.Context, data *schema.ResorceData) (interface{}, diag.Diagnostics) {
	key := data.Get("access_token").(string)
	var diags diag.Diagnostics
	return miro.NewClient(key), diags
}
