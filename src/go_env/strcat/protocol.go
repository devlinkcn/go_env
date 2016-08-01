package strcat

import "net/http"

type Handler interface {
	http.Handler
}

func New() Handler {
	return &strcat{}
}
