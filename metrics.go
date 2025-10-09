package main

import (
	"net/http"

	"github.com/naivary/cnapi/openapi"
)

const _regParam = `\{([^\/{}]+)\}`

type CreateMetricRequest struct{}

func metrics() http.Handler {
	fn := HandlerFuncErr(func(w http.ResponseWriter, r *http.Request) error {
		return nil
	})
	return &Endpoint{
		Handler:     fn,
		Error:       defaultErrorHandler(),
		Pattern:     "GET /metrics/{definition}",
		Summary:     "Get the metrics of the documetnation",
		Description: "<go doc string of the method if empty>",
		Tags:        []string{"metrics"},
		OperationID: "<name of the method>",
		RequestBody: CreateMetricRequest{},
		Responses: map[string]*openapi.Response{
			"200": {
				Summary:     "Successfull Response", // default,
				Description: "Successfull Response", // default,
				Headers:     map[string]*openapi.Header{},
			},
		},
	}
}
