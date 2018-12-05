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

package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/liipx/xlib/util"
)

// headerPretty could sort keys of header and fmt
func headerPretty(header http.Header) string {
	var fmtStr string

	var keys []string
	for k := range header {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmtStr += fmt.Sprintf("\n%30s: %s", k, header.Get(k))
	}

	return fmtStr
}

// bodyPretty could string the body
func bodyPretty(resp *Response) string {
	if resp.Response == nil || resp.Body == nil {
		return ""
	}

	// ReadAll will close resp.Body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return util.Bytes2String(b)
}
