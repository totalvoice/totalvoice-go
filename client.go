package totalvoice

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
func (c *Client) CreateResource(values map[string]string, path string) (string, error) {
	return c.makeRequest("POST", path, values)
}

// UpdateResource - HTTP PUT
func (c *Client) UpdateResource(values map[string]string, path string) (string, error) {
	return c.makeRequest("PUT", path, values)
}

// GetResource - HTTP GET
func (c *Client) GetResource(path string, params map[string]string) (string, error) {
	return c.makeRequest("GET", path, params)
}

// DeleteResource - HTTP DELETE
func (c *Client) DeleteResource(path string) (string, error) {
	return c.makeRequest("DELETE", path, nil)
}

// buildURI - Monta URI de acordo com o path
func (c *Client) buildURI(path string) string {
	base := c.GetBaseURI()
	s := []string{base, path}

	return strings.Join(s, "")
}

// MakeRequest -
func (c *Client) makeRequest(method string, path string, values map[string]string) (string, error) {

	uri := c.buildURI(path)

	if len(values) > 0 && (method == "POST" || method == "PUT") {
		rb, _ := json.Marshal(values)
		req, err := http.NewRequest(method, uri, bytes.NewBuffer(rb))
		if err != nil {
			return "", err
		}
		return c.generateRequest(req)
	}

	if method == "GET" {

		if len(values) > 0 {
			query := c.buildQueryString(values)
			uri = uri + "?" + query
		}
		req, err := http.NewRequest(method, uri, nil)
		if err != nil {
			return "", err
		}
		return c.generateRequest(req)
	}

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return "", err
	}

	return c.generateRequest(req)
}

func (c *Client) buildQueryString(values map[string]string) string {
	params := url.Values{}
	for i, v := range values {
		params.Add(i, v)
	}
	return params.Encode()
}

// generateRequest - Trata o response e retorna como string
func (c *Client) generateRequest(req *http.Request) (string, error) {

	client := c.client
	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Add("Access-Token", c.accessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SetBaseURI - seta a baseURI
func (c *Client) SetBaseURI(value string) {
	c.baseURI = value
}

// GetBaseURI - Get a base URL
func (c *Client) GetBaseURI() string {
	return c.baseURI
}
