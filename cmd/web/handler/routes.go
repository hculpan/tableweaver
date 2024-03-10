package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hculpan/tableweaver/pkg/session"
)

func Routes(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", authorized(log(WelcomeHandler)))
	mux.HandleFunc("/login", log(LoginHandler))
	mux.HandleFunc("/welcome", log(authorized(WelcomeHandler)))
	mux.HandleFunc("/logout", log(authorized(LogoutHandler)))
	mux.HandleFunc("/addtable", log(authorized(AddTableHandler)))
}

func authorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("checking authorization on " + r.URL.String())
		if userId := session.GetUsername(r); userId != "" {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	})
}

func log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("page %s retrieved", r.URL.String()))
		next.ServeHTTP(w, r)
	})
}
