package controllers

import beego "github.com/beego/beego/v2/server/web"

type AdminBaseController struct {
	beego.Controller
}

// api return result
type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//api return error msg
func ErrorMsg(msg string, code ...int) Result {
	var r Result
	if len(code) > 0 {
		r.Code = code[0]
	} else {
		r.Code = 1
	}
	r.Msg = msg
	r.Data = nil
	return r
}

//api return err data
func ErrorData(msg error, code... int) Result {
	var r Result
	if len(code) > 0 {
		r.Code = code[0]
	}else{
		r.Code = 1
	}
	r.Msg = msg.Error()
	r.Data = nil
	return r
}

//api return success data
func SuccessData(data interface{}) Result {
	var r Result
	r.Code = 0
	r.Msg = "ok"
	r.Data = data
	return r
}