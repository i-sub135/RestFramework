package response

import "net/http"

type output struct {
	Code      int         `json:"Code"`
	Message   string      `json:"Message"`
	ErrorCode int         `json:"ErrorCode"`
	ErrorMsg  string      `json:"ErrorMsg"`
	Result    interface{} `json:"Result"`
}

// RestResponse -- reciver Response
type RestResponse struct {
}

// RespOK -- response output OK
func (r *RestResponse) RespOK(data interface{}) output {
	out := output{
		Code:    http.StatusOK,
		Message: "OK",
		Result:  data,
	}

	return out
}

// RespBad -- response bad request
func (r *RestResponse) RespBad(errCode int, errMsg string) output {
	out := output{
		Code:      http.StatusBadRequest,
		Message:   "Bad Requests",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		Result:    map[string]interface{}{},
	}

	return out
}

// RespNotFound -- response 404
func (r *RestResponse) RespNotFound(errCode int, errMsg string) output {
	out := output{
		Code:      http.StatusNotFound,
		Message:   "Not Found",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		Result:    map[string]interface{}{},
	}
	return out
}
