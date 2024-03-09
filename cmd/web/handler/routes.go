package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", authorized(log(WelcomeHandler)))
	mux.HandleFunc("/login", log(LoginHandler))
	mux.HandleFunc("/welcome", log(authorized(WelcomeHandler)))
	//	http.HandleFunc("/logout", handler.LogoutHandler)

}

func authorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("checking authorization on " + r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("page %s retrieved", r.URL.String()))
		next.ServeHTTP(w, r)
	})
}
