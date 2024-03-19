package spa

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

type SPAHandler struct {
	staticPath string
	indexFile  string
	fs         IFS
}

type IFS interface {
	Exists(path string) bool
	Read(path string) ([]byte, error)
}

func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	log.Println(path)
	indexFile := filepath.Join(h.staticPath, h.indexFile)

	content, err := h.fs.Read(path)
	if os.IsNotExist(err) || path == h.staticPath {
		index, err := h.fs.Read(indexFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, file := filepath.Split(path)
	mt := mime.TypeByExtension(filepath.Ext(file))
	w.Header().Set("Content-Type", mt)
	w.Write(content)
}

func NewSPA(staticPath string, index string, fs IFS) *SPAHandler {
	return &SPAHandler{
		staticPath: staticPath,
		indexFile:  index,
		fs:         fs,
	}
}
