package totalvoice

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Client struct
type Client struct {
	accessToken string
	baseURI     string
	client      *http.Client
}

// CreateResource - HTTP POST
func (c *Client) CreateResource(values map[string]string, path string, v interface{}) error {
	return c.makeRequest("POST", path, values, v)
}

// UpdateResource - HTTP PUT
func (c *Client) UpdateResource(values map[string]string, path string, sid interface{}) error {
	url := c.url(sid, path)
	return c.makeRequest("PUT", url, values, nil)
}

// GetResource - HTTP GET
func (c *Client) GetResource(path string, sid interface{}, v interface{}) error {
	url := c.url(sid, path)
	return c.makeRequest("GET", url, nil, v)
	//if err := json.Unmarshal([]byte(r), &v); err != nil {
	//	return v, err
	//}
}

// ListResource - HTTP GET
func (c *Client) ListResource(path string, v interface{}, params map[string]string) error {
	return c.makeRequest("GET", path, params, v)
}

// DeleteResource - HTTP DELETE
func (c *Client) DeleteResource(path string, sid interface{}) error {
	url := c.url(sid, path)
	return c.makeRequest("DELETE", url, nil, nil)
}

// Make a request to the Twilio API.
func (c *Client) makeRequest(method string, path string, values map[string]string, v interface{}) error {

	uri := c.buildURI(path)
	client := c.client
	if client == nil {
		client = http.DefaultClient
	}

	rb := new(strings.Reader)
	if len(values) > 0 && (method == "POST" || method == "PUT") {
		r, _ := json.Marshal(values)
		rb = strings.NewReader(string(r))
	}

	if method == "GET" && len(values) > 0 {
		if len(values) > 0 {
			query := c.buildQueryString(values)
			uri = uri + "?" + query
		}
	}
	req, err := http.NewRequest(method, uri, rb)
	if err != nil {
		return err
	}
	//req = withContext(req, ctx)
	req.Header.Add("Access-Token", c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&v)

	return nil
}

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

func (c *Client) buildQueryString(values map[string]string) string {
	params := url.Values{}
	for i, v := range values {
		params.Add(i, v)
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
