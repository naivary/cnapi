package probe

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func newMux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	mux.Handle("/timeout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	return mux
}

func TestDoHTTP(t *testing.T) {
	server := httptest.NewServer(newMux())
	defer server.Close()
	tests := []struct {
		name    string
		target  string
		timeout time.Duration
		status  Status
	}{
		{
			name:    "success",
			target:  "/healthz",
			timeout: 5 * time.Second,
			status:  Success,
		},
		{
			name:    "timeout",
			target:  "/timeout",
			timeout: 1 * time.Second,
			status:  Failed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u, err := url.JoinPath(server.URL, tc.target)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
			req, err := http.NewRequest(http.MethodGet, u, nil)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				t.FailNow()
			}
			status, _ := DoHTTPWithClient(req, server.Client(), tc.timeout)
			if status != tc.status {
				t.Errorf("status not expected. got: %d; wanted: %d", status, tc.status)
			}
		})
	}
}
