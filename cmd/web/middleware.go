package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// No SURF Add CSRF Protection to ALl POST Request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Session Load session load and save on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
