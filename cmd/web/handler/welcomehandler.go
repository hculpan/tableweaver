package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/hculpan/tableweaver/pkg/session"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetUsername(r)

	tableEntries, err := model.GetAllTables(username)
	if err != nil {
		component := template.MessagePage(fmt.Sprintf("unable to retrieve table list: %s", err.Error()), true)
		component.Render(context.Background(), w)
		return
	}

	component := template.WelcomePage(tableEntries)
	component.Render(context.Background(), w)
}
