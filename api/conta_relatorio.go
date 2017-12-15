package api

import "github.com/totalvoice/totalvoice-go/api/model"

// ContaRelatorioService service
type ContaRelatorioService struct {
	client   HTTPClient
	response Response
}

// Gerar - Lista contas criadas por mim
func (s ContaRelatorioService) Gerar() (*model.ContaRelatorioResponse, error) {

	relatorio := new(model.ContaRelatorio)
	resp := new(model.ContaRelatorioResponse)

	http, err := s.client.ListResource(relatorio, RotaConta+"/relatorio", nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ContaRelatorioResponse), err
}
