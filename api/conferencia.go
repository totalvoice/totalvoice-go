package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ConferenciaService service
type ConferenciaService struct {
	client   HTTPClient
	response Response

	Relatorio *AudioRelatorioService
}

// NewConferenciaService - Serviço para realizar conferências
func NewConferenciaService(httpClient HTTPClient, response Response) *ConferenciaService {

	service := &ConferenciaService{
		client:   httpClient,
		response: response,
	}

	return service
}

// Criar - Cria uma conferência e retorna o ID
func (s ConferenciaService) Criar() (*model.TotalVoiceResponse, error) {

	conf := new(model.Conferencia)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(conf, RotaConferencia)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma chamada pelo seu ID
func (s ConferenciaService) Buscar(id int) (*model.ChamadaResponse, error) {

	sID := strconv.Itoa(id)
	conf := new(model.Conferencia)
	resp := new(model.ChamadaResponse) // retorna os dados de uma chamada

	http, err := s.client.GetResource(conf, RotaConferencia, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ChamadaResponse), err
}

// AddNumeroConferencia - Busca uma chamada pelo seu ID
func (s ConferenciaService) AddNumeroConferencia(id int, numero string, bina string, gravarAudio bool) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	conf := new(model.Conferencia)
	conf.Numero = numero
	conf.Bina = bina
	conf.GravarAudio = gravarAudio

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(conf, RotaConferencia+"/"+sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}
