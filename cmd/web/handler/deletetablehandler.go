package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hculpan/tableweaver/cmd/web/template"
	"github.com/hculpan/tableweaver/pkg/db"
	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/hculpan/tableweaver/pkg/session"
)

func DeleteTableHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetUsername(r)

	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		if name != "Select Table" {
			err := db.Delete(model.GetTableEntryDbKey(username, name))
			if err != nil {
				component := template.MessagePage(fmt.Sprintf("unable to find table: %s", err.Error()), true)
				component.Render(context.Background(), w)
				return
			}
		}

		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		tableEntries, err := model.GetAllTables(username)
		if err != nil {
			component := template.MessagePage(fmt.Sprintf("unable to retrieve table list: %s", err.Error()), true)
			component.Render(context.Background(), w)
			return
		}

		component := template.DeleteTablePage(tableEntries)
		component.Render(context.Background(), w)
	}
}
