package main

import (
	"net/http"
	"net/url"
	"strings"
)

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// TotalVoice struct
type TotalVoice struct {
	AccessToken string
	baseURI     interface{}
	client      *http.Client
}

// NewTotalVoiceClient - Cria TotalVoice struct.
//func NewTotalVoiceClient(accessToken string, baseURI interface{}) *TotalVoice {
//	return NewClient{accessToken, baseURI, http.DefaultClient}
//}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI interface{}, client *http.Client) *TotalVoice {

	if client == nil {
		client = http.DefaultClient
	}

	if baseURI == nil {
		baseURI = BaseURI
	}

	return &TotalVoice{accessToken, baseURI, client}
}

// post -
func (tvce *TotalVoice) post(formValues url.Values, path string) (*http.Response, error) {

	req, err := http.NewRequest("POST", path, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Access-Token", tvce.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	client := tvce.client
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req)
}

// get -
func (tvce *TotalVoice) get(path string) (*http.Response, error) {

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	client := tvce.client
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req)
}
