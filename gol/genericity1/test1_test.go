package genericity1

import (
	"encoding/json"
	"testing"
)

func TestGen1_f(t *testing.T) {
	t.Log(gen1_f1("100"))
	t.Log(gen1_f1(100))
}

func gen1_f1[T string | int](s T) string {
	c, _ := json.Marshal(s)
	return string(c)
}
