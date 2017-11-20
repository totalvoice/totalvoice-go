package totalvoice

import (
	"strconv"
	"time"
)

const (
	// RotaAudio - rota para manipulação de audios na API
	RotaAudio = "/audio/"
)

// Audio client
type Audio struct {
	client *TotalVoice
}

// Enviar - Envia uma mensagem de audio
func (p Audio) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string) (string, error) {

	params := map[string]string{
		"numero_destino":   numero,
		"url_audio":        urlAudio,
		"resposta_usuario": strconv.FormatBool(respostaUsuario),
		"bina":             bina,
	}

	return p.client.Post(params, RotaAudio)
}

// BuscaAudio - Busca uma mensagem de audio pelo seu ID
func (p Audio) BuscaAudio(id int) (string, error) {
	sID := strconv.Itoa(id)
	return p.client.Get(RotaAudio + sID)
}

// Relatorio - Relatório de mensagens de Audio
func (p Audio) Relatorio(dataInicial time.Time, dataFinal time.Time) (string, error) {

	return p.client.Get(RotaAudio + "relatorio")
}
