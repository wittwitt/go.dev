package os1

import (
	"os"
	"testing"
)

func TestOsLink(t *testing.T) {
	err := os.Symlink("1.txt", "some.txt") // some.txt不能存在
	t.Log(err)
}
