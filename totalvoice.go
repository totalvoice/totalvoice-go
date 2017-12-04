package totalvoice

import (
	"github.com/totalvoice/go-client/api"
)

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// TotalVoice struct
type TotalVoice struct {
	client api.HTTPClient

	Perfil   *api.PerfilService
	Audio    *api.AudioService
	Webhook  *api.WebhookService
	Saldo    *api.SaldoService
	Conta    *api.ContaService
	Composto *api.CompostoService
	Chamada  *api.ChamadaService
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	client := &Client{accessToken: accessToken, baseURI: BaseURI}
	tvce := &TotalVoice{client: client}

	handler := api.Response{}

	tvce.Perfil = api.NewPerfilService(client, handler)
	tvce.Audio = api.NewAudioService(client, handler)
	tvce.Webhook = api.NewWebhookService(client, handler)
	tvce.Saldo = api.NewSaldoService(client, handler)
	tvce.Conta = api.NewContaService(client, handler)
	tvce.Composto = api.NewCompostoService(client, handler)
	tvce.Chamada = api.NewChamadaService(client, handler)

	return tvce
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	tvce := NewTotalVoiceClient(accessToken)
	tvce.client.SetBaseURI(baseURI)

	return tvce
}
