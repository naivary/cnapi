package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/naivary/cnapi/probe"
)

func TestReadyz(t *testing.T) {
	getenv := func(s string) string {
		if s == "port" {
			return "7443"
		}
		return os.Getenv(s)
	}

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
	go run(ctx, os.Args, getenv, os.Stdin, os.Stdout, os.Stderr)
	r, err := http.NewRequest(http.MethodGet, "http://localhost:7443/livez", nil)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	status, err := probe.DoHTTP(r, 5*time.Second)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	if status == probe.Failed {
		t.Errorf("probe failed.")
		t.FailNow()
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}
