package util

import "net/http"

func Redirect(redirectPath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectPath, http.StatusFound)
	}
}
