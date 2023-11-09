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

func Test_t2(t *testing.T) {
	for i := 0; i < 100; i++ {

		ii := uint64(i)

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, ii)

		t.Logf("%d, %b, %d, %d", i, buf, bits.LeadingZeros64(ii), bits.TrailingZeros64(ii))

	}
}

func Test_t3(t *testing.T) {

	len_out := 4023

	chunks := len_out / 127

	t.Log("=== unpadReader", len_out, chunks)

	outTwoPow := 1 << (63 - bits.LeadingZeros64(uint64(chunks*128)))

	if outTwoPow < 128 {
		t.Log("minimum padded piece size is 128 bytes")
		return
	}
	if bits.OnesCount64(uint64(outTwoPow)) != 1 {
		t.Log("padded piece size must be a power of 2")
		return
	}
	t.Log("xxx1", outTwoPow)

	todo := outTwoPow

	ss2KiB := uint64(2 << 10)

	if ss2KiB < uint64(todo) {
		todo = 1 << (63 - bits.LeadingZeros64(ss2KiB))
	}

	t.Log(todo)
}
