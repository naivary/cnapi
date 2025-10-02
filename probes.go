package main

import (
	"net/http"
)

func readyz() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		err := r.Context().Err()
		if err != nil {
			return NewHTTPError("server shutting down", http.StatusServiceUnavailable)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return nil
	}
	return HandlerFuncErr(fn)
}

func livez() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return nil
	}
	return HandlerFuncErr(fn)
}
