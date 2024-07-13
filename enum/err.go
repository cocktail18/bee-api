package enum

import "github.com/pkg/errors"

type BussError struct {
	Err     error
	Code    ResCode
	Message string
}

func (err *BussError) Error() string {
	return err.Message
}

func NewBussErr(err error, code ResCode, message string) *BussError {
	if err == nil {
		err = ErrBuss
	}
	return &BussError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

var (
	ErrBuss         = errors.New("业务错误")
	ErrEmpty        = NewBussErr(ErrBuss, ResCodeEmpty, "数据为空")
	ErrNotImplement = NewBussErr(ErrBuss, ResCodeFail, "未实现")
	ErrParamError   = NewBussErr(ErrBuss, ResCodeForbidden, "参数错误")
	ErrReceived     = NewBussErr(ErrBuss, 30002, "已经领取过了")
	ErrIneligible   = NewBussErr(ErrBuss, 30003, "不符合条件")
)
