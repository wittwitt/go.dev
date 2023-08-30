package rand1

import (
	"math/rand"
	"testing"
)

func Test_rand1(t *testing.T) {
	pr := rand.Perm(10)
	t.Log(pr)
}
