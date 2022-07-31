package handlers

import "net/http"

func checkMethodAndPath(r *http.Request, path, method string) int {
	if r.Method != method {
		return http.StatusMethodNotAllowed
	}
	if r.URL.Path != path {
		return http.StatusNotFound
	}
	return 200
}
