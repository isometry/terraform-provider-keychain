package keychain

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

const defaultKind = "terraform password"

var classAllowedValues = []string{"generic", "internet"}
var classLookup = map[string]keychain.SecClass{
	"generic":  keychain.SecClassGenericPassword,
	"internet": keychain.SecClassInternetPassword,
}

func readDataSourcePassword(class keychain.SecClass, service, username string) ([]byte, error) {
	query := keychain.NewItem()
	query.SetSecClass(class)
	query.SetSynchronizable(keychain.SynchronizableNo)
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

func createResourcePassword(class keychain.SecClass, kind, service, username string, password []byte) error {
	item := keychain.NewItem()
	item.SetSecClass(class)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetDescription(kind)
	item.SetService(service)
	item.SetAccount(username)
	item.SetData(password)

	return keychain.AddItem(item)
}

func readResourcePassword(class keychain.SecClass, kind, service, username string) ([]byte, error) {
	query := keychain.NewItem()
	query.SetSecClass(class)
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

func deleteResourcePassword(class keychain.SecClass, kind, service, username string) error {
	item := keychain.NewItem()
	item.SetSecClass(class)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetDescription(kind)
	item.SetService(service)
	item.SetAccount(username)

	return keychain.DeleteItem(item)
}
