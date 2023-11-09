package io1

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_t1(t *testing.T) {

	r1 := bytes.NewBuffer([]byte{0x11, 0x22, 0x33, 0x44, 0x55})
	r2 := bytes.NewBuffer([]byte{0x01, 0x02, 0x03, 0x04, 0x05})

	r := io.MultiReader(r1, r2)

	data, err := io.ReadAll(r)
	require.NoError(t, err)

	fmt.Println(hex.EncodeToString(data))
}
