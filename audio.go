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
	client HTTPClient
}

// Enviar - Envia uma mensagem de audio
func (p Audio) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string) (string, error) {

	params := map[string]string{
		"numero_destino":   numero,
		"url_audio":        urlAudio,
		"resposta_usuario": strconv.FormatBool(respostaUsuario),
		"bina":             bina,
	}

	return p.client.CreateResource(params, RotaAudio)
}

// BuscaAudio - Busca uma mensagem de audio pelo seu ID
func (p Audio) BuscaAudio(id int) (string, error) {
	sID := strconv.Itoa(id)
	return p.client.GetResource(RotaAudio+sID, nil)
}

// Relatorio - Relatório de mensagens de Audio
func (p Audio) Relatorio(dataInicial time.Time, dataFinal time.Time) (string, error) {

	params := map[string]string{
		"data_inicial": dataInicial.UTC().String(),
		"data_final":   dataFinal.UTC().String(),
	}
	return p.client.GetResource(RotaAudio+"relatorio", params)
}
