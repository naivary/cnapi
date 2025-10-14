package main

import (
	"net/http"

	"github.com/naivary/cnapi/openapi"
)

const (
	_required = true
	_optional = false
)

// +openapi:schema:title="create metric response model"
type CreateMetricResponse struct{}

// +openapi:schema:title="create metric request model"
type CreateMetricRequest struct {
	// +openapi:schema:required
	Name string
}

func metrics() http.Handler {
	fn := HandlerFuncErr(func(w http.ResponseWriter, r *http.Request) error {
		return nil
	})
	return &Endpoint{
		Handler:     fn,
		Error:       defaultErrorHandler(),
		Pattern:     "GET /metrics/{definition}",
		Summary:     "Get the metrics of the documetnation",
		Description: "something",
		Tags:        []string{"metrics"},
		OperationID: "<name of the method if empty>",
		RequestBody: openapi.NewReqBody("", false, new(CreateMetricRequest)),
		Responses: map[string]*openapi.Response{
			"200": openapi.NewResponse("", new(CreateMetricRequest)).
				AddHeader("", openapi.NewHeader("", _optional)).
				AddHeader("", openapi.NewHeader("", _optional)),
			"404": openapi.NewResponse("", nil),
		},
		Parameters: []*openapi.Parameter{
			openapi.NewPathParam("definition", openapi.StringParameter),
			openapi.NewQueryParam("id", "", _required),
		},
	}
}
