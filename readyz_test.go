package main

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestReadyz(t *testing.T) {
	tests := []struct {
		name string
		code int
	}{
		{
			name: "readiness",
			code: http.StatusOK,
		},
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	baseURL, err := NewTestServer(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	endpoint, err := url.JoinPath(baseURL, "readyz")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	cl := &http.Client{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, endpoint, nil)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
			res, err := cl.Do(r)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
			if res.StatusCode != tc.code {
				t.Errorf("status code differ. Got: %d; Want: %d", res.StatusCode, tc.code)
				t.FailNow()
			}
		})
	}
}
