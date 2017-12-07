package model

// URA model
type URA struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Dados []Acao `json:"dados"`
}
