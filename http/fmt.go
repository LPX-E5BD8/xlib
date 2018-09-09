package http

import (
	"fmt"
	"io"
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
func bodyPretty(body io.ReadCloser) string {
	if body == nil {
		return ""
	}

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return ""
	}

	return string(bytes)
}
