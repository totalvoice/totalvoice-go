package api

import (
	"github.com/totalvoice/go-client/api/model"
)

// PerfilService service
type PerfilService struct {
	client  HTTPClient
	handler Response
}

// NewPerfilService - Serviço para o gerenciamento da Conta Perfil
func NewPerfilService(httpClient HTTPClient, handler Response) *PerfilService {

	service := &PerfilService{
		client:  httpClient,
		handler: handler,
	}

	return service
}

// MinhaConta - Consulta saldo atual
func (s PerfilService) MinhaConta() (*model.Perfil, error) {

	perfil := new(model.Perfil)
	http, err := s.client.GetResource(perfil, RotaConta, nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(perfil, http)
	return res.(*model.Perfil), err
}

// AtualizarConta - Consulta saldo atual
func (s PerfilService) AtualizarConta(perfil model.Perfil) (*model.Perfil, error) {

	http, err := s.client.UpdateResource(perfil, RotaConta, nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(perfil, http)
	return res.(*model.Perfil), err
}

// RelatorioRecarga - Gera um relatorio com as recargas efetuadas
func (s PerfilService) RelatorioRecarga() (*model.Recarga, error) {

	recarga := new(model.Recarga)
	http, err := s.client.GetResource(recarga, RotaConta+"/recargas", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(recarga, http)
	return res.(*model.Recarga), err
}

// GeraURLRecarga - Gera uma URL para recarga de créditos
func (s PerfilService) GeraURLRecarga() (*model.URLRecarga, error) {

	url := new(model.URLRecarga)
	http, err := s.client.GetResource(url, RotaConta+"/urlrecarga", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(url, http)
	return res.(*model.URLRecarga), err
}
