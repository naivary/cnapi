package main

import "net/http"

type Endpoint struct {
	Handler HandlerFuncErr
	Error   HTTPErrorHandler
}

func (e Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := e.Handler(w, r)
	if err == nil {
		return
	}
	e.Error.ServeError(w, r, err)
}

type HTTPErrorHandler interface {
	ServeError(w http.ResponseWriter, r *http.Request, err error)
}

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func (e ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	e(w, r, err)
}

func defaultErrorHandler() HTTPErrorHandler {
	fn := func(w http.ResponseWriter, r *http.Request, err error) {
		httpErr, isHTTPErr := err.(*HTTPError)
		msg := err.Error()
		code := http.StatusInternalServerError
		if isHTTPErr {
			msg = httpErr.Msg
			code = httpErr.StatusCode
		}
		http.Error(w, msg, code)
	}
	return ErrorHandlerFunc(fn)
}

type HandlerFuncErr func(w http.ResponseWriter, r *http.Request) error

func (h HandlerFuncErr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err == nil {
		return
	}
	defaultErrorHandler().ServeError(w, r, err)
}
