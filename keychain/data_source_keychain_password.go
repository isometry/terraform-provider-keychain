package keychain

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform/helper/schema"
	keychain "github.com/keybase/go-keychain"
)

func dataSourceKeychainPassword() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeychainPasswordRead,

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceKeychainPasswordRead(d *schema.ResourceData, _ interface{}) error {
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password, err := keychain.GetGenericPassword(service, username, "", "")
	if err != nil {
		return err
	}
	d.Set("password", string(password))

	checksum := sha256.Sum256([]byte(password))
	d.SetId(hex.EncodeToString(checksum[:]))

	return nil
}
