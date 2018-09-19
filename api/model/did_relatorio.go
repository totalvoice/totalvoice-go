package model

// DIDRelatorio struct
type DIDRelatorio struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// DIDRelatorioResponse struct
type DIDRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID                     int    `json:"id"`
			Ativa                  bool   `json:"ativa"`
			URLGravacao            string `json:"url_gravacao"`
			DataInicio             string `json:"data_inicio"`
			NumeroDestino          string `json:"numero_destino"`
			NumeroOrigem           string `json:"numero_origem"`
			Status                 string `json:"status"`
			DuracaoSegundos        int    `json:"duracao_segundos"`
			Duracao                string `json:"duracao"`
			DuracaoCobradaSegundos int    `json:"duracao_cobrada_segundos"`
			DuracaoCobrada         string `json:"duracao_cobrada"`
			DuracaoFaladaSegundos  int    `json:"duracao_falada_segundos"`
			DuracaoFalada          string `json:"duracao_falada"`
			Preco                  int    `json:"preco"`
			RamalID                int    `json:"ramal_id"`
		} `json:"relatorio"`
	} `json:"dados"`
}
