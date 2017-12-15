package api

import "github.com/totalvoice/totalvoice-go/api/model"

// WebhookService service
type WebhookService struct {
	client   HTTPClient
	response Response
}

// NewWebhookService - Serviço para o gerenciamento de Webhooks
func NewWebhookService(httpClient HTTPClient, response Response) *WebhookService {

	service := &WebhookService{
		client:   httpClient,
		response: response,
	}

	return service
}

// Listar - Consulta saldo atual
func (s WebhookService) Listar() (*model.WebhookResponse, error) {
	webhook := new(model.WebhookResponse)
	http, err := s.client.ListResource(webhook, RotaWebhook, nil)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(webhook, http)
	return res.(*model.WebhookResponse), err
}

// Excluir - Apaga um webhook
func (s WebhookService) Excluir(nome string) (*model.WebhookResponse, error) {
	webhook := new(model.WebhookResponse)
	http, err := s.client.DeleteResource(RotaWebhook, nome)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(webhook, http)
	return res.(*model.WebhookResponse), err
}

// Salva - Cadastra ou atualiza um webhook
func (s WebhookService) Salva(nome, url string) (*model.Webhook, error) {

	webhook := new(model.Webhook)
	webhook.URL = url
	response := new(model.WebhookResponse)

	http, err := s.client.UpdateResource(webhook, RotaWebhook, nome)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.Webhook), err
}
