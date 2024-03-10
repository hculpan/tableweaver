package handler

import (
	"context"
	"net/http"

	"github.com/hculpan/tableweaver/cmd/web/template"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	errorValue := r.URL.Query().Get("error")
	component := template.MessagePage(message, errorValue == "yes")
	component.Render(context.Background(), w)
}
