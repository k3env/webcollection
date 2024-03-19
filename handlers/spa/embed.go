package spa

import (
	"embed"
	"errors"
	"io/fs"
)

type EmbedFS struct {
	fs *embed.FS
}

func (f *EmbedFS) Read(path string) ([]byte, error) {
	return f.fs.ReadFile(path)
}

func (f *EmbedFS) Exists(path string) bool {
	_, err := f.fs.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func NewEmbedFS(fs *embed.FS) *EmbedFS {
	return &EmbedFS{fs: fs}
}
