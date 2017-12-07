package api

import (
	"time"

	"github.com/totalvoice/go-client/api/model"
)

// TTSRelatorioService service
type TTSRelatorioService struct {
	client  HTTPClient
	handler Response
}

// Gerar - Relatório de mensagens de TTS
func (s TTSRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.TTSRelatorioResponse, error) {

	relatorio := new(model.TTSRelatorio)
	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	response := new(model.TTSRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaTTS+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TTSRelatorioResponse), err
}
