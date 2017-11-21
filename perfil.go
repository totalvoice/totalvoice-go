package totalvoice

const (
	// RotaSaldo - rota para consulta de saldo
	RotaSaldo = "/saldo/"
	// RotaConta - rota para consulta dos dados da conta
	RotaConta = "/conta/"
	// RotaWebhook - rota para o webhook
	RotaWebhook = "/webhook/"
)

// Perfil client
type Perfil struct {
	client HTTPClient
}

// ConsultaSaldo - Consulta saldo atual
func (p Perfil) ConsultaSaldo() (string, error) {
	return p.client.GetResource(RotaSaldo, nil)
}

// MinhaConta - Consulta saldo atual
func (p Perfil) MinhaConta() (string, error) {
	return p.client.GetResource(RotaConta, nil)
}

// AtualizarConta - Consulta saldo atual
func (p Perfil) AtualizarConta(values map[string]string) (string, error) {
	return p.client.UpdateResource(values, RotaConta)
}

// RelatorioRecarga - Gera um relatorio com as recargas efetuadas
func (p Perfil) RelatorioRecarga() (string, error) {
	return p.client.GetResource(RotaConta+"recargas", nil)
}

// GeraURLRecarga - Gera uma URL para recarga de créditos
func (p Perfil) GeraURLRecarga() (string, error) {
	return p.client.GetResource(RotaConta+"urlrecarga", nil)
}

// Webhooks - Retorna a lista de webhooks configurados para esta conta
func (p Perfil) Webhooks() (string, error) {
	return p.client.GetResource(RotaWebhook, nil)
}

// ExcluirWebhook - Retorna a lista de webhooks configurados para esta conta
func (p Perfil) ExcluirWebhook(nome string) (string, error) {
	return p.client.DeleteResource(RotaWebhook + nome)
}

// SalvaWebhook - Cadastra ou atualiza um webhook
func (p Perfil) SalvaWebhook(nome string, url string) (string, error) {
	params := map[string]string{"url": url}
	return p.client.UpdateResource(params, RotaWebhook+nome)
}
