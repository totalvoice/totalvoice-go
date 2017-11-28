package api

import "strconv"

// Conta struct
type Conta struct {
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

// ContaService service
type ContaService struct {
	Client HTTPClient

	Relatorio *ContaRelatorioService
}

// Criar - Cria uma nova conta na plataforma
func (s ContaService) Criar(values map[string]string) (*Conta, error) {
	conta := new(Conta)
	err := s.Client.CreateResource(values, RotaConta, conta)
	return conta, err
}

// BuscaConta - Leitura dos dados de uma conta criada
func (s ContaService) BuscaConta(id int) (*Conta, error) {
	sID := strconv.Itoa(id)
	conta := new(Conta)
	err := s.Client.GetResource(RotaConta, sID, conta)
	return conta, err
}

// Excluir - Remove uma Conta
func (s ContaService) Excluir(id int) error {
	sID := strconv.Itoa(id)
	return s.Client.DeleteResource(RotaConta, sID)
}

// Atualizar - Atualiza os dados de uma Conta criada
func (s ContaService) Atualizar(id int, values map[string]string) (*Conta, error) {
	sID := strconv.Itoa(id)
	conta := new(Conta)
	err := s.Client.UpdateResource(values, RotaConta, sID, conta)
	return conta, err
}
