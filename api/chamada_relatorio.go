package api

import (
	"time"

	"github.com/totalvoice/go-client/api/model"
)

// ChamadaRelatorioService service
type ChamadaRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens de Audio
func (s ChamadaRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.ChamadaRelatorioResponse, error) {

	relatorio := new(model.ChamadaRelatorio)
	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.ChamadaRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaChamada+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ChamadaRelatorioResponse), err
}
