package response

import (
	"github.com/allen-ping/go-zero-response/v3/xerrors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Body interface{} `json:"body,omitempty"`
	Ret  *Result     `json:"ret"`
}
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	body := &Body{
		Body: resp,
		Ret: &Result{
			Code: 0,
			Msg:  "success",
		},
	}

	e := xerrors.FromError(err)
	if e != nil {
		body.Ret.Code = e.Code
		body.Ret.Msg = e.Msg
	}

	httpx.OkJson(w, body)
}
