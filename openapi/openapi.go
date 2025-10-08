package openapi

type OpenAPISecurityType int

const (
	APIKEY OpenAPISecurityType = iota + 1
	HTTP
	MUTUALTLS
	OAUTH2
	OPENIDCONNECT
)

type OpenAPIStyle int

const (
	MATRIX OpenAPIStyle = iota + 1
	LABEL
	SIMPLE
	FORM
	SPACEDELIM
	PIPEDELIM
	DEEPOBJ
	COOKIEE
)

type OpenAPIIn int

const (
	PATH OpenAPIIn = iota + 1
	QUERY
	QUERYSTR
	HEADER
	COOKIE
)

type OpenAPI struct {
	Version           string `json:"openapi"`
	Self              string `json:"$self"`
	Info              *OpenAPIInfo
	JSONSchemaDialect string
	Servers           []*OpenAPIServer
	Paths             map[string]*OpenAPIPathItem
	Webhooks          map[string]*OpenAPIPathItem
	Security          []*OpenAPISecurityRequirement
	Tags              []*OpenAPITag
	ExternalDocs      *OpenAPIExternalDoc
}

type OpenAPIInfo struct {
	Version        string
	Title          string
	Summary        string
	Description    string
	TermsOfService string
	Contact        OpenAPIContact
	License        OpenAPILicense
}

type OpenAPIContact struct {
	Name  string
	URL   string
	Email string
}

type OpenAPILicense struct {
	Name       string
	Identifier string
	URL        string
}

type OpenAPIServer struct {
	URL         string
	Description string
	Name        string
	Variables   map[string]OpenAPIServerVariable
}

type OpenAPIServerVariable struct {
	Default     string
	Enum        []string
	Description string
}

type OpenAPIPathItem struct {
	Ref                  string `json:"$ref"`
	Summary              string
	Description          string
	Get                  *OpenAPIOperation
	Put                  *OpenAPIOperation
	Post                 *OpenAPIOperation
	Delete               *OpenAPIOperation
	Options              *OpenAPIOperation
	Head                 *OpenAPIOperation
	Patch                *OpenAPIOperation
	Trace                *OpenAPIOperation
	Query                *OpenAPIOperation
	AdditionalProperties map[string]*OpenAPIOperation
	Servers              []*OpenAPIServer
	Parameters           []*OpenAPIParameter
}

type OpenAPIOperation struct {
	Tags         []string
	Summary      string
	Description  string
	ExternalDocs *OpenAPIExternalDoc
	OperationID  string
	Parameters   *OpenAPIParameter
	RequestBody  *OpenAPIRequestBody
	Responses    map[string]*OpenAPIResponse
	Callbacks    map[string]*OpenAPIPathItem
	Deprecated   bool
	Security     OpenAPISecurityRequirement
	Server       *OpenAPIServer
}

type OpenAPIExternalDoc struct {
	URL         string
	Description string
}

type OpenAPIParameter struct {
	Ref             string `json:"$ref"`
	Name            string
	In              OpenAPIIn
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool
	Example         any
	Examples        map[string]*OpenAPIExample

	// Used with Schema
	Style         OpenAPIStyle
	Explode       bool
	AllowReserved bool

	// Used with content
	Content map[string]*OpenAPIMediaType
}

type OpenAPIExample struct {
	Ref             string `json:"$ref"`
	Summary         string
	Description     string
	DataValue       any
	SerializedValue string
	ExternalValue   string
	Value           any
}

type OpenAPIMediaType struct {
	Schema         OpenAPISchema
	ItemSchema     OpenAPISchema
	Example        any
	Examples       map[string]*OpenAPIExample
	Encoding       map[string]*OpenAPIEncoding
	PrefixEncoding []*OpenAPIEncoding
	ItemEncoding   *OpenAPIEncoding
}

type OpenAPISchema struct{}

type OpenAPIEncoding struct {
	ContentType    string
	Headers        map[string]*OpenAPIHeader
	Encoding       map[string]*OpenAPIEncoding
	PrefixEncoding []*OpenAPIEncoding
	ItemEncoding   *OpenAPIEncoding
}

type OpenAPIHeader struct {
	Description string
	Required    bool
	Deprecated  bool
	Example     any
	Examples    map[string]*OpenAPIExample

	// Fixed Fields of schema
	Schema  *OpenAPISchema
	Style   OpenAPIStyle
	Explode bool

	Content map[string]*OpenAPIMediaType
}

type OpenAPIRequestBody struct {
	Description string
	Required    bool
	Content     map[string]*OpenAPIMediaType
}

type OpenAPIResponse struct {
	Ref         string `json:"$ref"`
	Summary     string
	Description string
	Headers     map[string]*OpenAPIHeader
	Content     map[string]*OpenAPIMediaType
	Links       map[string]*OpenAPILink
}

type OpenAPILink struct {
	OperationRef string
	OperationID  string
	Parameters   map[string]any
	RequestBody  map[string]any
	Description  string
	Server       *OpenAPIServer
}

type OpenAPISecurityRequirement map[string][]string

type OpenAPITag struct {
	Name         string
	Summary      string
	Description  string
	ExternalDocs *OpenAPIExternalDoc
	Parent       string
	Kind         string
}

type OpenAPIComponents struct {
	Schemas         map[string]*OpenAPISchema
	Responses       map[string]*OpenAPIResponse
	Parameters      map[string]*OpenAPIParameter
	Examples        map[string]*OpenAPIExample
	RequestBodies   map[string]*OpenAPIRequestBody
	Headers         map[string]*OpenAPIHeader
	SecuritySchemas map[string]*OpenAPISecurityScheme
	Links           map[string]*OpenAPILink
	Callbacks       map[string]*OpenAPIPathItem
	PathItems       map[string]*OpenAPIPathItem
	MediaTypes      map[string]*OpenAPIMediaType
}

type OpenAPISecurityScheme struct {
	Type             OpenAPISecurityType
	Description      string
	Name             string
	In               string
	Scheme           string
	BearerFormat     string
	Flows            *OpenAPIOAuthFlows
	OpenIDConnectURL string
	OAuth2MetdataURL string
	Deprecated       bool
}

type OpenAPIOAuthFlows struct {
	Implicit            *OpenAPIOAuthFlow
	Password            *OpenAPIOAuthFlow
	ClientCredentials   *OpenAPIOAuthFlow
	AuthorizationCode   *OpenAPIOAuthFlow
	DeviceAuthorization *OpenAPIOAuthFlow
}

type OpenAPIOAuthFlow struct {
	AuthorizationURL       string
	DeviceAuthorizationURL string
	TokenURL               string
	RefreshURL             string
	Scopes                 map[string]string
}
