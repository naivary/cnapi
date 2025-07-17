package main

import "net/http"

func Timeout(next http.Handler) http.Handler {
	return nil
}

func RequestID(next http.Handler) http.Handler {
	return nil
}

func Log(next http.Handler) http.Handler {
	return nil
}

func CORS(next http.Handler) http.Handler {
	return nil
}

func ContentType(next http.Handler) http.Handler {
	return nil
}

func Authz(next http.Handler) http.Handler {
	return nil
}

func Authn(next http.Handler) http.Handler {
	return nil
}
