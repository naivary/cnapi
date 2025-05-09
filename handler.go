package main

import "net/http"

type FuncErr func(w http.ResponseWriter, r *http.Request) error

func toHandler(fn FuncErr) http.Handler {
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
