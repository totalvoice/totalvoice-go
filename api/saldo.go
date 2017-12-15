package api

import "github.com/totalvoice/totalvoice-go/api/model"

// SaldoService service
type SaldoService struct {
	client   HTTPClient
	response Response
}

// NewSaldoService - Serviço para a consulta de Saldo
func NewSaldoService(httpClient HTTPClient, response Response) *SaldoService {

	service := &SaldoService{
		client:   httpClient,
		response: response,
	}

	return service
}

// ConsultaSaldo - Consulta saldo atual
func (s SaldoService) ConsultaSaldo() (*model.Saldo, error) {

	saldo := new(model.Saldo)
	http, err := s.client.GetResource(saldo, RotaSaldo, nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(saldo, http)
	return res.(*model.Saldo), err
}
