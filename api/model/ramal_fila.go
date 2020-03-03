package model

// Ramal model
type RamalFila struct {
	StatusPausa		bool   	`json:"status_pausa"`
	Fila			int 	`json:"fila"`
}

// RamalResponse struct
type RamalFilaResponse struct {
	Status   int	`json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		RamalId 	int	`json:"ramal_id"`
		StatusPausa	bool 	`json:"pausado"`
		Fila		int	`json:"fila"`
	} `json:"dados"`
}
