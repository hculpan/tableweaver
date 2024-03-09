package db

import (
	"strings"

	"github.com/dgraph-io/badger/v3"
)

var DB *badger.DB

func InitDB(path string) error {
	var err error
	DB, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return err
	}

	return nil
}

func SaveOrUpdate(key string, value string) error {
	err := DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	return err
}

func SearchKeysWithSubstring(substring string) ([]string, error) {
	var matchedKeys []string

	err := DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false // Only fetch keys
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())

			if containsSubstring(key, substring) {
				matchedKeys = append(matchedKeys, key)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchedKeys, nil
}

func containsSubstring(key, substring string) bool {
	return strings.Contains(key, substring)
}

func GetValuesForKeys(keys []string) (map[string][]byte, error) {
	valuesMap := make(map[string][]byte)

	err := DB.View(func(txn *badger.Txn) error {
		for _, key := range keys {
			item, err := txn.Get([]byte(key))
			if err != nil {
				if err == badger.ErrKeyNotFound {
					continue // Skip keys that are not found
				}
				return err // Return other errors
			}
			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			valuesMap[key] = val
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return valuesMap, nil
}

func Delete(key string) error {
	err := DB.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	return err
}

func GetValue(key string) ([]byte, error) {
	var value []byte
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}

func DeleteAllKeysWithSubstring(substring string) error {
	keys, err := SearchKeysWithSubstring(substring)
	if err != nil {
		return err
	}

	for _, key := range keys {
		err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
