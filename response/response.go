package response

import (
	"github.com/allen-ping/go-zero-response/v3/xerrors"
	"github.com/zeromicro/go-zero/core/metric"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

const serverNamespace = "http_server"

var metricBizCode = metric.NewCounterVec(&metric.CounterVecOpts{
	Namespace: serverNamespace,
	Subsystem: "requests",
	Name:      "biz_code_total",
	Help:      "business response code counter",
	Labels:    []string{"method", "path", "code"},
})

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
	metricBizCode.Inc(r.Method, r.URL.Path, strconv.Itoa(body.Ret.Code))

	httpx.OkJson(w, body)
}
