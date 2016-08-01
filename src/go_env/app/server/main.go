package main

import (
	_ "expvar"
	"net/http"
	_ "net/http/pprof"

	"go_env/strcat"
)

func main() {
	config := loadConfig()
	handler := strcat.New()
	http.Handle("/strcat", handler)
	http.ListenAndServe(config.Listen, nil)
}
