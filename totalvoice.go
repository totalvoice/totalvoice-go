package totalvoice

import (
	"github.com/totalvoice/totalvoice-go/api"
)

// BaseURI - URI Base
const (
	BaseURI = "https://voice-api.zenvia.com"
)

// TotalVoice struct
type TotalVoice struct {
	client api.HTTPClient

	Perfil         *api.PerfilService
	Bina           *api.BinaService
	Audio          *api.AudioService
	Webhook        *api.WebhookService
	WebhookDefault *api.WebhookDefaultService
	Saldo          *api.SaldoService
	Conta          *api.ContaService
	Composto       *api.CompostoService
	Chamada        *api.ChamadaService
	Fila           *api.FilaService
	Conferencia    *api.ConferenciaService
	SMS            *api.SMSService
	TTS            *api.TTSService
	Ramal          *api.RamalService
	RamalFila      *api.RamalFilaService
	URA            *api.URAService
	DID            *api.DIDService
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	client := NewClient(accessToken, BaseURI)
	return client
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	request := &Client{accessToken: accessToken, baseURI: baseURI}
	tvce := &TotalVoice{client: request}

	response := api.Response{}

	tvce.Perfil = api.NewPerfilService(request, response)
	tvce.Bina = api.NewBinaService(request, response)
	tvce.Audio = api.NewAudioService(request, response)
	tvce.Webhook = api.NewWebhookService(request, response)
	tvce.WebhookDefault = api.NewWebhookDefaultService(request, response)
	tvce.Saldo = api.NewSaldoService(request, response)
	tvce.Fila = api.NewFilaService(request, response)
	tvce.Conta = api.NewContaService(request, response)
	tvce.Composto = api.NewCompostoService(request, response)
	tvce.Chamada = api.NewChamadaService(request, response)
	tvce.Conferencia = api.NewConferenciaService(request, response)
	tvce.SMS = api.NewSMSService(request, response)
	tvce.TTS = api.NewTTSService(request, response)
	tvce.Ramal = api.NewRamalService(request, response)
	tvce.RamalFila = api.NewRamalFilaService(request, response)
	tvce.URA = api.NewURAService(request, response)
	tvce.DID = api.NewDIDService(request, response)
	return tvce
}
