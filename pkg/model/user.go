package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hculpan/tableweaver/pkg/auth"
	"github.com/hculpan/tableweaver/pkg/db"
)

type User struct {
	Username     string    `json:"username"`
	RealName     string    `json:"realname"`
	PasswordHash string    `json:"password"`
	Registered   time.Time `json:"registered"`
	Validated    time.Time `json:"validated"`
}

func New(username, realname, password string, isRegistered, isValidated bool) (*User, error) {
	result := User{
		Username: strings.ToLower(username),
		RealName: realname,
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return &result, err
	}
	result.PasswordHash = hash

	if isRegistered {
		result.Registered = time.Now()
	}

	if isValidated {
		result.Validated = time.Now()
	}

	return &result, nil
}

func (u *User) String() string {
	result := fmt.Sprintf("%s, %q, %s", u.Username, u.RealName, u.PasswordHash)
	if !u.Registered.IsZero() {
		result += fmt.Sprintf(", Registered at: %s", u.Registered.String())
	}
	if !u.Validated.IsZero() {
		result += fmt.Sprintf(", Validated at: %s", u.Validated.String())
	}
	return result
}

func (u *User) DbValue() (string, error) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (u *User) DbKey() string {
	return GetDbKey(u.Username)
}

func GetDbKey(username string) string {
	return username + "-userrecord"
}

func (u *User) Save() error {
	v, err := u.DbValue()
	if err != nil {
		return err
	}
	return db.SaveOrUpdate(u.DbKey(), v)
}

func GetAllUsers() ([]*User, error) {
	userKeys, err := db.SearchKeysWithSubstring("-userrecord")
	if err != nil {
		return []*User{}, err
	}

	userValues, err := db.GetValuesForKeys(userKeys)
	if err != nil {
		return []*User{}, err
	}

	result := []*User{}
	for _, v := range userValues {
		user := User{}
		err := json.Unmarshal([]byte(v), &user)
		if err != nil {
			return result, err
		}
		result = append(result, &user)
	}

	return result, nil
}
