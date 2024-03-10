package model

import (
	"encoding/json"
	"fmt"

	"github.com/hculpan/tableweaver/pkg/db"
)

type TableEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Source      string `json:"source"`
}

func NewTableEntry(name, description, url, source string) *TableEntry {
	return &TableEntry{
		Name:        name,
		Description: description,
		Url:         url,
		Source:      source,
	}
}

func (t *TableEntry) String() string {
	return fmt.Sprintf("%s: %s, %s, %s", t.Name, t.Description[:20], t.Source, t.Url)
}

func GetTableEntryDbKey(username, tableName string) string {
	return fmt.Sprintf("%s-table-%s", username, tableName)
}

func (t *TableEntry) DbValue() (string, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func GetAllTables(username string) ([]*TableEntry, error) {
	keys, err := db.SearchKeysWithSubstring(username + "-table-")
	if err != nil {
		return []*TableEntry{}, err
	}

	values, err := db.GetValuesForKeys(keys)
	if err != nil {
		return []*TableEntry{}, err
	}

	result := []*TableEntry{}
	for _, v := range values {
		entry, err := tableEntryFromDbValue(string(v))
		if err != nil {
			return result, err
		}
		result = append(result, entry)
	}

	return result, nil

}

func tableEntryFromDbValue(value string) (*TableEntry, error) {
	entry := TableEntry{}
	err := json.Unmarshal([]byte(value), &entry)
	return &entry, err
}

func (t *TableEntry) GetOnClick() string {
	return fmt.Sprintf("onclick=\"loadIntoIframe('%s');\"", t.Url)
}
