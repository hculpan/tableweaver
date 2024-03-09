package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/session"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		// password := r.FormValue("password")
		// Process login
		// On success:

		if strings.Trim(username, " ") == "" {
			component := template.LoginPage("invalid username or password, try again")
			component.Render(context.Background(), w)
		}

		session, _ := session.GetSession().Get(r, "session-name")
		session.Values["user_id"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		// Render login page
		component := template.LoginPage("")
		component.Render(context.Background(), w)
	}
}
