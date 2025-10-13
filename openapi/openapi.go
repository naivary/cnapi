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
	Version           string                 `json:"openapi"`
	Self              string                 `json:"$self,omitempty"`
	Info              *Info                  `json:"info"`
	JSONSchemaDialect string                 `json:"jsonSchemaDialect,omitempty"`
	Servers           []*Server              `json:"servers,omitempty"`
	Paths             map[string]*PathItem   `json:"paths"`
	Webhooks          map[string]*PathItem   `json:"webhooks,omitempty"`
	Components        *Components            `json:"components,omitempty"`
	Security          []*SecurityRequirement `json:"security,omitempty"`
	Tags              []*Tag                 `json:"tags,omitempty"`
	ExternalDocs      *ExternalDoc           `json:"externalDocs,omitempty"`
}

func New(info *Info) *OpenAPI {
	return &OpenAPI{
		Version: "3.2.0",
		Info:    info,
		Paths:   make(map[string]*PathItem),
		Components: &Components{
			Schemas: make(map[string]*Schema),
		},
	}
}

type Info struct {
	Version        string   `json:"version"`
	Title          string   `json:"title"`
	Summary        string   `json:"summary,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}

type Server struct {
	URL         string                     `json:"url"`
	Description string                     `json:"description,omitempty"`
	Name        string                     `json:"name,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}

type ServerVariable struct {
	Default     string   `json:"default"`
	Enum        []string `json:"enum,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs *ExternalDoc          `json:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   []*Parameter          `json:"parameters,omitempty"`
	RequestBody  *RequestBody          `json:"requestBody,omitempty"`
	Responses    map[string]*Response  `json:"responses"`
	Callbacks    map[string]*PathItem  `json:"callbacks,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty"`
	Servers      []*Server             `json:"servers,omitempty"`
}

type ExternalDoc struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type Example struct {
	Ref             string `json:"$ref,omitempty"`
	Summary         string `json:"summary,omitempty"`
	Description     string `json:"description,omitempty"`
	DataValue       any    `json:"dataValue,omitempty"`
	SerializedValue string `json:"serializedValue,omitempty"`
	ExternalValue   string `json:"externalValue,omitempty"`
	Value           any    `json:"value,omitempty"`
}

type MediaType struct {
	Schema         *Schema              `json:"schema,omitempty"`
	ItemSchema     *Schema              `json:"itemSchema,omitempty"`
	Example        any                  `json:"example,omitempty"`
	Examples       map[string]*Example  `json:"examples,omitempty"`
	Encoding       map[string]*Encoding `json:"encoding,omitempty"`
	PrefixEncoding []*Encoding          `json:"prefixEncoding,omitempty"`
	ItemEncoding   *Encoding            `json:"itemEncoding,omitempty"`
}

type JSONType string

const (
	InvalidType JSONType = "invalid"
	NullType    JSONType = "null"
	BooleanType JSONType = "boolean"
	NumberType  JSONType = "number"
	IntegerType JSONType = "integer"
	StringType  JSONType = "string"
	ArrayType   JSONType = "array"
	ObjectType  JSONType = "object"
)

type Schema struct {
	// metadata
	ID    string   `json:"$id,omitzero"`
	Draft string   `json:"$schema,omitzero"`
	Ref   string   `json:"$ref,omitzero"`
	Type  JSONType `json:"type,omitzero"`

	OneOf []*Schema `json:"oneOf,omitzero"`
	AnyOf []*Schema `json:"anyOf,omitzero"`
	Not   *Schema   `json:"not,omitzero"`

	// agnostic
	Enum []any `json:"enum,omitzero"`

	// annotations
	Title      string `json:"title,omitzero"`
	Desc       string `json:"description,omitzero"`
	Examples   []any  `json:"examples,omitzero"`
	Deprecated bool   `json:"deprecated,omitzero"`
	WriteOnly  bool   `json:"writeOnly,omitzero"`
	ReadOnly   bool   `json:"readOnly,omitzero"`
	Default    string `json:"default,omitzero"`

	// array
	MaxItems    int64   `json:"maxItems,omitzero"`
	MinItems    int64   `json:"minItems,omitzero"`
	UniqueItems bool    `json:"uniqueItems,omitzero"`
	Items       *Schema `json:"items,omitzero"`

	// object
	Properties           map[string]*Schema  `json:"properties,omitzero"`
	Required             []string            `json:"required,omitzero"`
	AdditionalProperties *Schema             `json:"additionalProperties,omitzero"`
	PatternProperties    map[string]*Schema  `json:"patternProperties,omitzero,omitempty"`
	DependentRequired    map[string][]string `json:"dependentRequired,omitzero,omitempty"`

	// string
	MinLength        int64  `json:"minLength,omitzero"`
	MaxLength        int64  `json:"maxLength,omitzero"`
	Pattern          string `json:"pattern,omitzero"`
	ContentEncoding  string `json:"contentEnconding,omitzero"`
	ContentMediaType string `json:"contentMediaType,omitzero"`
	Format           string `json:"format,omitzero"`

	// number
	Maximum          int64 `json:"maximum,omitzero"`
	Minimum          int64 `json:"minimum,omitzero"`
	ExclusiveMaximum int64 `json:"exclusiveMaximum,omitzero"`
	ExclusiveMinimum int64 `json:"exclusiveMinimum,omitzero"`
	MultipleOf       int64 `json:"multipleOf,omitzero"`
}

type Encoding struct {
	ContentType    string               `json:"contentType,omitempty"`
	Headers        map[string]*Header   `json:"headers,omitempty"`
	Encoding       map[string]*Encoding `json:"encoding,omitempty"`
	PrefixEncoding []*Encoding          `json:"prefixEncoding,omitempty"`
	ItemEncoding   *Encoding            `json:"itemEncoding,omitempty"`
}

type Link struct {
	OperationRef string         `json:"operationRef,omitempty"`
	OperationID  string         `json:"operationId,omitempty"`
	Parameters   map[string]any `json:"parameters,omitempty"`
	RequestBody  map[string]any `json:"requestBody,omitempty"`
	Description  string         `json:"description,omitempty"`
	Server       *Server        `json:"server,omitempty"`
}

type SecurityRequirement map[string][]string

type Tag struct {
	Name         string       `json:"name"`
	Summary      string       `json:"summary,omitempty"`
	Description  string       `json:"description,omitempty"`
	ExternalDocs *ExternalDoc `json:"externalDocs,omitempty"`
	Parent       string       `json:"parent,omitempty"`
	Kind         string       `json:"kind,omitempty"`
}

type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty"`
	Examples        map[string]*Example        `json:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]*Header         `json:"headers,omitempty"`
	SecuritySchemas map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
	Links           map[string]*Link           `json:"links,omitempty"`
	Callbacks       map[string]*PathItem       `json:"callbacks,omitempty"`
	PathItems       map[string]*PathItem       `json:"pathItems,omitempty"`
	MediaTypes      map[string]*MediaType      `json:"mediaTypes,omitempty"`
}

type SecurityScheme struct {
	Type             SecurityType `json:"type"`
	Description      string       `json:"description,omitempty"`
	Name             string       `json:"name,omitempty"`
	In               string       `json:"in,omitempty"`
	Scheme           string       `json:"scheme,omitempty"`
	BearerFormat     string       `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlows  `json:"flows,omitempty"`
	OpenIDConnectURL string       `json:"openIdConnectUrl,omitempty"`
	OAuth2MetdataURL string       `json:"oauth2MetadataUrl,omitempty"`
	Deprecated       bool         `json:"deprecated,omitempty"`
}

type OAuthFlows struct {
	Implicit            *OAuthFlow `json:"implicit,omitempty"`
	Password            *OAuthFlow `json:"password,omitempty"`
	ClientCredentials   *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode   *OAuthFlow `json:"authorizationCode,omitempty"`
	DeviceAuthorization *OAuthFlow `json:"deviceAuthorization,omitempty"`
}

type OAuthFlow struct {
	AuthorizationURL       string            `json:"authorizationUrl,omitempty"`
	DeviceAuthorizationURL string            `json:"deviceAuthorizationUrl,omitempty"`
	TokenURL               string            `json:"tokenUrl,omitempty"`
	RefreshURL             string            `json:"refreshUrl,omitempty"`
	Scopes                 map[string]string `json:"scopes"`
}
