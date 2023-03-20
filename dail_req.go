package main

import (
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

func (r *Requester) DoDailRequest(req *fasthttp.Request, resp *fasthttp.Response, rr *ReportRecord) {
	t1 := time.Since(startTime)
	var err error
	addr := strings.Split(r.clientOpt.url, "//")[1]
	conn, err := fasthttp.Dial(addr)
	//conn, err := r.httpClient.Dial(addr)
	//if r.clientOpt.doTimeout > 0 {
	//	err = r.httpClient.DoTimeout(req, resp, r.clientOpt.doTimeout)
	//} else {
	//	err = r.httpClient.Do(req, resp)
	//}
	var code string

	if err != nil {
		rr.cost = time.Since(startTime) - t1
		rr.code = ""
		rr.error = err.Error()
		return
	}
	conn.Close()
	code = "2xx"

	rr.cost = time.Since(startTime) - t1
	rr.code = code
	rr.error = ""
}
