package response

import (
	"github.com/allen-ping/go-zero-response/xerrors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	body := &Body{
		Code: 0,
		Msg:  "OK",
		Data: resp,
	}

	e := xerrors.FromError(err)
	if err != nil {
		body.Code = e.Code
		body.Msg = e.Msg
	}

	httpx.OkJson(w, body)
}
