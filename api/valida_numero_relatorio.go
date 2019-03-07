package api

import (	
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ValidaNumeroRelatorioService service
type ValidaNumeroRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de ValidaNuero
func (s ValidaNumeroRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time, filtros map[string]interface{}) (*model.ValidaNumeroRelatorioResponse, error) {

	relatorio := new(model.ValidaNumeroRelatorio)
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

	resp := new(model.ValidaNumeroRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaValidaNumero+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ValidaNumeroRelatorioResponse), err
}
