package api

// HTTPClient Interface for Clients
type HTTPClient interface {
	GetResource(path string, sid interface{}, v interface{}) error
	ListResource(path string, v interface{}, params map[string]string) error
	CreateResource(values map[string]string, path string, v interface{}) error
	UpdateResource(values map[string]string, path string, sid interface{}) error
	DeleteResource(path string, sid interface{}) error
	SetBaseURI(value string)
	GetBaseURI() string
}
