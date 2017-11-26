package api

const (
	// RotaWebhook - rota para o webhook
	RotaWebhook = "/webhook"
)

// Webhook - Struct de retorno
type Webhook struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Webhooks []struct {
			Webhook string `json:"webhook"`
			URL     string `json:"url"`
		} `json:"webhooks"`
	} `json:"dados"`
}

// WebhookService service
type WebhookService struct {
	Client HTTPClient
}

// Listar - Consulta saldo atual
func (s WebhookService) Listar() (*Webhook, error) {
	webhook := new(Webhook)
	err := s.Client.ListResource(RotaWebhook, webhook, nil)
	return webhook, err
}

// Excluir - Retorna a lista de webhooks configurados para esta conta
func (s WebhookService) Excluir(nome string) error {
	return s.Client.DeleteResource(RotaWebhook, nome)
}

// Salva - Cadastra ou atualiza um webhook
func (s WebhookService) Salva(nome string, url string) error {
	params := map[string]string{"url": url}
	return s.Client.UpdateResource(params, RotaWebhook, nome)
}
