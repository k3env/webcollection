package middlewares

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/urfave/negroni"
)

type LogRow struct {
	Prefix   string
	Request  *http.Request
	Response *LogRowResponse
}

type LogRowResponse struct {
	StatusCode    int
	ContentLength int
}

func (h *LogHandler) Handle(next http.Handler) http.Handler {
	return h.HandleFunc(next.ServeHTTP)
}

func (h *LogHandler) HandleFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)
		next(lrw, r)
		row := &LogRow{
			Prefix:  "Web",
			Request: r,
			Response: &LogRowResponse{
				StatusCode:    lrw.Status(),
				ContentLength: lrw.Size(),
			},
		}
		buff := bytes.NewBuffer([]byte(""))
		h.template.Execute(buff, row)
		log.Println(string(buff.Bytes()))
	})
}

type LogHandler struct {
	template *template.Template
}

func NewLogHandler(templateText string) (*LogHandler, error) {
	tpl, err := template.New("log").Parse(templateText)
	if err != nil {
		return nil, err
	}
	return &LogHandler{
		template: tpl,
	}, nil
}
