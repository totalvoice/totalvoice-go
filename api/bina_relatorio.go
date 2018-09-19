package api

import (
	"github.com/totalvoice/totalvoice-go/api/model"
)

// BinaRelatorioService service
type BinaRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Gera relatório com os números cadastrados
func (s BinaRelatorioService) Gerar() (*model.BinaRelatorioResponse, error) {

	relatorio := new(model.BinaRelatorio)

	response := new(model.BinaRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaBina+"/relatorio", nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.BinaRelatorioResponse), err
}
