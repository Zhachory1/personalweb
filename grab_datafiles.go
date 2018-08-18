package datafiles

import (
	"log"
	"os"
	"path/filepath"
)

func GetDataFilesPath(filename string) string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
  topDir, _ := filepath.Split(filepath.Dir(ex))
	exPath := filepath.Join(topDir, filename)
	if _, err := os.Stat(exPath); err != nil {
		log.Fatal(err)
	}
	return exPath
}
