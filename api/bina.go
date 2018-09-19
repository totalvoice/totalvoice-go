package api

import (
	"github.com/totalvoice/totalvoice-go/api/model"
)

// BinaService service
type BinaService struct {
	client   HTTPClient
	response Response

	Relatorio *BinaRelatorioService
}

// NewBinaService - Serviço para validação da Bina
func NewBinaService(httpClient HTTPClient, response Response) *BinaService {

	service := &BinaService{
		client:   httpClient,
		response: response,
		Relatorio: &BinaRelatorioService{
			client:   httpClient,
			response: response,
		},
	}

	return service
}

// Enviar - Envia um número pra receber um código de validação
func (s BinaService) Enviar(telefone string) (*model.TotalVoiceResponse, error) {

	bina := new(model.Bina)
	bina.Telefone = telefone

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(bina, RotaBina)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Validar - Verifica se o código é válido para o telefone
func (s BinaService) Validar(codigo string, telefone string) (*model.ValidoResponse, error) {

	params := map[string]interface{}{
		"codigo":   codigo,
		"telefone": telefone,
	}

	bina := new(model.Bina)
	response := new(model.ValidoResponse)

	http, err := s.client.ListResource(bina, RotaBina, params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.ValidoResponse), err
}

// Excluir - Remove o telefone cadastrado na sua Conta
func (s BinaService) Excluir(telefone string) (*model.TotalVoiceResponse, error) {

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaBina, telefone)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}
