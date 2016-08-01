package main

import (
	"net/http"

	"go_env/strcat"
)

func main() {
	config := loadConfig()
	handler := strcat.New()
	http.Handle("/strcat", handler)
	http.ListenAndServe(config.Listen, nil)
}
