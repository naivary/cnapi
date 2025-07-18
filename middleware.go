package main

import "net/http"

func timeout(next http.Handler) http.Handler {
	return nil
}

func requestID(next http.Handler) http.Handler {
	return nil
}

func log(next http.Handler) http.Handler {
	return nil
}

func cors(next http.Handler) http.Handler {
	return nil
}

func contentType(next http.Handler) http.Handler {
	return nil
}

func authz(next http.Handler) http.Handler {
	return nil
}

func authn(next http.Handler) http.Handler {
	return nil
}

// IsCanceled checks if the request context returns an error. If an error is
// returned the processed response by the handler before will not be sent.
func IsCanceled(next http.Handler) http.Handler {
	return nil
}
