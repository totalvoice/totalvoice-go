package api

import "strconv"

// Conta struct
type Conta struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID       int     `json:"id"`
		Nome     string  `json:"nome"`
		CpfCnpj  string  `json:"cpf_cnpj"`
		Login    string  `json:"login"`
		Saldo    float64 `json:"saldo"`
		Telefone string  `json:"telefone"`
	} `json:"dados"`
}

// ContaService service
type ContaService struct {
	Client HTTPClient
}

// BuscaConta - Leitura dos dados de uma conta criada
func (s ContaService) BuscaConta(id int) (*Conta, error) {
	sID := strconv.Itoa(id)
	conta := new(Conta)
	err := s.Client.GetResource(RotaConta, sID, conta)
	return conta, err
}
