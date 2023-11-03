package math1

import (
	"encoding/binary"
	"math/bits"
	"testing"
)

func Test_t1(t *testing.T) {

	for i := 0; i < 10; i++ {

		ii := uint64(i)

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, ii)

		t.Logf("%b", buf)

		t.Log(bits.OnesCount64(ii)) //
	}
}
