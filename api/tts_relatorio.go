package api

import (
	"time"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// TTSRelatorioService service
type TTSRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Relatório de mensagens de TTS
func (s TTSRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*model.TTSRelatorioResponse, error) {

	relatorio := new(model.TTSRelatorio)
	params := map[string]interface{}{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	resp := new(model.TTSRelatorioResponse)
	http, err := s.client.ListResource(relatorio, RotaTTS+"/relatorio", params)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TTSRelatorioResponse), err
}
