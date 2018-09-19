package totalvoice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstrutorNewClientRetornaInstanciaCorreta(t *testing.T) {

	tvce := NewClient("meu-access-token", "https://meudominio.com.br")
	assert.Equal(t, "*totalvoice.TotalVoice", reflect.TypeOf(tvce).String())
}

func TestConstrutorNewTotalVoiceClientRetornaInstanciaCorreta(t *testing.T) {

	tvce := NewTotalVoiceClient("meu-access-token")
	assert.Equal(t, "*totalvoice.TotalVoice", reflect.TypeOf(tvce).String())
}

func TestConstrutorRetornaInstanciaDosServicos(t *testing.T) {

	tvce := NewTotalVoiceClient("meu-access-token")

	assert.Equal(t, "*api.PerfilService", reflect.TypeOf(tvce.Perfil).String())
	assert.Equal(t, "*api.AudioService", reflect.TypeOf(tvce.Audio).String())
	assert.Equal(t, "*api.WebhookService", reflect.TypeOf(tvce.Webhook).String())
	assert.Equal(t, "*api.SaldoService", reflect.TypeOf(tvce.Saldo).String())
	assert.Equal(t, "*api.ContaService", reflect.TypeOf(tvce.Conta).String())
	assert.Equal(t, "*api.CompostoService", reflect.TypeOf(tvce.Composto).String())
	assert.Equal(t, "*api.ChamadaService", reflect.TypeOf(tvce.Chamada).String())
	assert.Equal(t, "*api.ConferenciaService", reflect.TypeOf(tvce.Conferencia).String())
	assert.Equal(t, "*api.SMSService", reflect.TypeOf(tvce.SMS).String())
	assert.Equal(t, "*api.TTSService", reflect.TypeOf(tvce.TTS).String())
	assert.Equal(t, "*api.RamalService", reflect.TypeOf(tvce.Ramal).String())
	assert.Equal(t, "*api.URAService", reflect.TypeOf(tvce.URA).String())
	assert.Equal(t, "*api.DIDService", reflect.TypeOf(tvce.DID).String())
	assert.Equal(t, "*api.FilaService", reflect.TypeOf(tvce.Fila).String())
	assert.Equal(t, "*api.BinaService", reflect.TypeOf(tvce.Bina).String())
}
