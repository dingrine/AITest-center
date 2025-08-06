package api

type Response[T any] struct {
	// 错误内容
	Error string `json:"error,omitempty"`
	// 错误描述
	Message string `json:"message,omitempty"`
	// 返回的数据
	Data *T `json:"data,omitempty"`
}

func NewResponse[T any](err string, msg string, data *T) *Response[T] {
	return &Response[T]{
		Error:   err,
		Message: msg,
		Data:    data,
	}
}

func NewErrorResponse[T any](err string, msg string) *Response[T] {
	var v T
	return NewResponse[T](err, msg, &v)
}

func NewSuccessResponse[T any](data *T) *Response[T] {
	return NewResponse[T]("", "", data)
}
