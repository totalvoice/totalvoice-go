package api

import (
	"github.com/totalvoice/totalvoice-go/api/model"
)

// RamalRelatorioService service
type RamalRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de mensagens de Ramal
func (s RamalRelatorioService) Gerar() (*model.RamalRelatorioResponse, error) {

	relatorio := new(model.RamalRelatorio)
	resp := new(model.RamalRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaRamal+"/relatorio", nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.RamalRelatorioResponse), err
}
