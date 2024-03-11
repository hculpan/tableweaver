package handler

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/hculpan/tableweaver/pkg/session"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetUsername(r)

	tableEntries, err := model.GetAllTables(username)
	sort.Slice(tableEntries, func(i, j int) bool {
		return tableEntries[i].Name < tableEntries[j].Name
	})

	if err != nil {
		component := template.MessagePage(fmt.Sprintf("unable to retrieve table list: %s", err.Error()), true)
		component.Render(context.Background(), w)
		return
	}

	component := template.WelcomePage(tableEntries)
	component.Render(context.Background(), w)
}
