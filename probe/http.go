package probe

import (
	"context"
	"errors"
	"net/http"
	"syscall"
	"time"
)

func DoHTTPWithClient(r *http.Request, client *http.Client, timeout time.Duration) (Status, error) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	req := r.Clone(ctx)
	for {
		res, err := client.Do(req)
		if errors.Is(err, syscall.ECONNREFUSED) && ctx.Err() == nil {
			time.Sleep(250 * time.Millisecond)
			continue
		}
		if err := ctx.Err(); err != nil {
			return Failed, err
		}
		res.Body.Close()
		if isSuccessful(res.StatusCode) {
			return Success, nil
		}
	}
}

func DoHTTP(r *http.Request, timeout time.Duration) (Status, error) {
	cl := defaultClient()
	return DoHTTPWithClient(r, cl, timeout)
}

// isSuccessful returns true if the code is any of the 2XX codes
func isSuccessful(code int) bool {
	return code >= http.StatusOK && code <= http.StatusIMUsed
}

// TODO: FOLLOW redirects
func defaultClient() *http.Client {
	return &http.Client{}
}
