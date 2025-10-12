package openapi

type Response struct {
	Ref         string `json:"$ref"`
	Summary     string
	Description string
	Headers     map[string]*Header
	Content     map[string]*MediaType
	Links       map[string]*Link
}

func NewResponse(desc string, model any) *Response {
	res := &Response{
		Headers: make(map[string]*Header),
		Links:   make(map[string]*Link),
	}
	return res
}

func (r *Response) AddHeader(name string, h *Header) *Response {
	r.Headers[name] = h
	return r
}

func (r *Response) RefTo(ref string) *Response {
	r.Ref = ref
	return r
}
