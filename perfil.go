package totalvoice

// ConsultaSaldo - Consulta saldo atual
func (tvce *TotalVoice) ConsultaSaldo() (string, error) {
	return tvce.Get("/saldo")
}
