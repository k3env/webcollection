package middlewares

import (
	"net/http"
)

type Origin string

const (
	OriginAll  = "*"
	OriginSame = "same"
)

func Cors(next http.Handler, origin Origin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var originVal string
		if origin == OriginSame {
			originVal = r.Host
		} else {
			originVal = string(origin)
		}

		w.Header().Set("Access-Control-Allow-Origin", originVal)
		next.ServeHTTP(w, r)
	})
}
