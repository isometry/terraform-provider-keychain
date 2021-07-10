package keychain

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"keychain_password": resourceKeychainPassword(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"keychain_password": dataSourceKeychainPassword(),
		},
	}
}
