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

func AddTableHandler(w http.ResponseWriter, r *http.Request) {
	username := session.GetUsername(r)

	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		description := r.FormValue("description")
		url := r.FormValue("url")
		source := r.FormValue("source")

		var tableEntry *model.TableEntry
		if name != "" && url != "" {
			tableEntry = model.NewTableEntry(name, description, url, source)
			value, err := tableEntry.DbValue()
			if err != nil {
				component := template.MessagePage(fmt.Sprintf("unable to save table: %s", err.Error()), true)
				component.Render(context.Background(), w)
				return
			}

			err = db.SaveOrUpdate(model.GetTableEntryDbKey(username, tableEntry.Name), value)
			if err != nil {
				component := template.MessagePage(fmt.Sprintf("unable to save table: %s", err.Error()), true)
				component.Render(context.Background(), w)
				return
			}
		} else {
			component := template.MessagePage("required fields are blank", true)
			component.Render(context.Background(), w)
			return
		}

		component := template.AddTablePage(tableEntry)
		component.Render(context.Background(), w)
	} else {
		// Render login page
		component := template.AddTablePage(nil)
		component.Render(context.Background(), w)
	}
}
