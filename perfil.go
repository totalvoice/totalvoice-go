package totalvoice

// Perfil client
type Perfil struct {
	client *TotalVoice
}

// ConsultaSaldo - Consulta saldo atual
func (p *Perfil) ConsultaSaldo() (string, error) {
	return p.client.Get("/saldo")
}
