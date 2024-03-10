package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/auth"
	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/hculpan/tableweaver/pkg/session"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := model.UserFromDb(username)
		if err != nil {
			slog.Error(fmt.Errorf("error logging in: %w", err).Error())
		}

		if strings.Trim(username, " ") == "" || !auth.CheckPasswordHash(password, user.PasswordHash) {
			component := template.LoginPage("invalid username or password, try again")
			component.Render(context.Background(), w)
			return
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
