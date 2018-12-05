/*
Copyright 2018 liipx(lipengxiang)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// AESEncryptCFB Encrypt string with CFB
func AESEncryptCFB(context string, salt []byte) (string, error) {
	c, err := aes.NewCipher(salt)
	if err != nil {
		return "", err
	}

	cryptText := make([]byte, len(context))
	cfb := cipher.NewCFBEncrypter(c, iv)
	cfb.XORKeyStream(cryptText, String2Bytes(context))

	return base64.StdEncoding.EncodeToString(cryptText), nil
}

// AESDecryptCFB Decrypt string with CFB
func AESDecryptCFB(cryptBytes string, salt []byte) (string, error) {
	cryptTextByte, _ := base64.StdEncoding.DecodeString(cryptBytes)
	c, err := aes.NewCipher(salt)

	if err != nil {
		return "", err
	}

	var result []byte
	context := make([]byte, 500)

	cFBDec := cipher.NewCFBDecrypter(c, iv)
	cFBDec.XORKeyStream(context, cryptTextByte)

	for index, v := range context {
		if v != 0 {
			result = append(result, context[index])
			continue
		}
	}

	return Bytes2String(result), nil
}
