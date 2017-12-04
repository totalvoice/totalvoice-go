package api

import (
	"strconv"
	"time"

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
func (s ChamadaService) Criar(numeroOrigem string, numeroDestino string, opcoes map[string]string) (*model.TotalVoiceResponse, error) {

	chamada := new(model.Chamada)
	chamada.NumeroOrigem = numeroOrigem
	chamada.NumeroDestino = numeroDestino

	for index, value := range opcoes {
		switch index {
		case "data_criacao":
			t, _ := time.Parse(DateFormatUTC, value)
			chamada.DataCriacao = t
		case "gravar_audio":
			b, _ := strconv.ParseBool(value)
			chamada.GravarAudio = b
		case "bina_origem":
			chamada.BinaOrigem = value
		case "bina_destino":
			chamada.BinaDestino = value
		case "tags":
			chamada.Tags = value
		}

	}

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

// DownloadGravacao - Download do áudio de uma chamada gravada
func (s ChamadaService) DownloadGravacao(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	chamada := new(model.Chamada)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.GetResource(chamada, RotaChamada+"/"+sID+"/gravacao", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}
