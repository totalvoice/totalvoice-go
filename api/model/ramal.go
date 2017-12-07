package model

// Ramal model
type Ramal struct {
	ID              int    `json:"id"`
	Ramal           string `json:"ramal"`
	Login           string `json:"login"`
	Senha           string `json:"senha"`
	Bina            string `json:"bina"`
	LigacaoExterna  bool   `json:"ligacao_externa"`
	LigacaoCelular  bool   `json:"ligacao_celular"`
	GravarAudio     bool   `json:"gravar_audio"`
	AcessoGravacoes bool   `json:"acesso_gravacoes"`
	UraID           int    `json:"ura_id"`
	Voicemail       bool   `json:"voicemail"`
}

// RamalResponse struct
type RamalResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID              int    `json:"id"`
		Ramal           string `json:"ramal"`
		Login           string `json:"login"`
		Senha           string `json:"senha"`
		LigacaoExterna  bool   `json:"ligacao_externa"`
		LigacaoCelular  bool   `json:"ligacao_celular"`
		GravarAudio     bool   `json:"gravar_audio"`
		AcessoGravacoes bool   `json:"acesso_gravacoes"`
	} `json:"dados"`
	UraID int `json:"ura_id"`
}
