package model

// Fila model
type Fila struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	EstrategiaRing string `json:"estrategia_ring"`
	TimeoutRing    string `json:"timeout_ring"`
	RamalID        string `json:"ramal_id"`
}

// FilaMembro model
type FilaMembro struct {
	RamalID int `json:"ramal_id"`
}

// FilaResponse model
type FilaResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID                 int    `json:"id"`
		Nome               string `json:"nome"`
		Chamadas           int    `json:"chamadas"`
		Completado         int    `json:"completado"`
		Cancelado          int    `json:"cancelado"`
		TempoFalado        string `json:"tempo_falado"`
		TempoEspera        string `json:"tempo_espera"`
		FilaMembroResponse []struct {
			ID                 int    `json:"id"`
			Nome               string `json:"nome"`
			Ramal              string `json:"ramal"`
			Login              string `json:"login"`
			Prioridade         string `json:"prioridade"`
			QtdLigacaoAtendida string `json:"qtd_ligacao_atendida"`
			UltimaChamada      string `json:"ultima_chamada"`
			EmLigacao          bool   `json:"em_ligacao"`
			Status             string `json:"status"`
			EmPausa            bool   `json:"em_pausa"`
			RazaoPausa         string `json:"razao_pausa"`
			Bina               string `json:"bina"`
			TempoStatus        string `json:"tempo_status"`
		} `json:"ramais"`
	} `json:"dados"`
}
