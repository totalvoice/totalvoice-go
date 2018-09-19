package api

import (
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// DIDRelatorioService service
type DIDRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de mensagens de Chamadas DID
func (s DIDRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.DIDRelatorioResponse, error) {

	relatorio := new(model.DIDRelatorio)
	params := map[string]interface{}{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.DIDRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaDID+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.DIDRelatorioResponse), err
}
