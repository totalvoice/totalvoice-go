package totalvoice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/totalvoice/totalvoice-go/api"
)

// Client struct
type Client struct {
	accessToken string
	baseURI     string
	client      *http.Client
}

// GetResource - HTTP GET
func (c *Client) GetResource(model api.Model, path string, sid interface{}) (*http.Response, error) {
	url := c.url(sid, path)
	return c.makeRequest("GET", model, url, nil)
}

// ListResource - HTTP GET
func (c *Client) ListResource(model api.Model, path string, params map[string]interface{}) (*http.Response, error) {
	return c.makeRequest("GET", model, path, params)
}

// CreateResource - HTTP POST
func (c *Client) CreateResource(model api.Model, path string) (*http.Response, error) {
	return c.makeRequest("POST", model, path, nil)
}

// UpdateResource - HTTP PUT
func (c *Client) UpdateResource(model api.Model, path string, sid interface{}) (*http.Response, error) {
	url := c.url(sid, path)
	return c.makeRequest("PUT", model, url, nil)
}

// DeleteResource - HTTP DELETE
func (c *Client) DeleteResource(path string, sid interface{}) (*http.Response, error) {
	url := c.url(sid, path)
	return c.makeRequest("DELETE", nil, url, nil)
}

// makeRequest Make a request to the TotalVoice API.
func (c *Client) makeRequest(method string, model api.Model, path string, params map[string]interface{}) (*http.Response, error) {

	uri := c.buildURI(path)
	client := c.client
	if client == nil {
		client = http.DefaultClient
	}

	b := new(bytes.Buffer)
	if method == "POST" || method == "PUT" {
		json.NewEncoder(b).Encode(model)
	}

	if method == "GET" && len(params) > 0 {
		if len(params) > 0 {
			query := c.buildQueryString(params)
			uri = uri + "?" + query
		}
	}
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return nil, err
	}

	// Headers
	req.Header.Add("Access-Token", c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

// url - monta a URL com o parametro ID
func (c *Client) url(sid interface{}, path string) string {
	if sid != nil {
		path = strings.Join([]string{path, sid.(string)}, "/")
	}
	return path
}

// buildURI - Monta URI de acordo com o path
func (c *Client) buildURI(path string) string {
	base := c.GetBaseURI()
	s := []string{base, path}

	return strings.Join(s, "")
}

// buildQueryString
func (c *Client) buildQueryString(values map[string]interface{}) string {
	params := url.Values{}
	for i, v := range values {
		switch v.(type) {
		case int:
			si := strconv.Itoa(v.(int))
			params.Add(i, si)
		case bool:
			sb := strconv.FormatBool(v.(bool))
			params.Add(i, sb)
		default:
			params.Add(i, v.(string))
		}
	}
	return params.Encode()
}

// SetBaseURI - seta a baseURI
func (c *Client) SetBaseURI(value string) {
	c.baseURI = value
}

// GetBaseURI - Get a base URL
func (c *Client) GetBaseURI() string {
	return c.baseURI
}
