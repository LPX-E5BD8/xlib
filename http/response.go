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
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/liipx/xlib/util"
)

// Response base on http.Response
type Response struct {
	*http.Response

	// Err is the error doing the request.
	Err error
}

func (resp *Response) debug() {
	if !Debug || resp.Response == nil {
		return
	}

	log.Debug("Response header:", headerPretty(resp.Header))
	log.Debug("Response body:\n", bodyPretty(resp))
}

// String return response body content in string type
func (resp *Response) String() (string, error) {
	if resp == nil || resp.Response == nil {
		return "", errorEmptyResp
	}

	if resp.Err != nil {
		return "", resp.Err
	}

	// get data from resp
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return util.Bytes2String(data), err
}

// JSONUnmarshal decode response body content to v
func (resp *Response) JSONUnmarshal(v interface{}) error {
	if resp == nil {
		return errors.New("empty response")
	}

	if resp.Err != nil {
		return resp.Err
	}

	// get data from resp
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	err = json.Unmarshal(data, &v)

	if err != nil {
		return err
	}

	return nil
}

// XML todo
func (resp *Response) XML() {}
