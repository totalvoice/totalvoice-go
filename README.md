# totalvoice-go
Cliente em Golang para consumo da API da TotalVoice

[![Build Status](https://travis-ci.org/totalvoice/totalvoice-go.svg?branch=master&style=flat-square)](https://travis-ci.org/totalvoice/totalvoice-go)

> ### Funcionalidades

- Gerenciamento das chamadas
- Consulta e envio de SMS
- Consulta e envio de TTS
- Consulta e envio de Audio
- Consulta e envio de Composto
- Consulta e envio de Conferência
- Gerenciamento da Conta
- Gerenciamento da Central
- Gerenciamento de DID
- Gerenciamento de Valida Número

> ### Requisitos

- Golang 1.6+

> ### Instalação

Para instalar a biblioteca basta digitar

```sh
go get github.com/totalvoice/totalvoice-go
```
> ### Utilização

Para utilizar esta biblioteca, primeiramente você deverá realizar um cadastro no site da [Total Voice](http://www.totalvoice.com.br).
Após a criação do cadastro será disponibilizado um AccessToken para acesso a API.

Com o AccessToken em mãos será possível realizar as consultas/cadastros conforme documentação da [API](https://api.totalvoice.com.br/doc/#/)

Os métodos da API que poderão ser invocados:
- audio
- central
- chamada
- composto
- conferencia
- conta
- perfil
- sms
- tts
- did
- validaNumero

A seguir exemplos de como pode ser utilizada esta biblioteca.

> ##### Realiza uma chamada telefônica entre dois números: A e B

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Chamada.Criar("4811111111", "4822222222", nil)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Consulta de chamada pelo ID

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Chamada.Buscar(123) // ID da chamada

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```


> ##### Encerra uma chamada ativa

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Chamada.Encerrar(123) // ID da chamada

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Relatório de Chamadas

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

	client := totalvoice.NewTotalVoiceClient("access-token")
	di := time.Date(2017, time.December, 06, 00, 20, 0, 0, time.Local)
	df := time.Date(2017, time.December, 06, 18, 20, 0, 0, time.Local)
	// parametros opcionais - filtros e paginacao
	filtros := map[string]interface{}{
		"origem": "48999999999",
		"destino": "489888888888",
		"posicao": 0,
		"limite":  100,
	}

	response, err := client.Chamada.Relatorio.Gerar(di, df, filtros)

  if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Envio de SMS

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.SMS.Enviar("4811111111", "Minha mensagem SMS", false, false, nil)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Relatório de SMS

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

	client := totalvoice.NewTotalVoiceClient("access-token")
	di := time.Date(2017, time.December, 06, 00, 20, 0, 0, time.Local)
	df := time.Date(2017, time.December, 06, 18, 20, 0, 0, time.Local)
	response, err := client.SMS.Relatorio.Gerar(di, df)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Envio de TTS


```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.TTS.Enviar("4811111111", "Minha mensagem TTS", nil)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Envio de TTS com opções adicionais


```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    opcoes := map[string]interface{}{
        "velocidade": 1,
        "resposta_usuario": true,
        "tipo_voz": "br-Vitoria",
        "bina": "4811111111",
	}
    response, err := client.TTS.Enviar("4811111111", "Minha mensagem TTS", opcoes)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Relatório de TTS

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

	client := totalvoice.NewTotalVoiceClient("access-token")
	di := time.Date(2017, time.December, 06, 00, 20, 0, 0, time.Local)
	df := time.Date(2017, time.December, 06, 18, 20, 0, 0, time.Local)
	response, err := client.TTS.Relatorio.Gerar(di, df)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Envio de Audio


```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Audio.Enviar("4811111111", "http://foooo.bar/audio.mp3", false, "4811111111", false)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Configurações de central telefonica

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    ramal = new(Ramal)
    ramal.Ramal = "4001"
    ramal.Login = "meulogin@login.com.br"
    ramal.Senha = "12345789"
    response, err := client.Ramal.Criar(ramal)

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Gerenciamento dos dados da Conta

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Conta.Buscar(123) // ID da Conta

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Visualizando os dados da Minha Conta

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Perfil.MinhaConta()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Consulta saldo da Minha Conta

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Saldo.ConsultaSaldo()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```
> ##### Gerar URL de Recarga

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Perfil.GeraURLRecarga()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Relatório de Recarga

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.Perfil.RelatorioRecarga()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Listando DIDs no estoque

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.DID.Estoque()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Adquirindo um DID

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewTotalVoiceClient("access-token")
    response, err := client.DID.Adquirir(123) // ID do DID

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Caso você necessite utilizar seu próprio endereço configurado na Total Voice

```go
package main

import (
	"fmt"

	"github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewClient("access-token", "https://seudominio.com.br")
    response, err := client.Saldo.ConsultaSaldo()

    if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

> ##### Usando o Valida Número para saber se o número está ativo

```go
package main

import (
    "fmt"

    "github.com/totalvoice/totalvoice-go"
)

func main() {

    client := totalvoice.NewClient("access-token", "https://seudominio.com.br")
    response, err := client.ValidaNumero.Criar("4811111111");

    if err != nil {
        panic(err)
    }

    fmt.Println(response)
}
```


Mais informações sobre os métodos disponíveis podem ser encontrados na documentação da [API](https://api.totalvoice.com.br/doc/#/)

> ### Contribua!

Quer contribuir? [clique aqui](https://github.com/totalvoice/totalvoice-go/blob/master/CONTRIBUTING.md)

> ### Licença

Esta biblioteca segue os termos de uso da [MIT](https://github.com/totalvoice/totalvoice-go/blob/master/LICENSE)
