package totalvoice

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	uri := c.buildURI(path)
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	return c.generateRequest(req)
}

// UpdateResource - HTTP PUT
func (c *Client) UpdateResource(values map[string]string, path string) (string, error) {

	uri := c.buildURI(path)
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	return c.generateRequest(req)
}

// GetResource - HTTP GET
func (c *Client) GetResource(path string) (string, error) {

	uri := c.buildURI(path)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	return c.generateRequest(req)
}

// DeleteResource - HTTP DELETE
func (c *Client) DeleteResource(path string) (string, error) {

	uri := c.buildURI(path)
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return "", err
	}

	return c.generateRequest(req)
}

// buildURI - Monta URI de acordo com o path
func (c *Client) buildURI(path string) string {
	base := c.GetBaseURI()
	s := []string{base, path}

	return strings.Join(s, "")
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
