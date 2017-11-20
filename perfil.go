package totalvoice

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

// RelatorioRecarga - Gera um relatorio com as recargas efetuadas
func (p Perfil) RelatorioRecarga() (string, error) {
	return p.client.Get(RotaConta + "/recargas")
}

// AtualizarConta - Consulta saldo atual
func (p Perfil) AtualizarConta(values map[string]string) (string, error) {
	return p.client.Put(values, RotaConta)
}

// GeraURLRecarga - Gera uma URL para recarga de créditos
func (p Perfil) GeraURLRecarga() (string, error) {
	return p.client.Get(RotaConta + "/urlrecarga")
}
