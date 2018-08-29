package http

import "errors"

var ErrorEmptyResp = errors.New("empty response")
var ErrorEmptyRespBody = errors.New("empty request body")
var ErrorEmptyReq = errors.New("empty request")
