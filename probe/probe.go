package probe

import (
	"net/http"
	"time"
)

type Status int

const (
	Success = iota + 1
	Failed
)

func DoHTTPWithClient(r *http.Request, client *http.Client, timeout time.Duration) (Status, error) {
	start := time.Now()
	for now := time.Now(); start.Add(timeout).Before(now); {
		res, err := client.Do(r)
		if err != nil {
			return Failed, err
		}
		if isSuccessful(res.StatusCode) {
			return Success, nil
		}
	}
	return Failed, nil
}

func DoHTTP(r *http.Request, timeout time.Duration) (Status, error) {
	cl := defaultClient()
	return DoHTTPWithClient(r, cl, timeout)
}

// isSuccessCode returns true if the code is any of the 200 codes
func isSuccessful(code int) bool {
	return code >= http.StatusOK && code <= http.StatusIMUsed
}

func defaultClient() *http.Client {
	return &http.Client{}
}
