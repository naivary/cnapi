package main

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/naivary/cnapi/probe"
)

var errProbeFailed = errors.New("probe failed")

func NewTestServer(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) (*http.Client, error) {
	go run(ctx, args, getenv, stdin, stdout, stderr)
	r, err := http.NewRequest(http.MethodGet, "http://localhost:9443/readyz", nil)
	if err != nil {
		return nil, err
	}
	status, err := probe.DoHTTP(r, 5*time.Second)
	if err != nil {
		return nil, err
	}
	if status == probe.Failed {
		return nil, errProbeFailed
	}
	cl := &http.Client{}
	return cl, nil
}

func randPort() (int, error) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return -1, err
	}
	defer lis.Close()
	return lis.Addr().(*net.TCPAddr).Port, nil
}
