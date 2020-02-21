package model

// WebhookResponse - Struct de retorno
type WebhookDefaultResponse struct {
	URL string `json:"url"`
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Webhook []struct {
			Webhook string `json:"webhook"`
			URL     string `json:"url"`
		} `json:"webhooks"`
	} `json:"dados"`
}
