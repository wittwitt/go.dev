package type1

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRef(t *testing.T) {

	f1 := new(big.Float)
	f1.SetInt64(100)

	f2 := big.NewFloat(1.3)

	f3 := big.NewFloat(1.6)

	f1.Add(f1, f2).Add(f1, f3)

	t.Log(fmt.Sprintf("%.8f", f1))
	t.Log(fmt.Sprintf("%.8f", f2))
	t.Log(fmt.Sprintf("%.8f", f3))

	f4 := big.NewFloat(10000)
	f5 := big.NewFloat(4)

	f4.Quo(f4, f5)
	// t.Log(fmt.Sprintf("%.8f", f6))

	t.Log(fmt.Sprintf("%.8f", f1))
	t.Log(fmt.Sprintf("%.8f", f2))
	t.Log(fmt.Sprintf("%.8f", f3))
	t.Log(fmt.Sprintf("%.8f", f4))
	t.Log(fmt.Sprintf("%.8f", f5))

}

func Test_float(t *testing.T) {
	var a int64 = 7
	var b int64 = 3

	c := a / b
	assert.Equal(t, int64(2), c)

	d := b / a
	assert.Equal(t, int64(0), d)
}
