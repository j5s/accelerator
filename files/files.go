package files

import (
	"os"
	"path/filepath"
)

func ReadAllClasses() []string {
	var data []string
	path, _ := filepath.Abs("temp")
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data = append(data, path)
		return nil
	})
	return data
}
