package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ContaService service
type ContaService struct {
	client   HTTPClient
	response Response

	Relatorio *ContaRelatorioService
}

// NewContaService - Serviço para o gerenciamento de Contas
func NewContaService(httpClient HTTPClient, response Response) *ContaService {

	service := &ContaService{
		client:   httpClient,
		response: response,
		Relatorio: &ContaRelatorioService{
			client:   httpClient,
			response: response,
		},
	}

	return service
}

// Criar - Cria uma nova conta na plataforma
func (s ContaService) Criar(conta model.Conta) (*model.ContaResponse, error) {

	resp := new(model.ContaResponse)
	http, err := s.client.CreateResource(conta, RotaConta)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ContaResponse), err
}

// Buscar - Leitura dos dados de uma conta criada
func (s ContaService) Buscar(id int) (*model.ContaResponse, error) {

	sID := strconv.Itoa(id)
	conta := new(model.Conta)
	resp := new(model.ContaResponse)

	http, err := s.client.GetResource(conta, RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ContaResponse), err
}

// Excluir - Remove uma Conta
func (s ContaService) Excluir(id int) (*model.ContaResponse, error) {

	sID := strconv.Itoa(id)
	resp := new(model.ContaResponse)

	http, err := s.client.DeleteResource(RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ContaResponse), err
}

// Atualizar - Atualiza os dados de uma Conta criada
func (s ContaService) Atualizar(conta model.Conta) (*model.ContaResponse, error) {

	sID := strconv.Itoa(conta.ID)
	resp := new(model.ContaResponse)

	http, err := s.client.UpdateResource(conta, RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ContaResponse), err
}
