package os1

import (
	"os"
	"testing"
	"time"
)

func TestOsLink(t *testing.T) {
	err := os.Symlink("1.txt", "some.txt") // some.txt不能存在
	t.Log(err)
}

func TestTmp(t *testing.T) {

	for {
		f, _ := os.CreateTemp("", "ssssss")
		t.Log(f.Name())

		f.Write([]byte("abc"))

		f.Close()

		time.Sleep(3 * time.Second)
	}
}
