package api

import (
	"github.com/totalvoice/go-client/api/model"
)

// RamalRelatorioService service
type RamalRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens de Ramal
func (s RamalRelatorioService) Gerar() (*model.RamalRelatorioResponse, error) {

	relatorio := new(model.RamalRelatorio)
	response := new(model.RamalRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaRamal+"/relatorio", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.RamalRelatorioResponse), err
}
