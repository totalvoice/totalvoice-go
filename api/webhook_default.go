package api

import "github.com/totalvoice/totalvoice-go/api/model"

type WebhookDefaultService struct {
	client   HTTPClient
	response Response
}

// Listar - Listar todos webhooks default
func (s WebhookDefaultService) Listar() (*model.WebhookDefaultResponse, error) {
	webhook := new(model.WebhookDefaultResponse)
	http, err := s.client.ListResource(webhook, RotaWebhookDefault, nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(webhook, http)
	return res.(*model.WebhookDefaultResponse), err
}

// Excluir - Apaga um webhook default
func (s WebhookDefaultService) Excluir(nome string) (*model.WebhookDefaultResponse, error) {
	webhook := new(model.WebhookDefaultResponse)
	http, err := s.client.DeleteResource(RotaWebhookDefault, nome)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(webhook, http)
	return res.(*model.WebhookDefaultResponse), err
}

// Salva - Cadastra ou atualiza um webhook default
func (s WebhookDefaultService) Salvar(nome, url string) (*model.WebhookDefault, error) {

	webhook := new(model.WebhookDefaultResponse)
	webhook.URL = url
	response := new(model.WebhookDefaultResponse)

	http, err := s.client.UpdateResource(webhook, RotaWebhookDefault, nome)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.WebhookDefault), err
}