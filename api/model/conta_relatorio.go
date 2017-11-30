package model

// ContaRelatorio Model
type ContaRelatorio struct{}

// ContaRelatorioResponse struct
type ContaRelatorioResponse struct {
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
