package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// ConferenciaService service
type ConferenciaService struct {
	client  HTTPClient
	handler Response

	Relatorio *AudioRelatorioService
}

// NewConferenciaService - Serviço para realizar conferências
func NewConferenciaService(httpClient HTTPClient, handler Response) *ConferenciaService {

	service := &ConferenciaService{
		client:  httpClient,
		handler: handler,
	}

	return service
}

// Criar - Cria uma conferência e retorna o ID
func (s ConferenciaService) Criar() (*model.TotalVoiceResponse, error) {

	conf := new(model.Conferencia)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(conf, RotaConferencia)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// BuscaChamada - Busca uma chamada pelo seu ID
func (s ConferenciaService) BuscaChamada(id int) (*model.ChamadaResponse, error) {

	sID := strconv.Itoa(id)
	conf := new(model.Conferencia)
	response := new(model.ChamadaResponse) // retorna os dados de uma chamada

	http, err := s.client.GetResource(conf, RotaConferencia, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ChamadaResponse), err
}

// AddNumeroConferencia - Busca uma chamada pelo seu ID
func (s ConferenciaService) AddNumeroConferencia(id int, numero string, bina string, gravarAudio bool) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	conf := new(model.Conferencia)
	conf.Numero = numero
	conf.Bina = bina
	conf.GravarAudio = gravarAudio

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(conf, RotaConferencia+"/"+sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}
