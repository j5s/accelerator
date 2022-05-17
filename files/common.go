package files

import (
	"os"
	"path/filepath"
)

func RemoveTempFiles() {
	dir, _ := filepath.Abs("temp")
	_ = os.RemoveAll(dir)
}
