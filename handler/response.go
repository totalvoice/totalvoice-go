package handler

// DataResponse - Retorna os dados especificos de cada chamada a API
type DataResponse interface {
	GetData() DataResponse
}

// TvceResponse -
type TvceResponse struct {
	Status   int       `json:"status"`
	Sucesso  bool      `json:"sucesso"`
	Motivo   int       `json:"motivo"`
	Mensagem string    `json:"mensagem"`
	Dados    *struct{} `json:"dados"`
}
