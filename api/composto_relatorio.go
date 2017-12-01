package api

import (
	"time"
)

// CompostoRelatorio struct
type CompostoRelatorio struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID                     int    `json:"id"`
			NumeroDestino          string `json:"numero_destino"`
			DataCriacao            string `json:"data_criacao"`
			DataInicio             string `json:"data_inicio"`
			Tipo                   string `json:"tipo"`
			Status                 string `json:"status"`
			DuracaoSegundos        int    `json:"duracao_segundos"`
			Duracao                string `json:"duracao"`
			DuracaoCobradaSegundos int    `json:"duracao_cobrada_segundos"`
			DuracaoCobrada         string `json:"duracao_cobrada"`
			DuracaoFaladaSegundos  int    `json:"duracao_falada_segundos"`
			DuracaoFalada          string `json:"duracao_falada"`
			Preco                  int    `json:"preco"`
			URLAudio               string `json:"url_audio"`
			RespostaUsuario        bool   `json:"resposta_usuario"`
			Resposta               string `json:"resposta"`
		} `json:"relatorio"`
	} `json:"dados"`
}

// CompostoRelatorioService service
type CompostoRelatorioService struct {
	Client HTTPClient
}

// Gerar - Relatório de mensagens de Audio
func (s CompostoRelatorioService) Gerar(dataInicial time.Time, dataFinal time.Time) (*AudioRelatorio, error) {

	params := map[string]string{
		"data_inicio": dataInicial.UTC().Format(DateFormat),
		"data_fim":    dataFinal.UTC().Format(DateFormat),
	}

	resp := new(AudioRelatorio)
	err := s.Client.ListResource(RotaAudio+"/relatorio", resp, params)
	return resp, err
}
