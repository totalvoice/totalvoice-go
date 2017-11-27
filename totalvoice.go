package totalvoice

import "github.com/totalvoice/go-client/api"

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// TotalVoice struct
type TotalVoice struct {
	client api.HTTPClient

	Perfil  *api.PerfilService
	Audio   *api.AudioService
	Webhook *api.WebhookService
	Saldo   *api.SaldoService
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	c := &Client{accessToken: accessToken, baseURI: BaseURI}

	tvce := &TotalVoice{client: c}

	tvce.Perfil = &api.PerfilService{Client: c}
	tvce.Audio = &api.AudioService{Client: c, Relatorio: &api.AudioRelatorioService{Client: c}}
	tvce.Webhook = &api.WebhookService{Client: c}
	tvce.Saldo = &api.SaldoService{Client: c}

	return tvce
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	tvce := NewTotalVoiceClient(accessToken)
	tvce.client.SetBaseURI(baseURI)

	return tvce
}
