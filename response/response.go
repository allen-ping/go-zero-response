package response

import (
	"github.com/allen-ping/go-zero-response/v2/xerrors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	body := &Body{
		Code:    0,
		Message: "OK",
		Data:    resp,
	}

	e := xerrors.FromError(err)
	if e != nil {
		body.Code = e.Code
		body.Message = e.Msg
	}

	httpx.OkJson(w, body)
}
