package test

import (
	"testing"
	"io/ioutil"

	"github.com/liipx/xlib/http"
)

func TestRequest(t *testing.T) {
	resp := http.Get("http://www.baidu.com", nil)
	if resp.Err != nil {
		t.Error(resp.Err)
	} else {
		defer resp.Body.Close()
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
	}
}
