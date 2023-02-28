package ua

import "net/http"

func GetUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}
