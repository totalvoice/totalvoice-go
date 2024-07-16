package totalvoice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropriedadesComValoresCorretos(t *testing.T) {

	client := &Client{accessToken: "meu-access-token", baseURI: BaseURI}
	assert.Equal(t, "meu-access-token", client.accessToken)
	assert.Equal(t, "https://voice-api.zenvia.com", client.baseURI)

	client2 := &Client{accessToken: "meu-access-token-2", baseURI: "https://meudominio.com.br"}
	assert.Equal(t, "meu-access-token-2", client2.accessToken)
	assert.Equal(t, "https://meudominio.com.br", client2.baseURI)
}

func TestMetodoGetBaseURI(t *testing.T) {

	client := &Client{accessToken: "meu-access-token", baseURI: "https://meudominio.com.br"}
	assert.Equal(t, "https://meudominio.com.br", client.GetBaseURI())
}

func TestMetodoBuildURI(t *testing.T) {

	client := &Client{accessToken: "meu-access-token", baseURI: BaseURI}
	path := "/path"
	uri := client.buildURI(path)

	assert.Equal(t, "https://voice-api.zenvia.com/path", uri)
}

func TestBuildQueryStringRetornaValorCorreto(t *testing.T) {

	client := &Client{accessToken: "meu-access-token", baseURI: BaseURI}
	params := map[string]interface{}{
		"value1": 1,
		"value2": true,
		"value3": "2",
	}
	query := client.buildQueryString(params)

	assert.Equal(t, "value1=1&value2=true&value3=2", query)
}
