package handler

import (
	"context"
	"net/http"

	"github.com/hculpan/tableweaver/cmd/web/template"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	username := "harry"
	ok := true
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	component := template.WelcomePage(username)
	component.Render(context.Background(), w)
}
