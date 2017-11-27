package api

// Perfil struct
type Perfil struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID              int         `json:"id"`
		Nome            string      `json:"nome"`
		CpfCnpj         string      `json:"cpf_cnpj"`
		Login           string      `json:"login"`
		Saldo           float64     `json:"saldo"`
		Telefone        string      `json:"telefone"`
		AccessToken     string      `json:"access_token"`
		PrecoFixo       string      `json:"preco_fixo"`
		PrecoCel        string      `json:"preco_cel"`
		PrecoRamal      string      `json:"preco_ramal"`
		EmailFinanceiro string      `json:"email_financeiro"`
		NomeFantasia    interface{} `json:"nome_fantasia"`
		MetodoPagamento string      `json:"metodo_pagamento"`
		FaturaAtual     float64     `json:"fatura_atual"`
	} `json:"dados"`
}

// Recarga -
type Recarga struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
	} `json:"dados"`
}

// URLRecarga -
type URLRecarga struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		URL string `json:"url"`
	} `json:"dados"`
}

// PerfilService service
type PerfilService struct {
	Client HTTPClient
}

// MinhaConta - Consulta saldo atual
func (p PerfilService) MinhaConta() (*Perfil, error) {
	perfil := new(Perfil)
	err := p.Client.GetResource(RotaConta, nil, perfil)
	return perfil, err
}

// AtualizarConta - Consulta saldo atual
func (p PerfilService) AtualizarConta(values map[string]string) (*Perfil, error) {
	perfil := new(Perfil)
	err := p.Client.UpdateResource(values, RotaConta, "", perfil)
	return perfil, err
}

// RelatorioRecarga - Gera um relatorio com as recargas efetuadas
func (p PerfilService) RelatorioRecarga() (*Recarga, error) {
	recarga := new(Recarga)
	err := p.Client.GetResource(RotaConta+"/recargas", nil, recarga)
	return recarga, err
}

// GeraURLRecarga - Gera uma URL para recarga de créditos
func (p PerfilService) GeraURLRecarga() (*URLRecarga, error) {
	url := new(URLRecarga)
	err := p.Client.GetResource(RotaConta+"/urlrecarga", nil, url)
	return url, err
}
