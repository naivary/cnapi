package main

import "net/http"

func healthz(w http.ResponseWriter, r *http.Request) error {
	return nil
}
