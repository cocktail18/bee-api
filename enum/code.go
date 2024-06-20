package enum

type ResCode int32

const (
	ResCodeSuccess      ResCode = 0
	ResCodeFail         ResCode = 1
	ResCodeCreated      ResCode = 201
	ResCodeUnauthorized ResCode = 401
	ResCodeForbidden    ResCode = 403
	ResCodeNotFound     ResCode = 404
	ResCodeEmpty        ResCode = 700

	ResCodeTokenInvalid ResCode = 2000
)
