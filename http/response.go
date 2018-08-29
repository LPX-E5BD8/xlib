package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response

	// Err is the error doing the request.
	Err error
}

func (resp *Response) String() (string, error) {
	if resp == nil || resp.Response == nil {
		return "", ErrorEmptyResp
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

	return string(data), err
}

// JSON return JSONResult decode by
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

func (resp *Response) XML() {}
