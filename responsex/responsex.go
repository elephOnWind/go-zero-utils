package reponsex

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

const (
	DEFAULT_SUCCESS_CODE      = 0
	DEFAULT_ERROR_CODE        = 500
	DEFAULT_PARAMS_ERROR_CODE = 300
)

type Body struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp,string"`
}

func Success(w http.ResponseWriter, resp interface{}) {
	var body Body
	body.Code = DEFAULT_SUCCESS_CODE
	body.Msg = "Success"
	body.Data = resp
	httpx.OkJson(w, body)
}
func Error(w http.ResponseWriter, err error) {
	var body Body
	body.Code = DEFAULT_ERROR_CODE
	body.Msg = err.Error()
	body.Data = nil
	httpx.OkJson(w, body)
}
func Response(w http.ResponseWriter, code int, resp interface{}, err error) {
	var body Body
	body.Code = code
	if err != nil {
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
