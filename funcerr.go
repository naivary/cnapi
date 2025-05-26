package main

import "net/http"

type funcErr func(w http.ResponseWriter, r *http.Request) error

func toHandler(fn funcErr) http.Handler {
	hl := func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err == nil {
			return
		}
		httpErr, isHTTPErr := err.(*HTTPError)
		msg := err.Error()
		code := http.StatusInternalServerError
		if isHTTPErr {
			msg = httpErr.Msg
			code = httpErr.StatusCode
		}
		http.Error(w, msg, code)
	}
	return http.HandlerFunc(hl)
}
