package session

import "github.com/gorilla/sessions"

var store *sessions.CookieStore

func InitializeSessionManagement(secretKey string) {
	store = sessions.NewCookieStore([]byte(secretKey))
}

func GetSession() *sessions.CookieStore {
	return store
}
