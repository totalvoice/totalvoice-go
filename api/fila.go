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

// Busca fila
func (s FilaService) Buscar(id int) (*model.FilaResponse, error) {

	fila := new(model.Fila)
	sID := strconv.Itoa(id)
	http, err := s.client.GetResource(fila, RotaFila, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(fila, http)
	return res.(*model.FilaResponse), err
}

// Busca ramal na fila
func (s FilaService) BuscarRamalFila(id int, ramalId int) (*model.FilaResponse, error) {

	fila := new(model.Fila)
	sID := strconv.Itoa(id)
	sRamalID := strconv.Itoa(ramalId)
	http, err := s.client.GetResource(fila, RotaFila+"/"+sID+"/"+sRamalID, nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(fila, http)
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

// Adicionar - Adiciona ramal na fila
func (s FilaService) AdicionaRamalFila(fila model.Fila) (*model.FilaResponse, error) {
	
	sID := strconv.Itoa(fila.ID)
	resp := new(model.FilaResponse)
	http, err := s.client.CreateResource(fila, RotaFila+"/"+sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.FilaResponse), err
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

// Excluir - Remove um Ramal da fila
func (s FilaService) ExcluirRamalFila(id int, ramalId int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	sRamalID := strconv.Itoa(ramalId)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaFila, sID+"/"+sRamalID) 
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}