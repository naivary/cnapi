package openapi

type Header struct {
	Description string
	Required    bool
	Deprecated  bool
	Example     any
	Examples    map[string]*Example

	// Fixed Fields of schema
	Schema  *Schema
	Style   Style
	Explode bool

	Content map[string]*MediaType
}

func NewHeader(desc string, required bool) *Header {
	h := &Header{
		Description: desc,
		Required:    required,
	}
	return h
}

func (h *Header) Deprecate() *Header {
	h.Deprecated = true
	return h
}
