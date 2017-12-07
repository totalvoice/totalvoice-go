package api

import "net/http"

// HTTPClient Interface for Clients
type HTTPClient interface {
	GetResource(model Model, path string, sid interface{}) (*http.Response, error)
	ListResource(model Model, path string, params map[string]interface{}) (*http.Response, error)
	CreateResource(model Model, path string) (*http.Response, error)
	UpdateResource(model Model, path string, sid interface{}) (*http.Response, error)
	DeleteResource(path string, sid interface{}) (*http.Response, error)
	SetBaseURI(value string)
	GetBaseURI() string
}
