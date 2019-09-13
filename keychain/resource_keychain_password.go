package keychain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	keychain "github.com/keybase/go-keychain"
)

const (
	defaultKind = "application password"
)

func resourceKeychainPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeychainPasswordCreate,
		Read:   resourceKeychainPasswordRead,
		Delete: resourceKeychainPasswordDelete,

		Schema: map[string]*schema.Schema{
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  defaultKind,
				ForceNew: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": &schema.Schema{
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

func createPassword(kind, service, username string, password []byte) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetDescription(kind)
	item.SetService(service)
	item.SetAccount(username)
	item.SetData(password)

	return keychain.AddItem(item)
}

func readPassword(kind, service, username string) ([]byte, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetSynchronizable(keychain.SynchronizableNo)
	query.SetDescription(kind)
	query.SetService(service)
	query.SetAccount(username)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, err
	}
	switch hits := len(results); {
	case hits > 1:
		return nil, fmt.Errorf("Too many results")
	case hits == 1:
		return results[0].Data, nil
	default:
		return nil, nil
	}
}

func deletePassword(kind, service, username string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetDescription(kind)
	item.SetService(service)
	item.SetAccount(username)

	return keychain.DeleteItem(item)
}

func resourceKeychainPasswordCreate(d *schema.ResourceData, _ interface{}) error {
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password := []byte(d.Get("password").(string))
	err := createPassword(kind, service, username, password)
	if err != nil {
		return err
	}
	checksum := sha256.Sum256(password)
	d.SetId(hex.EncodeToString(checksum[:]))
	return nil
}

func resourceKeychainPasswordRead(d *schema.ResourceData, _ interface{}) error {
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	password, err := readPassword(kind, service, username)
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
	kind := d.Get("kind").(string)
	service := d.Get("service").(string)
	username := d.Get("username").(string)

	return deletePassword(kind, service, username)
}
