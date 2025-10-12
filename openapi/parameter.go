package openapi

type In int

const (
	PATH In = iota + 1
	QUERY
	// TODO: Check when Query String is needed
	// QUERYSTR
	HEADER
	COOKIE
)

type Parameter struct {
	Ref             string `json:"$ref"`
	Name            string
	In              In
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool
	Example         any
	Examples        map[string]*Example

	// Used with Schema
	Schema        *Schema
	Style         Style
	Explode       bool
	AllowReserved bool

	// Used with content
	Content map[string]*MediaType
}

func NewQueryParam(name, desc string, required bool) *Parameter {
	param := &Parameter{
		Name:        name,
		Description: desc,
		Required:    required,
		In:          QUERY,
	}
	return param
}

func NewCookieParam(name, desc string, required bool) *Parameter {
	param := &Parameter{
		Name:        name,
		Description: desc,
		Required:    required,
		In:          COOKIE,
	}
	return param
}

func NewHeaderParam(name, desc string, required bool) *Parameter {
	param := &Parameter{
		Name:        name,
		Description: desc,
		Required:    required,
		In:          HEADER,
	}
	return param
}

func newPathParam(name, desc string, required bool) *Parameter {
	param := &Parameter{
		Name:        name,
		Description: desc,
		Required:    required,
		In:          PATH,
	}
	return param
}

func (p *Parameter) Deprecate() *Parameter {
	p.Deprecated = true
	return p
}
