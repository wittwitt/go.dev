package xor

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func XorMain() {
	k, c, err := EncodeXor("hu#hi2wt3262h627152@hh Y2RVXB6E2VGH5IBX")
	fmt.Println(k, c, err)

	fmt.Println(DecodeXor(k, c))
}

// EncodeXor EncodeXor
func EncodeXor(msg string) (keyStr string, ciphertext string, err error) {
	//string to bytes
	strBuf := []byte(msg)

	keyBuf, err := GenerateRandomBytes(len(strBuf))
	if err != nil {
		return
	}
	keyStr = base64.URLEncoding.EncodeToString(keyBuf)

	var outBuf = make([]byte, len(strBuf))
	for i := 0; i < len(strBuf); i++ {
		outBuf[i] = strBuf[i] ^ keyBuf[i]
	}
	ciphertext = base64.URLEncoding.EncodeToString(outBuf)

	return
}

// DecodeXor DecodeXor
func DecodeXor(key string, str string) (msg string, err error) {
	keyBuf, err := base64.URLEncoding.DecodeString(key)
	if err != nil {
		fmt.Println("decode: decode key err", err)
		return
	}

	strBuf, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("decode: decode str err", err)
		return
	}

	if len(keyBuf) != len(strBuf) {
		fmt.Println("decode: decode err")
		return
	}

	var outBuf = make([]byte, len(strBuf))
	for i := 0; i < len(strBuf); i++ {
		outBuf[i] = strBuf[i] ^ keyBuf[i]
	}

	msg = string(outBuf)

	return
}

// GenerateRandomBytes GenerateRandomBytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
