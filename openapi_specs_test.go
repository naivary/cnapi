package main

import (
	"encoding/json"
	"io"
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
			r, err := GenOpenAPISpecs(tc.root, tc.h)
			if err != nil {
				t.Errorf("generating open API specs: %v", err)
				t.FailNow()
			}
			err = json.NewEncoder(os.Stdout).Encode(&tc.root)
			if err != nil {
				t.Errorf("json encode: %v", err)
				t.FailNow()
			}
			io.Copy(os.Stdout, r)
		})
	}
}
