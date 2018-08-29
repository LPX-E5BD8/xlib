package test

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/liipx/xlib/http"
)

func TestRequest(t *testing.T) {
	res := make([]map[string]interface{}, 0)

	err := http.Get(
		"https://api.github.com/repos/vmg/redcarpet/issues?state=closed",
	).JSONUnmarshal(&res)

	if err != nil {
		t.Error(err)
	}

	pretty.Println(res)

}
