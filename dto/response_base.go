package dto

type HttpBaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SuccessResp(data ...any) HttpBaseResponse {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return HttpBaseResponse{
		Code: 0,
		Msg:  "success",
		Data: d,
	}
}

func ErrorResp(code int, msg ...string) HttpBaseResponse {
	var m string
	if len(msg) > 0 {
		m = msg[0]
	} else {
		m = "error"
	}
	return HttpBaseResponse{
		Code: code,
		Msg:  m,
	}
}
