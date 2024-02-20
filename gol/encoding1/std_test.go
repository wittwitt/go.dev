package encoding1

import (
	"testing"

	"github.com/multiformats/go-base32"
)

func Test_std_f1(t *testing.T) {

	encName := base32.RawStdEncoding.EncodeToString([]byte("wallet-"))

	t.Log(string(encName))
}
