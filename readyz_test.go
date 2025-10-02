package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReadyz(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		code int
	}{
		{
			name: "readiness",
			r:    httptest.NewRequest(http.MethodGet, "/readyz", nil),
			code: http.StatusOK,
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}
