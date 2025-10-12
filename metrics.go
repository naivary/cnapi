package main

import (
	"net/http"

	"github.com/naivary/cnapi/openapi"
)

const (
	_required = true
	_optional = false
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
		RequestBody: openapi.NewReqBody("", false, new(CreateMetricRequest)),
		Responses: map[string]*openapi.Response{
			"200": openapi.NewResponse("", new(CreateMetricRequest)).
				AddHeader("", openapi.NewHeader("", _optional)).
				AddHeader("", openapi.NewHeader("", _optional)),
			"404": openapi.NewResponse("", nil),
		},
		Parameters: []*openapi.Parameter{
			openapi.NewQueryParam("id", "", _required),
		},
	}
}
