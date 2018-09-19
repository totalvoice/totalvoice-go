package api

import (
	"encoding/json"
	"net/http"
)

// Response - struct
type Response struct{}

// HandleResponse - realiza o tratamento do Response
func (r Response) HandleResponse(modelResponse interface{}, httpResponse *http.Response) interface{} {

	if httpResponse != nil {
		defer httpResponse.Body.Close()
		json.NewDecoder(httpResponse.Body).Decode(&modelResponse)
	}
	return modelResponse
}
