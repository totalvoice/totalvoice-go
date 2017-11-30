package model

// Conta Model
type Conta struct {
	ID              int         `json:"id"`
	Nome            string      `json:"nome"`
	Login           string      `json:"login"`
	Senha           string      `json:"senha"`
	CpfCnpj         string      `json:"cpf_cnpj"`
	Telefone        string      `json:"telefone"`
	PrecoFixo       string      `json:"preco_fixo"`
	PrecoCel        string      `json:"preco_cel"`
	PrecoRamal      string      `json:"preco_ramal"`
	EmailFinanceiro string      `json:"email_financeiro"`
	NomeFantasia    interface{} `json:"nome_fantasia"`
}

// ContaResponse struct
type ContaResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID          int     `json:"id"`
		Nome        string  `json:"nome"`
		CpfCnpj     string  `json:"cpf_cnpj"`
		Login       string  `json:"login"`
		Saldo       float64 `json:"saldo"`
		Telefone    string  `json:"telefone"`
		AccessToken string  `json:"access_token"`
	} `json:"dados"`
}
