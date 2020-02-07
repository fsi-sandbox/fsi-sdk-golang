package nibss

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

type Crypt struct {
	AESKey   []byte
	IVKey    []byte
	Code     string
	Password string
}

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

func Encode(t string) string {
	return base64.StdEncoding.EncodeToString([]byte(t))
}

func Decode(t string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(t)
}

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
