package qycms_error_code

const (
	// ErrSuccess - 200 ：OK，执行成功
	ErrSuccess int = iota + 100001

	// ErrUnknown - 500 ：Internal server error. 未知异常
	ErrUnknown

	// ErrBind - 400 ：Error occurred while binding the request body to the struct. 请求参数绑定结构体异常
	ErrBind

	// ErrValidation - 400 ：Validation failed. 数据校验异常
	ErrValidation

	// ErrTokenInvalid - 401 ：Token invalid. Token 校验异常
	ErrTokenInvalid

	// ErrPageNotFound - 404 ：Page not found. 路径请求异常
	ErrPageNotFound
)

const (
	// ErrDatabase - 500 ：Database error. 数据库异常
	ErrDatabase int = iota + 100101
)

const (
	ErrEncrypt int = iota + 100201
	ErrSignatureInvalid
	ErrExpired
	ErrInvalidAuthHeader

	ErrPasswordIncorrect
)
