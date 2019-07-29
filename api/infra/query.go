package infra

import (
	"net/http"
	"strconv"
)

//GetLimit limit
func GetLimit(r *http.Request) (int, error) {
	if r.URL.Query().Get("limit") != "" {
		return strconv.Atoi(r.URL.Query().Get("limit"))
	}
	return 0, nil
}

//GetPage page
func GetPage(r *http.Request) (int, error) {
	if r.URL.Query().Get("page") != "" {
		return strconv.Atoi(r.URL.Query().Get("page"))
	}
	return 0, nil
}
