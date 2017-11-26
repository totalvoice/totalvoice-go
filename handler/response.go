package handler

// TvceResponse -
type TvceResponse struct {
	Status   int                 `json:"status"`
	Sucesso  bool                `json:"sucesso"`
	Motivo   int                 `json:"motivo"`
	Mensagem string              `json:"mensagem"`
	Dados    map[string]struct{} `json:"dados"`
}
