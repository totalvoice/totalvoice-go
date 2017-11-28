package api

// ContaRelatorio struct
type ContaRelatorio struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID              int         `json:"id"`
			Nome            string      `json:"nome"`
			CpfCnpj         string      `json:"cpf_cnpj"`
			Login           string      `json:"login"`
			Saldo           int         `json:"saldo"`
			Telefone        string      `json:"telefone"`
			AccessToken     string      `json:"access_token"`
			PrecoFixo       string      `json:"preco_fixo"`
			PrecoCel        string      `json:"preco_cel"`
			PrecoRamal      string      `json:"preco_ramal"`
			EmailFinanceiro string      `json:"email_financeiro"`
			NomeFantasia    interface{} `json:"nome_fantasia"`
			MetodoPagamento string      `json:"metodo_pagamento"`
			FaturaAtual     int         `json:"fatura_atual"`
		} `json:"relatorio"`
	} `json:"dados"`
}

// ContaRelatorioService service
type ContaRelatorioService struct {
	Client HTTPClient
}

// Gerar - Lista contas criadas por mim
func (s ContaRelatorioService) Gerar() (*ContaRelatorio, error) {
	resp := new(ContaRelatorio)
	err := s.Client.ListResource(RotaConta+"/relatorio", resp, nil)
	return resp, err
}
