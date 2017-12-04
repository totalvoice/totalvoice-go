package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// ContaService service
type ContaService struct {
	client  HTTPClient
	handler Response

	Relatorio *ContaRelatorioService
}

// NewContaService - Serviço para o gerenciamento de Contas
func NewContaService(httpClient HTTPClient, handler Response) *ContaService {

	service := &ContaService{
		client:  httpClient,
		handler: handler,
		Relatorio: &ContaRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Criar - Cria uma nova conta na plataforma
func (s ContaService) Criar(conta model.Conta) (*model.ContaResponse, error) {

	response := new(model.ContaResponse)
	http, err := s.client.CreateResource(conta, RotaConta)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ContaResponse), err
}

// BuscaConta - Leitura dos dados de uma conta criada
func (s ContaService) BuscaConta(id int) (*model.ContaResponse, error) {

	sID := strconv.Itoa(id)
	conta := new(model.Conta)
	response := new(model.ContaResponse)

	http, err := s.client.GetResource(conta, RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ContaResponse), err
}

// Excluir - Remove uma Conta
func (s ContaService) Excluir(id int) (*model.ContaResponse, error) {

	sID := strconv.Itoa(id)
	response := new(model.ContaResponse)

	http, err := s.client.DeleteResource(RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ContaResponse), err
}

// Atualizar - Atualiza os dados de uma Conta criada
func (s ContaService) Atualizar(conta model.Conta) (*model.ContaResponse, error) {

	sID := strconv.Itoa(conta.ID)
	response := new(model.ContaResponse)

	http, err := s.client.UpdateResource(conta, RotaConta, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.ContaResponse), err
}
