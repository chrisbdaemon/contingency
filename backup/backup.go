package backup

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type Backup struct {
	Size int64
}

var sizeMap = map[byte]float64{
	'k': 1024.0,
	'm': 1024.0 * 1024.0,
	'g': 1024.0 * 1024.0 * 1024.0,
	't': 1024.0 * 1024.0 * 1024.0 * 1024,
	'p': 1024.0 * 1024.0 * 1024.0 * 1024 * 1024,
}

func fileList(path string) ([]string, error) {
	log.Debugf("Getting list of files at %s", path)
	var files []string

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})

	return files, err
}
