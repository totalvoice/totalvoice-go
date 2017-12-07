package model

// RamalRelatorio struct
type RamalRelatorio struct{}

// RamalRelatorioResponse struct
type RamalRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID             int    `json:"id"`
			Ramal          string `json:"ramal"`
			Login          string `json:"login"`
			Senha          string `json:"senha"`
			LigacaoExterna bool   `json:"ligacao_externa"`
		} `json:"relatorio"`
	} `json:"dados"`
}
