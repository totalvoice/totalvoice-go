package api

import (
	"time"

	"github.com/totalvoice/go-client/api/model"
)

// SMSRelatorioService service
type SMSRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens de SMS
func (s SMSRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.SMSRelatorioResponse, error) {

	relatorio := new(model.SMSRelatorio)
	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.SMSRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaSMS+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.SMSRelatorioResponse), err
}
