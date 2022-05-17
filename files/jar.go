package files

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func unzipJar(path string, id string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("error jar path: %s", path))
	}
	r, err := zip.OpenReader(absPath)
	if r == nil {
		panic(fmt.Errorf("cannot read file: %s", absPath))
	}
	for _, f := range r.File {
		tempPath := filepath.Join("temp", id, f.Name)
		if strings.HasSuffix(f.Name, "/") {
			_ = os.MkdirAll(tempPath, 0644)
		} else {
			if !strings.HasSuffix(f.Name, ".class") {
				continue
			}
			reader, _ := f.Open()
			data, _ := ioutil.ReadAll(reader)
			_ = ioutil.WriteFile(tempPath, data, 0644)
		}
	}
}

func UnzipJars(dir string) {
	dirPath, err := filepath.Abs(dir)
	if err != nil {
		panic(fmt.Errorf("error dir path: %s", dir))
	}
	fileList, _ := ioutil.ReadDir(dirPath)
	for i, f := range fileList {
		if strings.HasSuffix(f.Name(), ".jar") {
			finalPath := filepath.Join(dirPath, f.Name())
			id := strconv.Itoa(i) + "_" + f.Name()
			unzipJar(finalPath, id)
		}
	}
}
