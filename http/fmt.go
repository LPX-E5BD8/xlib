package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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
	return string(b)
}
