package responsex

import (
	"net/http"
	"time"

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

func OkJson(w http.ResponseWriter, resp interface{}) {
	Response(w, DEFAULT_SUCCESS_CODE, resp, nil)
}
func OK(w http.ResponseWriter) {
	Response(w, DEFAULT_SUCCESS_CODE, nil, nil)
}
func Error(w http.ResponseWriter, err error) {
	Response(w, DEFAULT_ERROR_CODE, nil, err)
}
func Response(w http.ResponseWriter, code int, resp interface{}, err error) {
	var body Body
	body.Code = code
	if err != nil {
		body.Msg = err.Error()
	} else {
		body.Msg = "Success"
		body.Data = resp
	}
	body.Timestamp = timestamp()
	httpx.OkJson(w, body)
}
func timestamp() int64 {
	return time.Now().UnixNano() / 1000
}
