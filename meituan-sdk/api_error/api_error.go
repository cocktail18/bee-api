package api_error

import "fmt"

type ApiError struct {
	Code    string
	Msg     string
	TraceId string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("code=%s, msg=%s, traceId=%s", e.Code, e.Msg, e.TraceId)
}
