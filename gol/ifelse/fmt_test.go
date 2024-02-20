package ifelse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fmt_t1(t *testing.T) {

	str := "@3523xxx335vxzcv"

	var num, num2 int64
	var str2 string

	n, err := fmt.Sscanf(str, "@%dxxx%d%s", &num, &num2, &str2)
	require.NoError(t, err)
	t.Log(n)

	t.Log(num)
	t.Log(num2)
	t.Log(str2)
}
