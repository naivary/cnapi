package main

import "net/http"

func addRoutes(mux *http.ServeMux) {
	baseChain := chain{adminOnly}
	mux.Handle("GET /healthz", baseChain.funcErr(healthz))
}
