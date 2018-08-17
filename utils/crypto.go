package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// Encrypt
func StringEncrypt(context string, salt []byte) (string, error) {
	c, err := aes.NewCipher(salt)
	if err != nil {
		return "", err
	}

	cryptText := make([]byte, len(context))
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cfb.XORKeyStream(cryptText, []byte(context))

	return base64.StdEncoding.EncodeToString(cryptText), nil
}

// Decrypt
func StringDecrypt(cryptBytes string, salt []byte) (string, error) {
	cryptTextByte, _ := base64.StdEncoding.DecodeString(cryptBytes)
	c, err := aes.NewCipher(salt)

	if err != nil {
		return "", err
	}

	var result []byte
	context := make([]byte, 500)

	cFBDec := cipher.NewCFBDecrypter(c, commonIV)
	cFBDec.XORKeyStream(context, cryptTextByte)

	for index, v := range context {
		if v != 0 {
			result = append(result, context[index])
			continue
		}
	}

	return string(result), nil
}
