package aes1

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])
	return plainText[:(length - unpadding)]
}

func AesCBCEncrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("aes key: %w", err)
	}

	blockSize := block.BlockSize()
	// padding
	plainText = PKCS7Padding(plainText, blockSize)

	cipherText := make([]byte, blockSize+len(plainText))

	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte{}, fmt.Errorf("aes readfull: %w", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], plainText)

	return cipherText, nil
}

func AesCBCDncrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("aes key: %w", err)
	}

	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return []byte{}, fmt.Errorf("aes key: ciphertext too short: %d", blockSize)
	}
	iv := cipherText[:blockSize]
	if len(cipherText)%blockSize != 0 {
		return []byte{}, fmt.Errorf("aes key: ciphertext is not a multiple of the block size: %d", blockSize)
	}

	plainText := make([]byte, len(cipherText)-blockSize)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText[blockSize:])

	return PKCS7UnPadding(plainText), nil
}
