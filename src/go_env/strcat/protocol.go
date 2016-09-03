// Package strcat 使用 HTTP 实现了字符串拼接的服务。
package strcat

import (
	"expvar"
	"net/http"
)

// Handler 用来处理字符串拼接
type Handler interface {
	http.Handler
}

// New 生成一个新的 Handler
func New() Handler {
	count := expvar.Get("count")
	return &strcat{
		count: count.(*expvar.Int),
	}
}
