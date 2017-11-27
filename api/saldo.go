package api

// Saldo - struct
type Saldo struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Saldo int `json:"saldo"`
	} `json:"dados"`
}

// SaldoService service
type SaldoService struct {
	Client HTTPClient
}

// ConsultaSaldo - Consulta saldo atual
func (p SaldoService) ConsultaSaldo() (*Saldo, error) {
	saldo := new(Saldo)
	err := p.Client.GetResource(RotaSaldo, nil, saldo)
	return saldo, err
}
