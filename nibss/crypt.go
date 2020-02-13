// Package nibss implements access to nibss sandbox.
//
// It futher implements bvnr, fingerprint and BVNPlaceholder sandbox APIs
// as exported functions.
package nibss

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Crypt struct required for Encrypt and Decrypt Func.
type Crypt struct {
	AESKey   []byte
	IVKey    []byte
	Code     string
	Password string
}

// Encrypt converts byte slice to an AES-256 encrypted byte slice.
// It returns byte slice and any encrypt error encountered.
func (c Crypt) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.AESKey)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	mode := cipher.NewCBCEncrypter(block, c.IVKey)
	cipherText := make([]byte, len(origData))
	mode.CryptBlocks(cipherText, origData)
	return cipherText, nil
}

// Decrypt converts an AES-256 encrypted string to byte slice.
// It returns byte slice and any decrypt error encountered.
func (c Crypt) Decrypt(encryptedText string) ([]byte, error) {
	cipherText, _ := hex.DecodeString(encryptedText)
	block, err := aes.NewCipher(c.AESKey)

	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, c.IVKey)
	origData := make([]byte, len(cipherText))
	mode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// Encode converts a plain string to base64-encoded string.
func Encode(t string) string {
	return base64.StdEncoding.EncodeToString([]byte(t))
}

// Decode converts a base64-encoded string back to human-readable
// string.
// It returns byte slice and any decoding error encountered.
func Decode(t string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(t)
}

// Sha256 generates SHA256 Hash of a plain string.
func Sha256(t string) string {
	h := sha256.New()
	h.Write([]byte(t))
	return hex.EncodeToString(h.Sum(nil))
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
