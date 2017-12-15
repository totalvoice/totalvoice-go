package api

import (
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// SMSRelatorioService service
type SMSRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de mensagens de SMS
func (s SMSRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.SMSRelatorioResponse, error) {

	relatorio := new(model.SMSRelatorio)
	params := map[string]interface{}{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	resp := new(model.SMSRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaSMS+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.SMSRelatorioResponse), err
}
