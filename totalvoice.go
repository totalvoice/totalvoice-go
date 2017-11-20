package totalvoice

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// TotalVoice struct
type TotalVoice struct {
	accessToken string
	baseURI     string
	client      *http.Client

	Perfil *Perfil
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	c := &TotalVoice{accessToken: accessToken, baseURI: BaseURI}
	c.Perfil = &Perfil{client: c}

	return c
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	c := NewTotalVoiceClient(accessToken)
	c.SetBaseURI(baseURI)

	return c
}

// Post - HTTP POST
func (tvce *TotalVoice) Post(values map[string]string, path string) (string, error) {

	uri := tvce.buildURI(path)
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	return tvce.generateRequest(req)
}

// Put - HTTP PUT
func (tvce *TotalVoice) Put(values map[string]string, path string) (string, error) {

	uri := tvce.buildURI(path)
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	return tvce.generateRequest(req)
}

// Get - HTTP GET
func (tvce *TotalVoice) Get(path string) (string, error) {

	uri := tvce.buildURI(path)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	return tvce.generateRequest(req)
}

// Delete - HTTP DELETE
func (tvce *TotalVoice) Delete(path string) (string, error) {

	uri := tvce.buildURI(path)
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return "", err
	}

	return tvce.generateRequest(req)
}

// buildURI - Monta URI de acordo com o path
func (tvce *TotalVoice) buildURI(path string) string {
	base := tvce.GetBaseURI()
	s := []string{base, path}

	return strings.Join(s, "")
}

// generateRequest - Trata o response e retorna como string
func (tvce *TotalVoice) generateRequest(req *http.Request) (string, error) {

	client := tvce.client
	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Add("Access-Token", tvce.accessToken)
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
func (tvce *TotalVoice) SetBaseURI(value string) {
	tvce.baseURI = value
}

// GetBaseURI - Get a base URL
func (tvce *TotalVoice) GetBaseURI() string {
	if tvce.baseURI == "" {
		tvce.baseURI = BaseURI
	}
	return tvce.baseURI
}
