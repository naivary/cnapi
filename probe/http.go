package probe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func DoHTTPWithClient(r *http.Request, client *http.Client, timeout time.Duration) (Status, error) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	req := r.Clone(ctx)
	for {
		res, err := client.Do(req)
		if err != nil {
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return Failed, fmt.Errorf("request timed out after %v: %w", timeout, err)
			}
			return Failed, err
		}
		res.Body.Close()
		if isSuccessful(res.StatusCode) {
			return Success, nil
		}
		if ctx.Err() != nil {
			return Failed, ctx.Err()
		}
		time.Sleep(100 * time.Millisecond)
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
