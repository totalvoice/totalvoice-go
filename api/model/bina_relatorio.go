package model

// BinaRelatorio struct
type BinaRelatorio struct {
}

// BinaRelatorioResponse struct
type BinaRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID             int    `json:"id"`
			NumeroTelefone string `json:"numero_telefone"`
			DataCriacao    string `json:"data_criacao"`
			Fixo           bool   `json:"fixo"`
			Confirmado     bool   `json:"confirmado"`
		} `json:"relatorio"`
	} `json:"dados"`
}
