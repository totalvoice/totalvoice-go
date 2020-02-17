package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// RamalFilaService service
type RamalFilaService struct {
	client   HTTPClient
	response Response
}

// NewRamalService - Serviço para o gerenciamento de Ramais
func NewRamalFilaService(httpClient HTTPClient, response Response) *RamalFilaService {

	service := &RamalFilaService{
		client:   httpClient,
		response: response,
	}

	return service
}

// Atualizar - Atualiza um ramal
func (s RamalFilaService) Atualizar_ramal_fila(ramal model.RamalFila) (*model.RamalFilaResponse, error) {

	sID := strconv.Itoa(ramal.ID)
	resp := new(model.RamalFilaResponse)
	http, err := s.client.UpdateResource(ramal, RotaRamal, sID + RotaFila)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.RamalFilaResponse), err
}

