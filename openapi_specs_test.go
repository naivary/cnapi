package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/naivary/cnapi/openapi"
)

func TestGenOpenAPISpec(t *testing.T) {
	tests := []struct {
		name string
		root *openapi.OpenAPI
		h    *Endpoint
	}{
		{
			name: "pattern param queries",
			root: openapi.New(&openapi.Info{Version: "", Title: ""}),
			h: &Endpoint{
				Pattern:     "GET /path/{p1}/{p2}",
				OperationID: "testFunc",
				Responses: map[string]*openapi.Response{
					"200": openapi.NewResponse("", new(CreateMetricRequest)),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := GenOpenAPISpecs(tc.root, tc.h)
			t.Log(err)
			jerr := json.NewEncoder(os.Stdout).Encode(&tc.root)
			t.Log(jerr)
		})
	}
}
