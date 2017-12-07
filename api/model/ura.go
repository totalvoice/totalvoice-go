package model

// URA model
type URA struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Dados []Acao `json:"dados"`
}

// AddAcao - Adiciona uma ação no JSON
func (ura *URA) AddAcao(acao Acao) []Acao {
	ura.Dados = append(ura.Dados, acao)
	return ura.Dados
}
