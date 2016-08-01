package strcat

import "net/http"

// Handler 用来处理字符串拼接
type Handler interface {
	http.Handler
}

// New 生成一个新的 Handler
func New() Handler {
	return &strcat{}
}
