package openapi

type SecurityType int

const (
	API_KEY SecurityType = iota + 1
	HTTP
	MUTUAL_TLS
	OAUTH2
	OPENID_CONNECT
)

type Style int

const (
	MATRIX Style = iota + 1
	LABEL
	SIMPLE
	FORM
	SPACE_DELIM
	PIPE_DELIM
	DEEPOBJ
	S_COOKIE
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
