package model

// CompostoRelatorio struct
type CompostoRelatorio struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// CompostoRelatorioResponse struct
type CompostoRelatorioResponse struct {
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
