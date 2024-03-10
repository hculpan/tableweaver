package handler

import (
	"context"
	"net/http"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/session"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session.Logout(w, r)

	component := template.LogoutPage()
	component.Render(context.Background(), w)
}
