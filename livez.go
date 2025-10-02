package main

import "net/http"

func livez() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return nil
	}
	return HandlerFuncErr(fn)
}
