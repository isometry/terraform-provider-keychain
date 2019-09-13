package keychain

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceKeychainPassword() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeychainPasswordRead,

		Schema: map[string]*schema.Schema{
			"class": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "generic",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(classAllowedValues, false),
			},
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
	class := d.Get("class").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password, err := readDataSourcePassword(classLookup[class], service, username)
	if err != nil {
		return err
	}
	d.Set("password", string(password))

	checksum := sha256.Sum256(password)
	d.SetId(hex.EncodeToString(checksum[:]))

	return nil
}
