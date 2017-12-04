package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// ChamadaService service
type ChamadaService struct {
	client  HTTPClient
	handler Response

	Relatorio *ChamadaRelatorioService
}

// NewChamadaService - Serviço para o envio de Chamada
func NewChamadaService(httpClient HTTPClient, handler Response) *ChamadaService {

	service := &ChamadaService{
		client:  httpClient,
		handler: handler,
		Relatorio: &ChamadaRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Criar - Envia uma mensagem de Chamada
func (s ChamadaService) Criar(numeroOrigem string, numeroDestino string, opcoes map[string]interface{}) (*model.TotalVoiceResponse, error) {

	chamada := new(model.Chamada)
	chamada.NumeroOrigem = numeroOrigem
	chamada.NumeroDestino = numeroDestino
	//chamada.URLChamada = urlChamada
	//chamada.RespostaUsuario = respostaUsuario
	//chamada.Bina = bina

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(chamada, RotaChamada)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// BuscaChamada - Busca uma Chamada pelo seu ID
func (s ChamadaService) BuscaChamada(id int) (*model.ChamadaResponse, error) {

	sID := strconv.Itoa(id)
	chamada := new(model.Chamada)
	response := new(model.ChamadaResponse)

	http, err := s.client.GetResource(chamada, RotaChamada, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ChamadaResponse), err
}

// Excluir - Remove uma Conta
func (s ChamadaService) Excluir(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaChamada, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}
