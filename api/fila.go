package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// FilaService service
type FilaService struct {
	client   HTTPClient
	response Response
}

// NewFilaService - Serviço para o gerenciamento da fila
func NewFilaService(httpClient HTTPClient, response Response) *FilaService {

	service := &FilaService{
		client:   httpClient,
		response: response,
	}

	return service
}

// Buscar fila
func (s FilaService) Buscar(id int) (*model.FilaResponse, error) {

	sID := strconv.Itoa(id)
	fila := new(model.Fila)
	response := new(model.FilaResponse)

	http, err := s.client.GetResource(fila, RotaFila, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.FilaResponse), err
}

// BuscarRamalFila - Busca ramal na fila
func (s FilaService) BuscarRamalFila(id int, ramalID int) (*model.FilaResponse, error) {

	sID := strconv.Itoa(id)
	sRamalID := strconv.Itoa(ramalID)

	fila := new(model.Fila)
	response := new(model.FilaResponse)

	http, err := s.client.GetResource(fila, RotaFila+"/"+sID+"/"+sRamalID, nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.FilaResponse), err
}

// Criar - Criar uma nova fila
func (s FilaService) Criar(fila model.Fila) (*model.FilaResponse, error) {

	resp := new(model.FilaResponse)
	http, err := s.client.CreateResource(fila, RotaFila)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.FilaResponse), err
}

// AdicionarRamalFila - Adiciona ramal na fila
func (s FilaService) AdicionarRamalFila(id int, membro model.FilaMembro) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	resp := new(model.TotalVoiceResponse)
	http, err := s.client.CreateResource(membro, RotaFila+"/"+sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Atualizar - Atualiza uma fila
func (s FilaService) Atualizar(fila model.Fila) (*model.FilaResponse, error) {

	sID := strconv.Itoa(fila.ID)
	resp := new(model.FilaResponse)
	http, err := s.client.UpdateResource(fila, RotaFila, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.FilaResponse), err
}

// ExcluirRamalFila - Remove um Ramal da fila
func (s FilaService) ExcluirRamalFila(id int, ramalID int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	sRamalID := strconv.Itoa(ramalID)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaFila, sID+"/"+sRamalID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}
