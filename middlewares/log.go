package middlewares

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)
		log.Printf("WEB: %s - %s%s (%s) - %d", r.Method, r.Host, r.URL.Path, r.Proto, lrw.Status())
	})
}

func LogFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)
		next(lrw, r)
		log.Printf("WEB: %s - %s%s (%s) - %d", r.Method, r.Host, r.URL.Path, r.Proto, lrw.Status())
	})
}
