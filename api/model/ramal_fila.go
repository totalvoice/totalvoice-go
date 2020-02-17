package model

// Ramal model
type RamalFila struct {
	ID              int    `json:"id"`
	Fila			string `json: "fila"`
	PausaStatus		bool   `json: "pausa_status"`
}

// RamalResponse struct
type RamalFilaResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID              int    `json:"id"`
		Fila			string `json: "fila"`
		PausaStatus		bool   `json: "pausa_status"`
	} `json:"dados"`
}
