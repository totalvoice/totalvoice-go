package api

import (
	"strconv"
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ChamadaService service
type ChamadaService struct {
	client   HTTPClient
	response Response

	Relatorio *ChamadaRelatorioService
}

// NewChamadaService - Serviço para o envio de Chamada
func NewChamadaService(httpClient HTTPClient, response Response) *ChamadaService {

	service := &ChamadaService{
		client:   httpClient,
		response: response,
		Relatorio: &ChamadaRelatorioService{
			client:   httpClient,
			response: response,
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

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(chamada, RotaChamada)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma Chamada pelo seu ID
func (s ChamadaService) Buscar(id int) (*model.ChamadaResponse, error) {

	sID := strconv.Itoa(id)
	chamada := new(model.Chamada)
	resp := new(model.ChamadaResponse)

	http, err := s.client.GetResource(chamada, RotaChamada, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ChamadaResponse), err
}

// Encerrar - Encerra uma chamada ativa
func (s ChamadaService) Encerrar(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaChamada, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// DownloadGravacao - Download do áudio de uma chamada gravada
func (s ChamadaService) DownloadGravacao(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	chamada := new(model.Chamada)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.GetResource(chamada, RotaChamada+"/"+sID+"/gravacao", nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Escutar - Escuta uma chamada ativa
func (s ChamadaService) Escutar(id int, numero string, modo int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	escuta := new(model.Escuta)
	escuta.Numero = numero
	escuta.Modo = modo

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(escuta, RotaChamada+"/"+sID+"/escuta")
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Transferir - Faz uma transferência da chamada atual
func (s ChamadaService) Transferir(id int, numero string, perna string) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	transfer := new(model.Transferencia)
	transfer.Numero = numero
	transfer.Perna = perna

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(transfer, RotaChamada+"/"+sID+"/transfer")
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Avaliar - Faz uma transferência da chamada atual
func (s ChamadaService) Avaliar(id int, nota string, comentario string) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	avaliacao := new(model.Avaliacao)
	avaliacao.Nota = nota
	avaliacao.Comentario = comentario

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(avaliacao, RotaChamada+"/"+sID+"/avaliar")
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}
