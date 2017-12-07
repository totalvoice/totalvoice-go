package model

// DID model
type DID struct {
	ID      int    `json:"id"`
	DidID   string `json:"did_id"`
	UraID   string `json:"ura_id"`
	RamalID string `json:"ramal_id"`
}

// DIDResponse model
type DIDResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Did []interface{} `json:"did"`
	} `json:"dados"`
}
