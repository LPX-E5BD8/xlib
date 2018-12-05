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
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

// Bytes2String no copy to change bytes slice to string, unsafe
func Bytes2String(b []byte) (s string) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pStr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pStr.Data = pBytes.Data
	pStr.Len = pBytes.Len
	return
}

// String2Bytes no copy to change string to bytes slice, unsafe
func String2Bytes(s string) (b []byte) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pStr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pBytes.Data = pStr.Data
	pBytes.Len = pStr.Len
	pBytes.Cap = pStr.Len
	return
}

// RandomStrSeed is the seed of RandomString
const RandomStrSeed = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generate random string
func RandomString(length int) string {
	bytes := String2Bytes(RandomStrSeed)
	result := make([]byte, 0, length)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return Bytes2String(result)
}
