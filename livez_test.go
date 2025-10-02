package main

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/naivary/cnapi/probe"
)

func TestLivez(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		code int
	}{
		{
			name: "liveness",
			code: http.StatusOK,
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr)
	r, err := http.NewRequest(http.MethodGet, "http://localhost:6443/readyz", nil)
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
	cl := &http.Client{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "http://localhost:6443/livez", nil)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
			_, err = cl.Do(r)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
		})
	}
}
