package totalvoice

import "net/url"

const (
	// RotaSaldo - rota para consulta de saldo
	RotaSaldo = "/saldo"
	// RotaConta - rota para consulta dos dados da conta
	RotaConta = "/conta"
	// RotaWebhook - rota para o webhook
	RotaWebhook = "/webhook"
)

// Perfil client
type Perfil struct {
	client *TotalVoice
}

// ConsultaSaldo - Consulta saldo atual
func (p Perfil) ConsultaSaldo() (string, error) {
	return p.client.Get(RotaSaldo)
}

// MinhaConta - Consulta saldo atual
func (p Perfil) MinhaConta() (string, error) {
	return p.client.Get(RotaConta)
}

// AtualizarConta - Consulta saldo atual
func (p Perfil) AtualizarConta(values url.Values) (string, error) {
	return p.client.Put(values, RotaConta)
}
