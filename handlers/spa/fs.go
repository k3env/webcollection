package spa

import (
	"errors"
	"io/fs"
	"os"
)

type RealFS struct {
}

func (h *RealFS) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (h *RealFS) Exists(path string) bool {
	_, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}
