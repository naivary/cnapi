package openapi

const _contentTypeJSON = "application/json"

type HeaderDataType int

const (
	STRING = iota + 1
	INT
)

type SecurityType int

const (
	APIKEY SecurityType = iota + 1
	HTTP
	MUTUALTLS
	OAUTH2
	OPENIDCONNECT
)

type Style int

const (
	MATRIX Style = iota + 1
	LABEL
	SIMPLE
	FORM
	SPACEDELIM
	PIPEDELIM
	DEEPOBJ
	COOKIEE
)

type In int

const (
	PATH In = iota + 1
	QUERY
	QUERYSTR
	HEADER
	COOKIE
)

type OpenAPI struct {
	Version           string `json:"openapi"`
	Self              string `json:"$self"`
	Info              *Info
	JSONSchemaDialect string
	Servers           []*Server
	Paths             map[string]*PathItem
	Webhooks          map[string]*PathItem
	Security          []*SecurityRequirement
	Tags              []*Tag
	ExternalDocs      *ExternalDoc
}

type Info struct {
	Version        string
	Title          string
	Summary        string
	Description    string
	TermsOfService string
	Contact        *Contact
	License        *License
}

type Contact struct {
	Name  string
	URL   string
	Email string
}

type License struct {
	Name       string
	Identifier string
	URL        string
}

type Server struct {
	URL         string
	Description string
	Name        string
	Variables   map[string]*ServerVariable
}

type ServerVariable struct {
	Default     string
	Enum        []string
	Description string
}

type PathItem struct {
	Ref                  string                `json:"$ref"`
	Summary              string                `json:"summary"`
	Description          string                `json:"description"`
	Get                  *Operation            `json:"get"`
	Put                  *Operation            `json:"put"`
	Post                 *Operation            `json:"post"`
	Delete               *Operation            `json:"delete"`
	Options              *Operation            `json:"options"`
	Head                 *Operation            `json:"head"`
	Patch                *Operation            `json:"patch"`
	Trace                *Operation            `json:"trace"`
	Query                *Operation            `json:"query"`
	AdditionalProperties map[string]*Operation `json:"additionalProperties"`
	Servers              []*Server             `json:"servers"`
	Parameters           []*Parameter          `json:"parameters"`
}

type Operation struct {
	Tags         []string
	Summary      string
	Description  string
	ExternalDocs *ExternalDoc
	OperationID  string
	Parameters   []*Parameter
	RequestBody  *RequestBody
	Responses    map[string]*Response
	Callbacks    map[string]*PathItem
	Deprecated   bool
	Security     []SecurityRequirement
	Servers      []*Server
}

type ExternalDoc struct {
	URL         string
	Description string
}

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

type Example struct {
	Ref             string `json:"$ref"`
	Summary         string
	Description     string
	DataValue       any
	SerializedValue string
	ExternalValue   string
	Value           any
}

type MediaType struct {
	Schema         *Schema
	ItemSchema     *Schema
	Example        any
	Examples       map[string]*Example
	Encoding       map[string]*Encoding
	PrefixEncoding []*Encoding
	ItemEncoding   *Encoding
}

type Schema struct {
	Ref string `json:"$ref"`
}

type Encoding struct {
	ContentType    string
	Headers        map[string]*Header
	Encoding       map[string]*Encoding
	PrefixEncoding []*Encoding
	ItemEncoding   *Encoding
}

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

func (h *Header) Deprecate() *Header {
	h.Deprecated = true
	return h
}

func (h *Header) Exploded() *Header {
	h.Explode = true
	return h
}

type RequestBody struct {
	Description string
	Required    bool
	Content     map[string]*MediaType
}

type Response struct {
	Ref         string `json:"$ref"`
	Summary     string
	Description string
	Headers     map[string]*Header
	Content     map[string]*MediaType
	Links       map[string]*Link
}

type Link struct {
	OperationRef string
	OperationID  string
	Parameters   map[string]any
	RequestBody  map[string]any
	Description  string
	Server       *Server
}

type SecurityRequirement map[string][]string

type Tag struct {
	Name         string
	Summary      string
	Description  string
	ExternalDocs *ExternalDoc
	Parent       string
	Kind         string
}

type Components struct {
	Schemas         map[string]*Schema
	Responses       map[string]*Response
	Parameters      map[string]*Parameter
	Examples        map[string]*Example
	RequestBodies   map[string]*RequestBody
	Headers         map[string]*Header
	SecuritySchemas map[string]*SecurityScheme
	Links           map[string]*Link
	Callbacks       map[string]*PathItem
	PathItems       map[string]*PathItem
	MediaTypes      map[string]*MediaType
}

type SecurityScheme struct {
	Type             SecurityType
	Description      string
	Name             string
	In               string
	Scheme           string
	BearerFormat     string
	Flows            *OAuthFlows
	OpenIDConnectURL string
	OAuth2MetdataURL string
	Deprecated       bool
}

type OAuthFlows struct {
	Implicit            *OAuthFlow
	Password            *OAuthFlow
	ClientCredentials   *OAuthFlow
	AuthorizationCode   *OAuthFlow
	DeviceAuthorization *OAuthFlow
}

type OAuthFlow struct {
	AuthorizationURL       string
	DeviceAuthorizationURL string
	TokenURL               string
	RefreshURL             string
	Scopes                 map[string]string
}

func Request(desc string, required bool, model any) *RequestBody {
	return &RequestBody{
		Description: desc,
		Required:    required,
		Content: map[string]*MediaType{
			_contentTypeJSON: {
				Schema: &Schema{Ref: "#/components/schemas/<model-name>"},
			},
		},
	}
}

func Res(desc, summary string, model any, headers ...*Header) *Response {
	res := &Response{
		Description: desc,
		Summary:     summary,
	}
	return res
}

func NewHeader(desc string, required bool, dataType HeaderDataType) *Header {
	h := &Header{
		Description: desc,
		Required:    required,
		Deprecated:  deprecated,
	}
	return h
}
