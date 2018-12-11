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
func (s ChamadaRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time, filtros map[string]interface{}) (*model.ChamadaRelatorioResponse, error) {

	relatorio := new(model.ChamadaRelatorio)
	params := map[string]interface{}{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	for index, value := range filtros {
		switch index {
		case "origem":
			params["origem"] = value.(string)
		case "destino":
			params["destino"] = value.(string)
		case "posicao":
			params["posicao"] = value.(int)
		case "limite":
			params["limite"] = value.(int)
		}
	}

	response := new(model.ChamadaRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaChamada+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.ChamadaRelatorioResponse), err
}
