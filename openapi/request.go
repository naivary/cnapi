package openapi

type RequestBody struct {
	Description string
	Required    bool
	Content     map[string]*MediaType
}

func NewReqBody(desc string, required bool, model any) *RequestBody {
	req := &RequestBody{
		Description: desc,
		Required:    required,
	}
	return req
}
