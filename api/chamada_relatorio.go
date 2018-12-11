package api

import (
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ChamadaRelatorioService service
type ChamadaRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de mensagens de Audio
func (s ChamadaRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time, limite int, posicao int) (*model.ChamadaRelatorioResponse, error) {

	relatorio := new(model.ChamadaRelatorio)
	params := map[string]interface{}{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
		"limite":      limite,
		"posicao":     posicao,
	}

	response := new(model.ChamadaRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaChamada+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.ChamadaRelatorioResponse), err
}
