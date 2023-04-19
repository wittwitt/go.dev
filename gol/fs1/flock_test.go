package fs1

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_flock(t *testing.T) {

	for {
		isDLock := WriteLayerDLock()
		if !isDLock {
			log.Printf("wait!, other worker is writing!")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Println("no worker is writing, so do del action")
		time.Sleep(1 * time.Second)
	}
}

func WriteLayerDLock() bool {
	GOLANG_LOCK_FILE, isSet := os.LookupEnv("GOLANG_LOCK_FILE")
	if !isSet {
		return true
	}

	dlockDir := filepath.Dir(GOLANG_LOCK_FILE)
	_, err := os.Stat(dlockDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(dlockDir, 0775); err != nil {
				log.Printf("MkdirAll %s, %v", dlockDir, err)
				return false
			}
		} else {
			log.Printf("%s, %v", dlockDir, err)
			return false
		}
	}
	log.Printf("WriteLayerDLock, dlockDir: %s", dlockDir)

	list := []string{}
	filepath.WalkDir(dlockDir, func(path string, d fs.DirEntry, err error) error {
		if d != nil && !d.IsDir() {
			list = append(list, d.Name())
		}
		return nil
	})
	return len(list) == 0
}
