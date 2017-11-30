package model

// Recarga -
type Recarga struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		URL string `json:"url"`
	} `json:"dados"`
}
