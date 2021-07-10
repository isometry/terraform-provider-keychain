package keychain

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKeychainPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeychainPasswordCreate,
		Read:   resourceKeychainPasswordRead,
		Delete: resourceKeychainPasswordDelete,

		Schema: map[string]*schema.Schema{
			"class": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "generic",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(classAllowedValues, false),
			},
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  defaultKind,
				ForceNew: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceKeychainPasswordCreate(d *schema.ResourceData, _ interface{}) error {
	class := d.Get("class").(string)
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password := []byte(d.Get("password").(string))
	err := createResourcePassword(classLookup[class], kind, service, username, password)
	if err != nil {
		return err
	}
	checksum := sha256.Sum256(password)
	d.SetId(hex.EncodeToString(checksum[:]))
	return nil
}

func resourceKeychainPasswordRead(d *schema.ResourceData, _ interface{}) error {
	class := d.Get("class").(string)
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password, err := readResourcePassword(classLookup[class], kind, service, username)
	if err != nil || password == nil {
		d.SetId("")
		return err
	}

	checksum := sha256.Sum256(password)
	if hex.EncodeToString(checksum[:]) != d.Id() {
		d.SetId("")
	}

	return nil
}

func resourceKeychainPasswordDelete(d *schema.ResourceData, _ interface{}) error {
	class := d.Get("class").(string)
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)

	return deleteResourcePassword(classLookup[class], kind, service, username)
}
