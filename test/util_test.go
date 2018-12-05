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

package test

import (
	"testing"

	"github.com/liipx/xlib/util"
)

func TestUtil_Crypto(t *testing.T) {
	count := 10000
	s := "testData"
	randSalt := make([]string, 0)
	for i := 0; i < count; i++ {
		randSalt = append(randSalt, util.RandomString(16))
	}

	var err error
	var res string
	for _, saltStr := range randSalt {
		salt := util.String2Bytes(saltStr)
		res, err = util.AESEncryptCFB(s, salt)
		if err != nil {
			t.Error(err)
			return
		}

		res, err = util.AESDecryptCFB(res, salt)
		if err != nil {
			t.Error(err)
			return
		}

		if res != s {
			t.Error("want", s, "got", res)
			return
		}
	}
}
