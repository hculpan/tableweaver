package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitializeSessionManagement(secretKey string) {
	store = sessions.NewCookieStore([]byte(secretKey))
}

func GetSession() *sessions.CookieStore {
	return store
}

func GetUsername(r *http.Request) string {
	session, _ := GetSession().Get(r, "session-name")
	if userID, ok := session.Values["user_id"].(string); ok && userID != "" {
		return userID
	}

	return ""
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	delete(session.Values, "user_id")
	return session.Save(r, w)
}
