package model

// ChamadaRelatorio struct
type ChamadaRelatorio struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// ChamadaRelatorioResponse struct
type ChamadaRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID            int         `json:"id"`
			DataCriacao   string      `json:"data_criacao"`
			Ativa         bool        `json:"ativa"`
			ClienteID     int         `json:"cliente_id"`
			RamalIDOrigem interface{} `json:"ramal_id_origem"`
			Origem        struct {
				DataInicio             string  `json:"data_inicio"`
				Numero                 string  `json:"numero"`
				Tipo                   string  `json:"tipo"`
				Status                 string  `json:"status"`
				DuracaoSegundos        int     `json:"duracao_segundos"`
				Duracao                string  `json:"duracao"`
				DuracaoCobradaSegundos int     `json:"duracao_cobrada_segundos"`
				DuracaoCobrada         string  `json:"duracao_cobrada"`
				DuracaoFaladaSegundos  int     `json:"duracao_falada_segundos"`
				DuracaoFalada          string  `json:"duracao_falada"`
				Preco                  float64 `json:"preco"`
			} `json:"origem"`
			Destino struct {
				DataInicio             string      `json:"data_inicio"`
				Numero                 string      `json:"numero"`
				Tipo                   string      `json:"tipo"`
				Status                 string      `json:"status"`
				DuracaoSegundos        int         `json:"duracao_segundos"`
				Duracao                string      `json:"duracao"`
				DuracaoCobradaSegundos int         `json:"duracao_cobrada_segundos"`
				DuracaoCobrada         string      `json:"duracao_cobrada"`
				DuracaoFaladaSegundos  interface{} `json:"duracao_falada_segundos"`
				DuracaoFalada          interface{} `json:"duracao_falada"`
				Preco                  float64     `json:"preco"`
			} `json:"destino"`
		} `json:"relatorio"`
	} `json:"dados"`
}
