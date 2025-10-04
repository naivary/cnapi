package main

import "net/http"

var ErrCtxDead = &HTTPError{
	StatusCode: http.StatusServiceUnavailable,
	Msg:        `Shutting Down. No new connections accepted`,
}
