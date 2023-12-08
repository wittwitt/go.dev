package encoding1

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"testing"
	"time"
)

func Test_binary_f1(t *testing.T) {

	t.Log(uint64(time.Now().Unix()))

	list := []uint64{0x10, 0x7f, 0x80, 0x81, 0x99, 0xaa, 0xff, 0x1000, 0x8000, 838}
	// for i := 0; i < 100; i++ {
	for _, item := range list {

		buf := make([]byte, 8)
		binary.PutUvarint(buf, item)

		t.Log(item, hex.EncodeToString(buf))
	}

}

func Test_binary_f2(t *testing.T) {

	data, _ := hex.DecodeString("c6060155a0e402202a8883a8e6df3fec9d8391a081")

	r := bytes.NewBuffer(data)

	n, err := binary.ReadUvarint(r)
	t.Log(err)
	t.Log(n)

	n2, err := binary.ReadUvarint(r)
	t.Log(err)
	t.Log(n2)
}
