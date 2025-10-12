package main

import (
	"net/http"

	"github.com/naivary/cnapi/openapi"
)

const _regParam = `\{([^\/{}]+)\}`

type (
	CreateMetricRequest  struct{}
	CreateMetricResponse struct{}
)

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
		RequestBody: openapi.Request("", false, new(CreateMetricRequest)),
		Responses: map[string]*openapi.Response{
			"200": openapi.Res("", "", new(CreateMetricResponse),
				openapi.NewHeader("", true, openapi.STRING).Deprecate(),
			),
		},
	}
}
