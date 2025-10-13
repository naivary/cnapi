package openapi

import "reflect"

type RequestBody struct {
	Description string                `json:"description,omitempty"`
	Required    bool                  `json:"required"`
	Content     map[string]*MediaType `json:"content,omitempty"`
}

func NewReqBody(desc string, required bool, model any) *RequestBody {
	req := &RequestBody{
		Description: desc,
		Required:    required,
		Content: map[string]*MediaType{
			"application/json": &MediaType{
				Schema: &Schema{Ref: reflect.TypeOf(model).Name()},
			},
		},
	}
	return req
}
