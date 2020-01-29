package model

// Webhook - Struct de retorno
type WebhookDefault struct {
	URL string `json:"url"`
}

// WebhookResponse - Struct de retorno
type WebhookDefaultResponse struct {
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
