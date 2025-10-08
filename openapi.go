package main

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
	Info              OpenAPIInfo
	JSONSchemaDialect string
	Servers           []OpenAPIServer
	Paths             map[string]OpenAPIPathItem
	Webhooks          map[string]OpenAPIPathItem
	Components        OpenAPIComponents
	Security          []OpenAPISecurityRequirement
	Tags              []OpenAPITag
	ExternalDocs      OpenAPIExternalDoc
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

type OpenAPIComponents struct {
	Schemas   map[string]OpenAPISchema
	Responses map[string]OpenAPIResponse
	Parameters
	Examples
	RequestBodies
	Headers
	SecuritySchemes
	Links
	Callbacks
	PathItems
	MediaTypes
}

type OpenAPISchema struct{}

type OpenAPIResponse struct {
	Summary     string
	Description string
	// TODO: HJeaders is not this type
	Headers map[string]string
}

type OpenAPISecurityRequirement map[string][]string

type OpenAPIPathItem struct {
	Ref         string `json:"$ref"`
	Summary     string
	Description string
	Get
}

type OpenAPIOperation struct {
	Tags         []string
	Summary      string
	Description  string
	ExternalDocs OpenAPIExternalDoc
	OperationID  string
	Parameters
	RequestBody
	Responses
	Callbacks
	Deprecated bool
	Security   OpenAPISecurityRequirement
	Servers    []OpenAPIServer
}

type OpenAPITag struct {
	Name         string
	Summary      string
	Description  string
	ExternalDocs OpenAPIExternalDoc
	Parent       string
	Kind         string
}

type OpenAPIExternalDoc struct {
	URL         string
	Description string
}
