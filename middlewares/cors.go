package middlewares

import (
	"net/http"
	"strings"
)

type Origin string
type Method string

const (
	MethodGet     Method = "GET"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodDelete  Method = "DELETE"
	MethodHead    Method = "HEAD"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
	MethodConnect Method = "CONNECT"
)

const (
	OriginAll  Origin = "*"
	OriginSame Origin = "same"
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

type CorsMiddleware struct {
	origin  Origin
	methods []Method
	headers []string
}

func NewCorsMiddleware(origin Origin, methods []Method, headers []string) *CorsMiddleware {
	return &CorsMiddleware{
		origin:  origin,
		methods: methods,
		headers: headers,
	}
}

func (m *CorsMiddleware) HandleFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var originVal string
		var methodVals []string

		if m.origin == OriginSame {
			originVal = r.Host
		} else {
			originVal = string(m.origin)
		}

		if len(m.methods) == 0 {
			methodVals = append(methodVals, "*")
		} else {
			for _, v := range m.methods {
				methodVals = append(methodVals, string(v))
			}
		}

		w.Header().Set("Access-Control-Allow-Origin", originVal)
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(methodVals, ","))
		if len(m.headers) == 0 {
			w.Header().Set("Access-Control-Allow-Headers", "*")
		} else {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(m.headers, ","))
		}
		next(w, r)
	})
}

func (m *CorsMiddleware) Handle(next http.Handler) http.Handler {
	return m.HandleFunc(next.ServeHTTP)
}
